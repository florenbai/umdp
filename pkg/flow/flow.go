package flow

import (
	"context"
	"encoding/json"
	"github.com/tidwall/gjson"
	"minos/pkg/response"
	"minos/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

// Flow 流程引擎
type Flow struct {
	Template string `json:"template"`
	FormData string `json:"formData"`
	// map nodeId:nodeType
	NodeTypeMap map[string]string `json:"nodeTypeMap"`
	ctx         context.Context
}

func NewFlow(ctx context.Context, template string, formData string) (*Flow, error) {
	if !gjson.Valid(template) {
		return nil, response.JsonParamErr
	}
	if !gjson.Valid(formData) {
		return nil, response.JsonParamErr
	}
	dataMap := getNodeTypeMap(template)
	return &Flow{
		ctx:         ctx,
		Template:    template,
		FormData:    formData,
		NodeTypeMap: dataMap,
	}, nil
}

const (
	StartNode = "start"
	EndNode   = "end"
	AuditNode = "rect"
	TermNode  = "polygon"
)

const (
	NodeArr = "nodes"
	edgeArr = "edges"
)

const (
	sourceNodeName = "sourceNodeId"
	targetNodeName = "targetNodeId"
)

const (
	NodeHandleAgree    = 1
	NodeHandleTransfer = 2
	NodeHandleReject   = 3
)

const (
	ConditionGE = ">"
	ConditionGT = ">="
	ConditionLE = "<"
	ConditionLT = "<="
	ConditionEQ = "="
	ConditionNE = "!="
)

type Node struct {
	Id             string   `json:"id"`             //节点编号
	NodeName       string   `json:"nodeName"`       //节点名称
	Reviewer       []string `json:"reviewer"`       //通知人
	Cc             []string `json:"cc"`             //抄送
	Task           uint64   `json:"task"`           //任务编号
	EnableFallback bool     `json:"enableFallback"` //是否开启退回
	Notify         []string `json:"notify"`         //通知类型 email,wechat
	AutomaticAudit bool     `json:"automaticAudit"` //是否开启自动审批
	AuditDays      int64    `json:"auditDays"`      //自动审批天数
}

type Term struct {
	Field      string `json:"field"`
	Condition  string `json:"condition"`
	Value      string `json:"value"`
	NodeId     string `json:"nodeId"`
	FailNodeId string `json:"failNodeId"`
}

type Overview struct {
	NodeId   string `json:"nodeId"`
	NodeName string `json:"nodeName"`
	Complete bool   `json:"complete"`
}

func (flow *Flow) GetFlowOverview() ([]Node, error) {
	var nodes []Node
	nid := ""
	for {
		node, err := flow.GetNextAuditNode(nid)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node...)
		if len(node) > 0 {
			nid = node[0].Id
		} else {
			break
		}
	}
	return nodes, nil
}

// GetNextAuditNode 获取下一个审核节点
// 返回节点数组（当存在并行节点时，返回节点可能是多个）
func (flow *Flow) GetNextAuditNode(nodeId string) ([]Node, error) {
	var nodes []Node
	templateJson := gjson.Parse(flow.Template)
	if len(templateJson.Get(NodeArr).Array()) < 1 {
		return nodes, response.NodeNotFound
	}
	if nodeId == "" {
		for k, v := range flow.NodeTypeMap {
			if v == StartNode {
				nodeId = k
			}
		}
	}
	target, targetType := flow.GetTargetIdArrBySourceId(nodeId)
	nodes = flow.GetAuditNode(target, targetType)
	return nodes, nil
}

func getNodeTypeMap(template string) map[string]string {
	data := make(map[string]string)
	gjson.Parse(template).Get(NodeArr).ForEach(func(key, value gjson.Result) bool {
		data[value.Get("id").String()] = value.Get("type").String()
		return true
	})
	return data
}

// GetPrevAuditNode 获取上一个审核节点编号
func (flow *Flow) GetPrevAuditNode(nodeId string) (string, map[string]string, error) {
	if nodeId == "" {
		return "", flow.NodeTypeMap, response.NodeNotFound
	}
	if !gjson.Valid(flow.Template) {
		return "", flow.NodeTypeMap, response.JsonParamErr
	}
	templateJson := gjson.Parse(flow.Template)
	if len(templateJson.Get(NodeArr).Array()) < 1 {
		return "", flow.NodeTypeMap, response.NodeNotFound
	}
	for {
		sourceId := flow.GetSourceIdByNodeId(nodeId)
		if strings.Contains(flow.NodeTypeMap[sourceId], AuditNode) || strings.Contains(flow.NodeTypeMap[sourceId], StartNode) {
			return sourceId, flow.NodeTypeMap, nil
		}
		nodeId = sourceId
	}
}

// GetSourceIdByNodeId 根据目标编号获取源节点编号
func (flow *Flow) GetSourceIdByNodeId(nodeId string) string {
	var sourceId string
	gjson.Parse(flow.Template).Get(edgeArr).ForEach(func(key, value gjson.Result) bool {
		if value.Get(targetNodeName).String() == nodeId {
			sourceId = value.Get(sourceNodeName).String()
			return false
		}
		return true
	})
	return sourceId
}

// GetTargetIdArrBySourceId 根据源编号获取目标编号
func (flow *Flow) GetTargetIdArrBySourceId(sourceId string) ([]string, string) {
	var targetIds []string
	var targetType string
	templateJson := gjson.Parse(flow.Template)
	templateJson.Get(edgeArr).ForEach(func(key, value gjson.Result) bool {
		if value.Get(sourceNodeName).String() == sourceId {
			targetIds = append(targetIds, value.Get(targetNodeName).String())
			targetType = flow.NodeTypeMap[value.Get(targetNodeName).String()]
		}
		return true
	})
	return targetIds, targetType
}

