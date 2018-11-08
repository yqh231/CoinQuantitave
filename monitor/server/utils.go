package main

import(
	"time"
	"strconv"
	l "github.com/yqh231/CoinQuantitave/monitor/log"
)

type Response struct{
	Code int
	Data interface{}
}

func strIsNull(str string) bool{
	if str == ""{
		return true
	}
	return false
}

func boolIsNull(str string) bool{
	if str == "true"{
		return true
	}
	return false
}

func timestamp2time(tm int64) string{
	t := time.Unix(tm, 0)
	return t.Format("2006-01-02 15:04:05")
}

func str2int64(str string) int64{
	Int64, err := strconv.ParseInt(str, 10, 64)
	if err != nil{
		l.Logger.Println(err.Error())
	}
	return Int64
}

func ResOk(res interface{}) *Response{
	r := new(Response)
	r.Code = 0
	r.Data = res
	return r
}

func ResError(res interface{}) *Response{
	r := new(Response)
	r.Code = -1
	r.Data = nil 
	return r
}