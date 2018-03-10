package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

// If an defacto environment has not been set, cause an error with some messages.
func main() {

	// Iterate an Call of the orchestrator

	// session, err := mgo.Dial("localhost")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// c := session.DB("test").C("people")
	// err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	// 	&Person{"Cla", "+53 53 8402 8510"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print("Phone:", result.Phone)

	fmt.Println("Sentinel - Health check daemon for the DevOps - OpenHack")
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for t := range ticker.C {
			RunAllPokers()
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(30 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker Stopped")
}
