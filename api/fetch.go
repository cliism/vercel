package handler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API入口函数
func GetHandler(response http.ResponseWriter, request *http.Request) {
	target := target(request)
	if target != "" {
		resp, err := http.Get(target)
		if err != nil {
			fmt.Fprintf(response, err.Error())
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			response.Write(body)
		}
	} else {
		fmt.Fprintf(response, "404")
	}
}

// 获取目标链接
func target(request *http.Request) string {
	err := request.ParseForm()
	if err != nil {
		return ""
	}
	target := request.Form["target"]
	if len(target) > 0 {
		b64 := target[0]
		ret, err := base64.StdEncoding.DecodeString(b64)
		if err != nil {
			return ""
		} else {
			return string(ret)
		}
	} else {
		return ""
	}

}
