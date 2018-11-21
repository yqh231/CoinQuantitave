package main

import (
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
  )
  
func main() {
	switch kingpin.MustParse(monitor.Parse(os.Args[1:])) {

	case bp.FullCommand():
			bpCommand()	
	case market.FullCommand():
			marketCommand()
	}

}