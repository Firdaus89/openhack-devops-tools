package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const COSMOS_DB_TEST_NAME = "sadb"
const COSMOS_COLLECTION_TEST_NAME = "col"

func SetUpSampleTemplate() {
	// Create database and collection if not exists
	// Create a Initial Data Models
	// Clear exsiting Database
	teamdb := NewDB(COSMOS_DB_TEST_NAME, COSMOS_COLLECTION_TEST_NAME, GetConfig())
	teamdb.RemoveDB()
	teamdb = NewDB(COSMOS_DB_TEST_NAME, COSMOS_COLLECTION_TEST_NAME, GetConfig())
	// Team01 -> Team20
	flag := 0
	for i := 0; i < 20; i++ {
		team := &Team{
			Name:  fmt.Sprintf("Team%0d", i+1),
			Score: 0,
		}
		if flag == 0 {
			setupPattern01(team)
			team.StatusCheck()
			flag++
		} else if flag == 1 {
			setupPattern02(team)
			team.StatusCheck()
			team.StatusCheck()
			flag++
		} else if flag == 2 {
			setupPattern03(team)
			team.StatusCheck()
			team.StatusCheck()
			flag = 0
		}
		teamdb.Add(team)
	}

	// Print the created data
	fmt.Println("sentinel ---------> ")
	fmt.Println("Data Generated!")
	fmt.Println("")
	teams, err := teamdb.GetAll()
	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(teams, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))

}

// Initialize for production
func SetupTemplate() {
	teamdb := NewDB(COSMOS_DB_TEST_NAME, COSMOS_COLLECTION_TEST_NAME, GetConfig())
	teamdb.RemoveDB()
	teamdb = NewDB(COSMOS_DB_TEST_NAME, COSMOS_COLLECTION_TEST_NAME, GetConfig())
	for i := 0; i < 20; i++ {
		team := &Team{
			Name:  fmt.Sprintf("Team%0d", i+1),
			Score: 0,
		}
		teamdb.Add(team)
	}
}

// Challenge 2, Service 1
func setupPattern01(team *Team) {
	challenges := []Challenge{Challenge{
		Id:        "1",
		StartDate: time.Date(2018, 3, 10, 0, 0, 0, 0, time.Local),
		EndDate:   time.Date(2018, 3, 10, 0, 30, 0, 0, time.Local),
	},
		Challenge{
			Id:        "2",
			StartDate: time.Date(2018, 3, 10, 1, 0, 0, 0, time.Local),
		},
	}
	services := []Service{Service{
		Id:   "1",
		Name: "EP01",
		Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
	}}
	team.Challenges = &challenges
	team.Services = &services
}

// Challenge 5, Service 5
func setupPattern02(team *Team) {
	challenges := []Challenge{
		Challenge{
			Id:        "1",
			StartDate: time.Date(2018, 3, 10, 0, 0, 0, 0, time.Local),
			EndDate:   time.Date(2018, 3, 10, 0, 30, 0, 0, time.Local),
		},
		Challenge{
			Id:        "2",
			StartDate: time.Date(2018, 3, 10, 1, 0, 0, 0, time.Local),
			EndDate:   time.Date(2018, 3, 10, 1, 30, 0, 0, time.Local),
		},
		Challenge{
			Id:        "3",
			StartDate: time.Date(2018, 3, 10, 2, 0, 0, 0, time.Local),
			EndDate:   time.Date(2018, 3, 10, 2, 30, 0, 0, time.Local),
		},
		Challenge{
			Id:        "4",
			StartDate: time.Date(2018, 3, 10, 3, 0, 0, 0, time.Local),
			EndDate:   time.Date(2018, 3, 10, 3, 30, 0, 0, time.Local),
		},
		Challenge{
			Id:        "5",
			StartDate: time.Date(2018, 3, 10, 4, 0, 0, 0, time.Local),
		},
	}

	services := []Service{
		Service{
			Id:   "1",
			Name: "EP01",
			Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
		},
		Service{
			Id:   "2",
			Name: "EP02",
			Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
		},
		Service{
			Id:   "3",
			Name: "EP03",
			Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
		},
		Service{
			Id:   "4",
			Name: "EP04",
			Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
		},
		Service{
			Id:   "5",
			Name: "EP05",
			Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
		},
	}
	team.Challenges = &challenges
	team.Services = &services
}

// Challenge 3, Service 2
func setupPattern03(team *Team) {
	challenges := []Challenge{
		Challenge{
			Id:        "1",
			StartDate: time.Date(2018, 3, 10, 0, 0, 0, 0, time.Local),
			EndDate:   time.Date(2018, 3, 10, 0, 30, 0, 0, time.Local),
		},
		Challenge{
			Id:        "2",
			StartDate: time.Date(2018, 3, 10, 1, 0, 0, 0, time.Local),
			EndDate:   time.Date(2018, 3, 10, 1, 30, 0, 0, time.Local),
		},
		Challenge{
			Id:        "3",
			StartDate: time.Date(2018, 3, 10, 2, 0, 0, 0, time.Local),
		},
	}

	services := []Service{
		Service{
			Id:   "1",
			Name: "EP01",
			Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
		},
		Service{
			Id:   "2",
			Name: "EP02",
			Uri:  "https://sarmopenhack.azurewebsites.net/api/team01/health",
		},
	}
	team.Challenges = &challenges
	team.Services = &services
}
