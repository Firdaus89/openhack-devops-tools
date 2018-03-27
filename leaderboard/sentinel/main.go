package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

// If an defacto environment has not been set, cause an error with some messages.
func main() {
	app := cli.NewApp()
	app.Name = "Sentinel"
	app.Usage = "Status checking tool for OpenHack - DevOps"
	//	app.UsageText = "This is the usage text" // default will be nice
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize Sentinel. Insert initial Data for an Open Hack",
			Action: func(c *cli.Context) error {
				fmt.Println("***** I do initialize!")
				Initialize(c.Bool("t"))
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "template, t",
					Usage: "Initialize template data set for Testing",
				},
			},
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start Sentinel App and monitor the endpoint",
			Action: func(c *cli.Context) error {
				StartJobs()
				return nil
			},
		},
	}
	a := os.Args
	app.Run(a)
}

// Initialize initialize the Database settings. If the template is true, it inserts template data for testing.
func Initialize(template bool) {
	if template {
		SetUpSampleTemplate()
	} else {
		SetupTemplate()
	}
}

// StartJobs start Sentinel orchestration jobs
func StartJobs() {
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
			fmt.Println("********* Tick at", t)
			RunAllPokers()
		}
	}()
	time.Sleep(30 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker Stopped")
}
