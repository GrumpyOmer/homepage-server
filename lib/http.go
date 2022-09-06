package lib

import (
	"io"
	"net/http"
	urls "net/url"
)

func Get(url string, header, data map[string]string) (string, error) {
	var (
		param = urls.Values{}
		res   string
		err   error
	)
	client := &http.Client{}
	// add param
	for k, v := range data {
		param.Add(k, v)
	}
	// add header
	req, err := http.NewRequest("GET", url+"?"+param.Encode(), nil)
	if err != nil {
		return res, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	rsp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	res = string(body)
	return res, err
}
