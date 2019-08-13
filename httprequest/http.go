package httprequest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func Post(tokenvalue, url, reqdata string) (string, error) {
	const TokenName = `Authorization`
	const tokenPrefix = `BEARER `
	const httptype = "POST"

	if len(strings.TrimSpace(url)) == 0 {
		err := errors.New("url is null.")
		return "", err
	}
	client := &http.Client{}
	req, err := http.NewRequest(httptype, url, strings.NewReader(reqdata))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(TokenName, tokenPrefix+tokenvalue)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
