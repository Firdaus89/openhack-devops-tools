package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/caarlos0/env"
)

type config struct {
	Endpoint      string `env:"SENTINEL_ENDPOINT,required"`
	PORT          int    `env:"SENTINEL_PORT" envDefault:"80"`
	TeamID        string `env:"SENTINEL_TEAM_ID, required"`
	ServiceID     string `env:"SENTINEL_SERVICE_ID,required"`
	Addrs         string `env:"SENTINEL_MONGO_ADDRESS,required"`
	Database      string `env:"SENTINEL_MONGO_DATABASE" envDefault:"sentineldb"`
	Username      string `env:"SENTINEL_MONGO_USERNAME,required"`
	Password      string `env:"SENTINEL_MONGO_PASSWORD,required"`
	Collection    string `env:"SENTINEL_MONGO_COLLECTION_NAME" envDefault:"collection"`
	Interval      int    `env:"SENTINEL_POLLING_INTERVAL" envDefault:"1"`
	RetryDuration int    `env:"SENTINEL_RETRY_DURATION" envDefault:"1000"`
}

func main() {
	// Get enviornment variables and validate it
	fmt.Println("Sentinel - Monitor an endpoint status for Openhack - DevOps")
	fmt.Println("\nStarting...")
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Println("Initialize Error")
		fmt.Printf("%+v\n", err)
		return
	}

	db, err := NewLogDB(&cfg)
	if err != nil {
		fmt.Println("Database Connection Error")
		fmt.Println("%+v\n", err)
		return
	}
	defer db.Close()

	ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	for t := range ticker.C {
		fmt.Println("Tick ...", t)
		statusCode, err := HelathCheck(&cfg)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Server: %s Status: %d", cfg.Endpoint, statusCode))
		// Endpoint is dead
		if statusCode != 200 {
			err = db.Insert(&Log{
				TeamId:    cfg.TeamID,
				ServiceId: cfg.ServiceID,
				Date:      time.Now(),
			})
			if err != nil {
				panic(err)
			}
			fmt.Println("Dead! wait for recovery for 1000 ms")
			time.Sleep(time.Duration(cfg.RetryDuration) * time.Millisecond)
		}
	}

	fmt.Println("Hello")
}

func HelathCheck(cfg *config) (int, error) {
	res, err := http.Get((*cfg).Endpoint)
	return res.StatusCode, err
}
