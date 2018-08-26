package trace_ctl

import (
	"github.com/globalsign/mgo"
	conf "quan/config"
	"fmt")

var Session *mgo.Session 

func init() {
	host := conf.Conf.GetValue("database", "host")
	port := conf.Conf.GetValue("database", "port")
	db := conf.Conf.GetValue("database", "db")
	url := fmt.Sprintf("mongodb://%s:%s/%s", host, port, db)
	Session, _ = mgo.Dial(url)
}
