package license

import (
	"encoding/json"
	"fmt"
	"log"
	"minos/pkg/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var LicKey = "lasj!kjl23ads-bd"

type Claims struct {
	Comp    string `json:"comp"`
	Dest    string `json:"dest"`
	ExpDate string `json:"exp_date"`
	jwt.StandardClaims
}

func LicGen(sub string, expired int64, hostids, services []string) (string, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(expired) * time.Second)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//claims["name"] = "xxx"
	//jwt.StandardClaims{}
	if expired > 0 {
		claims["exp"] = expiresAt.Unix()
		var timeFmt = expiresAt.Format("2006-01-02 15:04:05")
		claims["exp_date"] = timeFmt
	} else {
		delete(claims, "exp")
	}
	claims["iat"] = now.Unix() // 签发人
	claims["iss"] = "air jwt."
	claims["sub"] = sub
	//claims["jti"] = "go jwt."  // id
	svcs, _ := json.Marshal(services)
	bytesData, err := json.Marshal(hostids)
	if err != nil {
		panic(err)
	}
	encHostStr := utils.EnTxtByAes(string(bytesData), LicKey)
	claims["dest"] = encHostStr
	claims["services"] = string(svcs)

	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(LicKey))
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	//fmt.Println("token:", tokenStr)
	return tokenStr, nil
}

func LicParse(token string) (bool, error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("许可错误")
		}
		return []byte(LicKey), nil
	})
	//fmt.Println("token:", token, len(keys), keys)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return false, fmt.Errorf("服务许可错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				tt, _, er := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
				if er != nil {
					log.Println("claims er:", er)
				}

				claims, okk := tt.Claims.(jwt.MapClaims)
				if okk {
					timeS := time.Unix(int64(claims["iat"].(float64)), 0)
					//fmt.Println("tt:", tt, timeS.Format("2006-01-02 15:04:05"))
					startTime := timeS.Format("2006-01-02 15:04:05")
					err = fmt.Errorf("服务许可已经过期: %s 至 %s", startTime, claims["exp_date"])
				} else {
					err = fmt.Errorf("服务许可已经过期")
				}
				return false, err
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return false, fmt.Errorf("服务许可错误")
			} else {
				return false, fmt.Errorf("服务许可错误")
			}
		}
		return false, fmt.Errorf("服务许可错误")
	} else if !parseToken.Valid {
		return false, fmt.Errorf("服务许可错误")
	}

	//fmt.Println("token:", parseToken)
	decHostStr := utils.DeTxtByAes(parseToken.Claims.(jwt.MapClaims)["dest"].(string), LicKey)
	//fmt.Println("decHostStr:", decHostStr)

	if strings.Contains(decHostStr, GetHostId()) {
		return true, nil
	} else {
		return false, fmt.Errorf("该主机服务许可未授权")
	}
}

type Lic struct {
	License   string `json:"license"`
	ErrMsg    string `json:"err_msg"`
	Dest      string `json:"dest"`
	StartTime string `json:"start_time"`
	ExpDate   string `json:"exp_date"`
	jwt.StandardClaims
}

func LicInfo(token string) Lic {
	li := Lic{License: token}
	tt, _, er := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if er != nil {
		log.Println("claims er:", er)
		li.ErrMsg = fmt.Sprintf("无效的许可：%s", er.Error())
		return li
	}

	claims, okk := tt.Claims.(jwt.MapClaims)
	if okk {
		timeS := time.Unix(int64(claims["iat"].(float64)), 0)
		marshal, _ := json.Marshal(claims)
		_ = json.Unmarshal(marshal, &li)
		//fmt.Println("tt:", tt, timeS.Format("2006-01-02 15:04:05"))
		li.StartTime = timeS.Format("2006-01-02 15:04:05")
	} else {
		li.ErrMsg = fmt.Sprintf("获取数据失败")
	}
	return li
}
