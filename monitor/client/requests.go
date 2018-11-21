package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"io/ioutil"
	"bytes"
)

type Response struct{
	Code int
	Data interface{}
}

type BpList struct{
	CurPrice string
	TopPrice string
	Market string
	CreateTime string
	Flag bool 
}

const Endpoint = "http://127.0.0.1:8080/"

var client = &http.Client{}

func DoRequest(req *http.Request) []byte{
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func bpList(filter map[string]string) *Response{
	response := new(Response)

	response.Data = make([]BpList, 1)

	new, _ := http.NewRequest("GET", Endpoint+"market/bp", nil)

	q := new.URL.Query()

	if market, ok := filter["market"];ok{
		q.Add("market", market)
	}

	if flag, ok := filter["flag"];ok{
		q.Add("flag", flag)
	}

	if limit, ok := filter["limit"];ok{
		q.Add("limit", limit)
	}

	if offset, ok := filter["offset"];ok{
		q.Add("offset", offset)
	}
	new.URL.RawQuery = q.Encode()	

	resp := DoRequest(new)

	json.Unmarshal(resp, response)

	return response
}

func addMarket(market string) *Response{
	response := new(Response)
	data := url.Values{}
	data.Set("market", market)
	new, _ := http.NewRequest("POST", Endpoint+"market/add", bytes.NewBufferString(data.Encode()))
	new.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	resp := DoRequest(new)

	json.Unmarshal(resp, response)
	return response
}

func listMarket() *Response{
	response := new(Response)

	response.Data = make([]string, 1)

	new, _ := http.NewRequest("GET", Endpoint+"market/list", nil)
	resp := DoRequest(new)

	json.Unmarshal(resp, response)

	return response
}