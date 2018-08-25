package trace_ctl

import "github.com/globalsign/mgo"

var Session *mgo.Session 

func init() {
	url := "mongodb://yqhdev:27000/trx"
	Session, _ = mgo.Dial(url)
}
