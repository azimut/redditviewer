package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetFromParam() (string, error) {
	client := &http.Client{
		Timeout: 2 * time.Second}
	s := os.Args[1]
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	if u.Host != "www.reddit.com" {
		return "", fmt.Errorf("not supported host")
	}
	s = s + ".json"
	// Request
	req, err := http.NewRequest(http.MethodGet, s, nil)
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
		fmt.Println(s)
		fmt.Println(string(r))
		return "", fmt.Errorf("invalid http status code %d", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}
