package trx

import ("net/http"
	"time"
	"io/ioutil"
		)
/*var api = map[string]string{
	"auth": "xxxxx"}


func genSign(){
	api_key := "4A4928DDD06A4F14BDD91C3EA2998C72"
	api_secret := "232CC50F6BF34E5BA2E336937AC282BEC6949F2BDD8995AE"
}*/

var Client *http.Client

type CoinexRequest struct{
	req *http.Request
}

func init(){
	tr := &http.Transport{
		MaxIdleConns: 10,
		IdleConnTimeout: 30 * time.Second,
		DisableCompression: true,
	}

	Client = &http.Client{Transport: tr}
}

func(c *CoinexRequest) SetHeader(){
	c.req.Header.Add("Content-Type", "application/json; charset=utf-8")
	c.req.Header.Add("Accept", "application/json")
	c.req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
}

func(c *CoinexRequest) DoRequest() []byte{
	resp, err := Client.Do(c.req)
	if err != nil{

	}
	body, _:= ioutil.ReadAll(resp.Body)
	return body
}
