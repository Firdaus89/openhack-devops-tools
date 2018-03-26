package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

// RunAllPokers pokes all services and update the status
func RunAllPokers() {
	// Read all Teams
	teams := ReadAllTeams()
	// Parallel execution of Check Status per each Team
	// Wait for all finished
	newTeams := ExecStatusCheckAsync(teams)
	// Update the status
	err := UpdateTeamStatus(newTeams)
	if err != nil {
		log.Printf(fmt.Sprint("UpdateTeamStatus failed: "))
		log.Printf(err.Error())
	}
}

// ExecStatusCheckAsync execute StatusCheck() for all Teams. It will be sored by Team.Name
func ExecStatusCheckAsync(teams *[]Team) *[]Team {
	ch := make(chan Team)
	for _, v := range *teams {
		go func(team Team) {
			team.StatusCheck()
			ch <- team
		}(v)
	}
	var newTeams []Team
	for i := 0; i < len(*teams); i++ {
		result := <-ch
		newTeams = append(newTeams, result)
	}
	sort.Slice(newTeams, func(i, j int) bool {
		return (newTeams[i].Name < newTeams[j].Name)
	})
	return &newTeams
}

const COSMOS_DB_NAME = "sadb"
const COSMOS_COLLECTION_NAME = "col"

// ReadAllTeams reads from MongoDB the Team, Challenge, Service, History
func ReadAllTeams() *[]Team {
	// Read Teams from MongoDB
	teamdb := NewDB(COSMOS_DB_NAME, COSMOS_COLLECTION_NAME, GetConfig())
	teams, err := teamdb.GetAll()
	if err != nil {
		log.Fatalf("Team DB GetAll() error! %s", err.Error())
	}
	newTeams := *teams
	sort.Slice(newTeams, func(i, j int) bool {
		return (newTeams[i].Name < newTeams[j].Name)
	})
	return &newTeams
}

// UpdateTeamStatus updates MongoDB's team Status
func UpdateTeamStatus(teams *[]Team) error {
	// Write Teams for MongoDB
	teamdb := NewDB(COSMOS_DB_NAME, COSMOS_COLLECTION_NAME, GetConfig())
	for i, v := range *teams {
		err := teamdb.Add(&v)
		if err != nil {
			fmt.Printf("upsert error! %d: '%s'", i, err.Error())

		}
	}
	return nil
}

func GetConfig() *Config {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	config := &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		panic(err)
	}
	return config
}
