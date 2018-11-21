package main

import (
	"encoding/json"
	"reflect"
	"bytes"
)


func toSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	l := v.Len()
	ret := make([]interface{}, l)
	for i:=0; i<l; i++{
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func toString(ins interface{}) string{
	return ins.(string)
}

func toPostBody(body map[string]interface{}) *bytes.Buffer{
	bytesRepresentation, _ := json.Marshal(body)
	return bytes.NewBuffer(bytesRepresentation)
} 

func toStringSlice(arr interface{}) (res []string){
	r := toSlice(arr)

	for _, s := range r{
		res = append(res, toString(s))
	}
	return
}