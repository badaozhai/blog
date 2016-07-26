package tool

import (
	"io/ioutil"
	"net/http"
)

func GetHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	if err == nil {
		return s
	}
	return ""
}