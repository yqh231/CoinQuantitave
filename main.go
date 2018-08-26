package main 

import (
	"fmt"
	_ "quan/config"
	_ "quan/trace_ctl"
	"quan/trx"
)

func StartTrade(){
	var h trx.Hedging

	h.Hedge()
}

func main() {
	var message chan string = make(chan string)
	go func(){
		for{
			StartTrade()
		}

		message <- "finish"
	}()	
	fmt.Println(<-message)
}