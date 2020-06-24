package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	target := request.Form["target"]
	if len(target) > 0 {
		resp, err := http.Get(target[0])
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
