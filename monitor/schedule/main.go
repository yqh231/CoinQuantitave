package main

import(
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("@hourly", taskBreakPoint)
	c.Start()
	select{}
}