package main

import (
	"errors"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Team struct {
	Id         string
	Name       string
	Score      int
	Services   *[]Service
	Challenges *[]Challenge
}

type Challenge struct {
	Id        string
	StartDate time.Time
	EndDate   time.Time
	Histories *[]History
}

type History struct {
	Id        string
	ServiceId string
	Status    string
	Date      time.Time
}

type Service struct {
	Id            string
	Name          string
	Uri           string
	CurrentStatus bool
	PokeClient    GetStatusClient
}

func (c *Team) GetCurrentChallenge() (*Challenge, error) {
	// Find a challenge which is Started but not finished.
	for _, v := range *c.Challenges {
		if !v.StartDate.IsZero() && v.EndDate.IsZero() {
			return &v, nil
		}
	}
	return nil, errors.New("Current Challenge not found. Please check the Challenge data which have StartDate but not EndDate")
}

func (c *Team) insertNewHistory(challengeId string, serviceId string, status string, date time.Time) {

	newChallenges := *c.Challenges
	for i, v := range newChallenges {
		if v.Id == challengeId {
			// insert
			var newHistories []History
			if v.Histories == nil {
				newHistories = []History{
					History{
						Id:        uuid.New().String(),
						ServiceId: serviceId,
						Status:    status,
						Date:      date,
					},
				}
			} else {
				newHistories = append(*v.Histories, History{
					Id:        uuid.New().String(),
					ServiceId: serviceId,
					Status:    status,
					Date:      date,
				})
			}
			newChallenges[i].Histories = &newHistories

		}
	}
	c.Challenges = &newChallenges
}

func statusConverter(status bool) string {
	if status == true {
		return "Alive"
	} else {
		return "Dead"
	}
}

func (c *Team) StatusCheck() {
	// Get a current Challenge
	challenge, err := c.GetCurrentChallenge()
	if err != nil {
		log.Printf(err.Error()) // We need to consider if it should be panic.
		return
	}
	// Loop through Services and Health Check it.
	for _, s := range *c.Services {
		s.HealthCheck()
		_, hasHistory := challenge.GetLatestHistory(s.Id)
		if hasHistory == true {

		} else {
			c.insertNewHistory(challenge.Id, s.Id, statusConverter(s.CurrentStatus), time.Now())
		}
	}

	// Check if the history exists which is the same Histories.

	// Insert History if neccessary
}

func (c *Challenge) GetLatestHistory(serviceId string) (*History, bool) {
	if c.Histories == nil {
		return nil, false
	}
	sortedHistories := *c.Histories
	sort.Slice(sortedHistories, func(i, j int) bool {
		return (sortedHistories[i].Date.Sub(sortedHistories[j].Date) > 0)
	})
	c.Histories = &sortedHistories
	for i, v := range *c.Histories {
		if v.ServiceId == serviceId {
			return &(*c.Histories)[i], true
		}
	}
	return nil, false
}

type GetStatusClient func(uri string) (*http.Response, error)

func (c *Service) HealthCheck() bool {
	resp, err := c.PokeClient(c.Uri)
	c.CurrentStatus = false
	if err != nil {
		log.Printf(err.Error()) // Compromise: original log.Fatal(err)
		return false
	}

	if resp.StatusCode == 200 {
		c.CurrentStatus = true
		return true
	} else {
		return false
	}
}
func RestGetClientImpl(uri string) (*http.Response, error) {
	return http.Get(uri)
}

type IService interface {
	HealthCheck() bool
}
