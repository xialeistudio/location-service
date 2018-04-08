package driver

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

type AMap struct {
	Requester
	AK      string
	baseUrl string
}

func NewAMap(ak string) *AMap {
	return &AMap{
		AK:      ak,
		baseUrl: "http://restapi.amap.com/v3",
	}
}

func (m AMap) LocationWithIp(ip string) (*Location, error) {
	link := m.baseUrl + "/ip"
	params := &url.Values{
		"key":    []string{m.AK},
		"ip":     []string{ip},
		"output": []string{"json"},
	}
	data, err := m.Request(link, params)
	if err != nil {
		return nil, err
	}

	mapData := make(map[string]string)
	if err := json.Unmarshal(data, &mapData); err != nil {
		return nil, ErrInvalidResponse
	}
	if mapData["status"] == "0" {
		return nil, errors.New(mapData["infocode"] + ":" + mapData["info"])
	}
	code, err := strconv.Atoi(mapData["adcode"])
	if err != nil {
		return nil, ErrInvalidResponse
	}
	return &Location{
		Country:  "中国",
		Province: mapData["province"],
		City:     mapData["city"],
		Code:     code,
	}, nil
}

func (m AMap) LocationWithAddress(address string) (*Location, error) {
	link := m.baseUrl + "/geocode/geo"
	params := &url.Values{
		"key":     []string{m.AK},
		"address": []string{address},
		"output":  []string{"json"},
	}
	data, err := m.Request(link, params)
	if err != nil {
		return nil, err
	}

	mapData := make(map[string]interface{})
	if err := json.Unmarshal(data, &mapData); err != nil {
		return nil, ErrInvalidResponse
	}
	if mapData["status"] == "0" {
		return nil, errors.New(mapData["infocode"].(string) + ":" + mapData["info"].(string))
	}

	geocodes, ok := mapData["geocodes"].([]interface{})
	if !ok {
		return nil, ErrInvalidResponse
	}
	if len(geocodes) == 0 {
		return nil, ErrAddressNotExists
	}
	geocode, ok := geocodes[0].(map[string]interface{})
	if !ok {
		return nil, ErrInvalidResponse
	}
	code, err := strconv.Atoi(geocode["adcode"].(string))
	if err != nil {
		return nil, ErrInvalidResponse
	}
	return &Location{
		Country:  "中国",
		Province: geocode["province"].(string),
		City:     geocode["city"].(string),
		District: geocode["district"].(string),
		Code:     code,
	}, nil
}

func (m AMap) LocationWithCoordinate(point string) (*Location, error) {
	link := m.baseUrl + "/geocode/regeo"
	params := &url.Values{
		"key":      []string{m.AK},
		"location": []string{point},
		"output":   []string{"json"},
	}
	data, err := m.Request(link, params)
	if err != nil {
		return nil, err
	}

	mapData := make(map[string]interface{})
	if err := json.Unmarshal(data, &mapData); err != nil {
		return nil, ErrInvalidResponse
	}
	if mapData["status"] == "0" {
		return nil, errors.New(mapData["infocode"].(string) + ":" + mapData["info"].(string))
	}

	regeocode, ok := mapData["regeocode"].(map[string]interface{})
	if !ok {
		return nil, ErrInvalidResponse
	}
	address, ok := regeocode["addressComponent"].(map[string]interface{})
	if !ok {
		return nil, ErrInvalidResponse
	}
	code, err := strconv.Atoi(address["adcode"].(string))
	if err != nil {
		return nil, ErrInvalidResponse
	}
	return &Location{
		Country:  address["country"].(string),
		Province: address["province"].(string),
		City:     address["city"].(string),
		District: address["district"].(string),
		Code:     code,
	}, nil
}
