package main

import (
	"sync"
)

// RunAllPokers pokes all services and update the status
func RunAllPokers() {
	// Read all Teams
	c := GetTeamsChanel(ReadAllTeams())
	// Parallel execution of Check Status per each Team
	// Wait for all finished
	newTeamsChanel := ExecAsync(c)
	// Update the status
	UpdateTeamStatus(newTeamsChanel)
}

func GetTeamsChanel(teams *[]Team) <-chan Team {
	out := make(chan Team)
	go func() {
		for _, n := range *teams {
			out <- n
		}
		close(out)
	}()
	return out
}

func ExecAsync(teams ...<-chan Team) <-chan Team {
	var wg sync.WaitGroup
	out := make(chan Team)
	output := func(ts <-chan Team) {
		for t := range ts {
			t.StatusCheck()
			out <- t
		}
		wg.Done()
	}
	wg.Add(len(teams))
	for _, t := range teams {
		go output(t)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// ReadAllTeams reads from MongoDB the Team, Challenge, Service, History
func ReadAllTeams() *[]Team {
	// Read Teams from MongoDB
	return nil
}

// UpdateTeamStatus updates MongoDB's team Status
func UpdateTeamStatus(teams ...<-chan Team) error {
	// Write Teams for MongoDB
	return nil
}
