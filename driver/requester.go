package driver

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Requester struct{}

func (Requester) Request(link string, params *url.Values) ([]byte, error) {
	resp, err := http.Get(link + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
