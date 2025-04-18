package wework

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type EventPushQueryBinding struct {
	MsgSign   string `form:"msg_signature" binding:"required"`
	Timestamp string `form:"timestamp" binding:"required"`
	Nonce     string `form:"nonce" binding:"required"`
	EchoStr   string `form:"echostr"`
	CorpId    string `form:"corpid"`
}

type InfoType string

const (
	SuiteTicket        InfoType = "suite_ticket"
	CreateAuth         InfoType = "create_auth"
	ChangeAuth         InfoType = "change_auth"
	CancelAuth         InfoType = "cancel_auth"
	ResetPermanentCode InfoType = "reset_permanent_code"
)

type BizEvent struct {
	XMLName   xml.Name `xml:"xml"`
	Text      string   `xml:",chardata"`
	SuiteId   string   `xml:"SuiteId"`
	InfoType  InfoType `xml:"InfoType"`
	TimeStamp int64    `xml:"TimeStamp"`
}

// 数据回调通知
// 企业安装应用时会把相应数据通知到回调URL
type BizData struct {
	xml.Name   `xml:"xml"`
	Text       string `xml:",chardata"`
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
	AgentID    string `xml:"AgentID"`
}

// BizResponse 基础返回类型，定义错误代码及错误消息
type BizResponse struct {
	Error
}

func httpRequestWithContext(ctx context.Context, request *http.Request, resChan chan<- []byte) (err error) {
	request = request.WithContext(ctx)
	request.Header.Set("User-Agent", "wecom-golang-sdk")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("client.Do Error: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("response from weixin with status %v", resp.StatusCode)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll Error: %s", err.Error())
	}
	defer resp.Body.Close()
	resChan <- data
	return nil
}

func HttpGet(apiUrl string) (body []byte, err error) {
	resChan := make(chan []byte)
	repoUrl := fmt.Sprintf("%s%s", qyApiHost, apiUrl)
	request, err := http.NewRequest(http.MethodGet, repoUrl, nil)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(15)*time.Second))
	defer cancel()

	go httpRequestWithContext(ctx, request, resChan)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.Tick(time.Duration(15) * time.Second):
		return nil, errors.New("time over")
	case body = <-resChan:
		return body, nil
	}
}

func HttpPost(apiUrl string, params interface{}) (body []byte, err error) {
	resChan := make(chan []byte)
	repoUrl := fmt.Sprintf("%s%s", qyApiHost, apiUrl)
	data, err := json.Marshal(params)
	request, err := http.NewRequest(http.MethodPost, repoUrl, bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(15)*time.Second))
	defer cancel()

	go httpRequestWithContext(ctx, request, resChan)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.Tick(time.Duration(15) * time.Second):
		log.Println(repoUrl, string(data))
		return nil, errors.New("time over")
	case body = <-resChan:
		return body, nil
	}
}

func HttpUploadMedia(apiUrl string, filePath string, filename string) (body []byte, err error) {
	repoUrl := fmt.Sprintf("%s%s", qyApiHost, apiUrl)
	reader, writer := io.Pipe()
	request, err := http.NewRequest(http.MethodPost, repoUrl, reader)
	mwriter := multipart.NewWriter(writer)
	request.Header.Set("Content-Type", mwriter.FormDataContentType())
	errchan := make(chan error)

	go func() {
		defer close(errchan)
		defer writer.Close()
		defer mwriter.Close()
		if filename == "" {
			_, filename = filepath.Split(filePath)
		}
		w, err := mwriter.CreateFormFile("media", filename)
		if err != nil {
			errchan <- err
			return
		}
		n, err := os.Open(filePath)
		if err != nil {
			errchan <- err
			return
		}
		defer n.Close()
		if written, err := io.Copy(w, n); err != nil {
			errchan <- fmt.Errorf("error copying %s (%d bytes written): %v", filePath, written, err)
		}
		if err = mwriter.Close(); err != nil {
			errchan <- err
			return
		}
	}()
	client := &http.Client{}
	resp, err := client.Do(request)
	merr := <-errchan

	if err != nil || merr != nil {
		return nil, fmt.Errorf("http error: %v, multipart error: %v", err, merr)
	}
	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return
}
