package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yqh231/CoinQuantitave/monitor/mongo"
	"sort"
	"net/http"
)

type Bp struct{
	CurPrice string 
	TopPrice string
	Market string
	CreateTime string
	Flag bool 
}

func breakPoint(c *gin.Context){
	market := c.DefaultQuery("market", "")
	flag := str2int64(c.DefaultQuery("flag", "0"))
	limit := str2int64(c.DefaultQuery("limit", "20"))
	offset := str2int64(c.DefaultQuery("offset", "0"))

	var filters []map[string]interface{}
	var results []Bp

	if !strIsNull(market){
		filters = append(filters, map[string]interface{}{"market": market})
	}
	if flag == 1{
		filters = append(filters, map[string]interface{}{"break_point": true})
	}
	bp := mongo.FilterBp(mongo.MakeDoc(filters))
	if bp == nil{
		c.JSON(http.StatusOK, ResOk("no data"))	
	}
	for i, r := range bp{
		if i < int(offset){
			continue
		}
		if len(results) >= int(limit){
			continue
		}
		results = append(results, Bp{
			CurPrice: r.Lookup("cur_price").StringValue(),
			TopPrice: r.Lookup("top_price").StringValue(),
			Market: r.Lookup("market").StringValue(),
			CreateTime: timestamp2time(r.Lookup("create_time").Int64()),
			Flag: r.Lookup("break_point").Boolean(),
		})
	}
	sort.SliceStable(results, func(i, j int)bool{
		return results[i].CreateTime > results[j].CreateTime
	})

	c.JSON(http.StatusOK, ResOk(results))
}