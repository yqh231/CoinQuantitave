package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	monitor = kingpin.New("monitor client", "A command-line for monitor server")

	bp     = monitor.Command("bp", "Get bp list")
	bpMarket = bp.Flag("market", "Specify break point market").String()
	bpFlag   = bp.Flag("flag", "Flag for break record").String()
	bpLimit  = bp.Flag("limit", "Limit to reveal").String()
	bpOffset = bp.Flag("offset", "Offset for current page").String()

	market   = monitor.Command("market", "Get market list")
	marketAdd = market.Flag("add", "Add market").String()
	marketList = market.Flag("list", "List market").String()
)

const (
	CodeOk = 0
)

func bpCommand(){
	filter := make(map[string]string)
	if *bpMarket != ""{
		filter["market"] = *bpMarket
	}
	
	if *bpFlag != ""{
		filter["flag"] = *bpFlag
	}

	if *bpLimit != ""{
		filter["limit"] = *bpLimit
	}

	if *bpOffset != ""{
		filter["offset"] = *bpOffset
	}
	bp := bpList(filter)
	if bp.Data == nil || bp.Code != CodeOk{
		err()
		return
	}
	bpSlice := toSlice(bp.Data)
	
	Fprint("CurPrice", "TopPrice", "Market", "CreateTime")
	for _, item := range bpSlice{
		v := item.(map[string]interface{})
		if v["Flag"] == true{
			RedFprint(toString(v["CurPrice"]), toString(v["TopPrice"]), toString(v["Market"]), toString(v["CreateTime"]))
			continue
		}
		Fprint(toString(v["CurPrice"]), toString(v["TopPrice"]), toString(v["Market"]), toString(v["CreateTime"]))
	}

}


func marketCommand(){
	if *marketAdd != ""{
		add := addMarket(*marketAdd)	
		if add.Code != CodeOk{
			goto errs
		}
		Fprint("Add market success!")
		return
	}else if *marketList != ""{
		markets := listMarket()	
		if markets.Code != CodeOk{
			goto errs
		}
		Fprint(toStringSlice(markets.Data)...)
		return
	}

errs:
	err()
	return

}