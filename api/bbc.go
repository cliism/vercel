package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://news.bbc.co.uk")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, string(body))
	}
}
