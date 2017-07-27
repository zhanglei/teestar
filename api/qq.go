package api

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

func HasQQUser(qq string) bool {
	resp, err := http.Get("http://www.webxml.com.cn/webservices/qqOnlineWebService.asmx/qqCheckOnline?qqCode=" + qq)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var res string
	xml.Unmarshal(body, &res)
	// Y = 在线
	// N = 离线
	// E = QQ号码错误
	// A = 商业用户验证失败
	// V = 免费用户超过数量

	if res == "Y" || res == "N" {
		return true
	} else if res == "E" {
		return false
	} else {
		panic(errors.New("Invalid QQ check reponse: " + res))
	}

	return false
}
