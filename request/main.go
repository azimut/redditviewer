package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func GetFromParam(timeout int, uri string) (string, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second}

	if uri == "" {
		return "", fmt.Errorf("-u parameter not provided")
	}

	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	if u.Host != "www.reddit.com" {
		return "", fmt.Errorf("not supported host")
	}
	uri = uri + ".json"
	// Request
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Reddit_Cli/0.1")
	// Response
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	//
	if resp.StatusCode != http.StatusOK {
		r, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(uri)
		fmt.Println(string(r))
		return "", fmt.Errorf("invalid http status code %d", resp.StatusCode)
	}
	if b, err := ioutil.ReadAll(resp.Body); err == nil {
		return string(b), nil
	}
	return "", fmt.Errorf("no body read")
}
