package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type LogDB struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

func (db *LogDB) Close() {
	db.Session.Close()
}

func (db *LogDB) Insert(log *Log) error {
	return db.Collection.Insert(log)
}

type Log struct {
	TeamId    string
	ServiceId string
	Date      time.Time
}

func NewLogDB(cfg *config) (*LogDB, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%d", (*cfg).Addrs)},
		Timeout:  60 * time.Second,
		Database: (*cfg).Database,
		Username: (*cfg).Username,
		Password: (*cfg).Password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Printf("Can not connect to mongodb, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	collection := session.DB((*cfg).Database).C((*cfg).Collection)
	return &LogDB{
		Session:    session,
		Collection: collection,
	}, nil
}