// GetAuditNode 获取节点
func (flow *Flow) GetAuditNode(target []string, targetType string) []Node {
	switch targetType {
	case AuditNode:
		return flow.auditNode(target)
	case TermNode:
		return flow.termNode(target)
	case EndNode:
		return nil
	}
	return nil
}

// termNode 从条件节点获取审核节点
func (flow *Flow) termNode(target []string) []Node {
	var nodes []Node
	var term Term
	templateJson := gjson.Parse(flow.Template)
	templateJson.Get(NodeArr).ForEach(func(key, value gjson.Result) bool {
		if value.Get("id").String() == target[0] {
			json.Unmarshal([]byte(value.Get("properties").String()), &term)
			return false
		}
		return true
	})
	auditNodes, _ := flow.GetTargetIdArrBySourceId(target[0])
	templateJson.Get(NodeArr).ForEach(func(key, value gjson.Result) bool {
		if utils.InOfT(value.Get("id").String(), auditNodes) {
			var reviewer []string
			var cc []string
			var notify []string
			value.Get("properties.reviewer").ForEach(func(key, value gjson.Result) bool {
				reviewer = append(reviewer, value.String())
				return true
			})
			value.Get("properties.cc").ForEach(func(key, value gjson.Result) bool {
				cc = append(cc, value.String())
				return true
			})
			value.Get("properties.notify").ForEach(func(key, value gjson.Result) bool {
				notify = append(notify, value.String())
				return true
			})
			node := Node{
				Id:             value.Get("id").String(),
				NodeName:       value.Get("properties.nodeName").String(),
				Reviewer:       reviewer,
				Cc:             cc,
				Task:           value.Get("properties.task").Uint(),
				EnableFallback: value.Get("properties.enableFallback").Bool(),
				Notify:         notify,
				AutomaticAudit: value.Get("properties.automaticAudit").Bool(),
				AuditDays:      value.Get("properties.auditDays").Int(),
			}
			if node.Id != term.NodeId {
				term.FailNodeId = node.Id
			}
			nodes = append(nodes, node)
		}
		return true
	})

	nodeId := flow.getTermNode(term)
	for _, n := range nodes {
		if nodeId == n.Id {
			return []Node{n}
		}
	}
	return nil
}

// AuditNode 根据目标节点编号返回目标节点
func (flow *Flow) auditNode(target []string) []Node {
	var nodes []Node
	templateJson := gjson.Parse(flow.Template)
	templateJson.Get(NodeArr).ForEach(func(key, value gjson.Result) bool {
		if utils.InOfT(value.Get("id").String(), target) {
			var reviewer []string
			var cc []string
			var notify []string
			value.Get("properties.reviewer").ForEach(func(key, value gjson.Result) bool {
				reviewer = append(reviewer, value.String())
				return true
			})
			value.Get("properties.cc").ForEach(func(key, value gjson.Result) bool {
				cc = append(cc, value.String())
				return true
			})
			value.Get("properties.notify").ForEach(func(key, value gjson.Result) bool {
				notify = append(notify, value.String())
				return true
			})
			node := Node{
				Id:             value.Get("id").String(),
				NodeName:       value.Get("properties.nodeName").String(),
				Reviewer:       reviewer,
				Cc:             cc,
				Task:           value.Get("properties.task").Uint(),
				EnableFallback: value.Get("properties.enableFallback").Bool(),
				Notify:         notify,
				AutomaticAudit: value.Get("properties.automaticAudit").Bool(),
				AuditDays:      value.Get("properties.auditDays").Int(),
			}
			nodes = append(nodes, node)
		}
		return true
	})
	return nodes
}

// GetNextNodeName 获取下一个节点名称
func (flow *Flow) GetNextNodeName(nodes []Node) string {
	var nameArr []string
	for _, v := range nodes {
		nameArr = append(nameArr, v.NodeName)
	}
	return strings.Join(nameArr, ",")
}

// 获取条件节点
func (flow *Flow) getTermNode(terms Term) string {
	var ok bool
	var node string
	formJson := gjson.Parse(flow.FormData)
	// 匹配条件
	field := formJson.Get(terms.Field).String()
	pattern := "(^[1-9]\\d*\\.\\d+$|^0\\.\\d+$|^[1-9]\\d*$|^0$)"
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(field) {
		thresholdFload, _ := strconv.ParseFloat(terms.Value, 64)
		ok = compare(terms.Condition, thresholdFload, formJson.Get(terms.Field).Float())
	} else {
		ok = compare(terms.Condition, terms.Value, formJson.Get(terms.Field).String())
	}
	if ok {
		node = terms.NodeId
	}
	// 当不存在匹配的node时，则选择默认的node
	if node == "" {
		node = terms.FailNodeId
	}
	return node
}

func compare[T int64 | float64 | float32 | string](condition string, threshold T, value T) bool {
	switch condition {
	case ConditionGE:
		if value >= threshold {
			return true
		}
	case ConditionGT:
		if value > threshold {
			return true
		}
	case ConditionLE:
		if value <= threshold {
			return true
		}
	case ConditionLT:
		if value < threshold {
			return true
		}
	case ConditionEQ:
		if value == threshold {
			return true
		}
	case ConditionNE:
		if threshold != value {
			return true
		}
	default:
		return false
	}
	return false
}
