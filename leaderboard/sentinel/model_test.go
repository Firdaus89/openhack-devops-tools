package main_test

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	. "github.com/Azure-Samples/openhack-devops-tools/leaderboard/sentinel"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getTeamWithHistoriesNullSample(statusClient GetStatusClient) *Team {
	mockClient := statusClient

	aService := Service{
		Id:         "1",
		Name:       "A Service",
		Uri:        "http://aaa.azure.com/health",
		PokeClient: mockClient,
	}

	aChallenge := Challenge{
		Id:        "1",
		StartDate: time.Date(2019, 1, 9, 23, 59, 59, 0, time.Local),
	}
	aTeam := &Team{
		Id:         "1",
		Name:       "Team 1",
		Score:      100,
		Challenges: &[]Challenge{aChallenge},
		Services:   &[]Service{aService},
	}
	return aTeam
}

var _ = Describe("Team", func() {
	Context("Get a current challenge", func() {
		It("Challenge exists", func() {
			aTeam := &Team{
				Challenges: &[]Challenge{
					Challenge{
						Id:        "1",
						StartDate: time.Date(2019, 1, 9, 1, 10, 00, 0, time.Local),
						EndDate:   time.Date(2019, 1, 9, 2, 10, 00, 0, time.Local),
					},
					Challenge{
						Id:        "2",
						StartDate: time.Date(2019, 1, 9, 3, 10, 00, 0, time.Local),
					},
				},
			}
			currentChallenge, _ := aTeam.GetCurrentChallenge()
			Expect(currentChallenge.Id).To(Equal("2"))
		})
		It("Challenge doesn't exists", func() {
			aTeam := &Team{
				Challenges: &[]Challenge{
					Challenge{
						Id:        "1",
						StartDate: time.Date(2019, 1, 9, 1, 10, 00, 0, time.Local),
						EndDate:   time.Date(2019, 1, 9, 2, 10, 00, 0, time.Local),
					},
				},
			}
			_, err := aTeam.GetCurrentChallenge()
			Expect(fmt.Sprint(err)).To(Equal("Current Challenge not found. Please check the Challenge data which have StartDate but not EndDate"))
		})
	})

	Context("When the team has One Challenge", func() {
		Context("History is null", func() {
			Context("Service is Alive", func() {
				It("Should write a new history as Alive", func() {
					aTeam := getTeamWithHistoriesNullSample(
						func(uri string) (*http.Response, error) {
							return &http.Response{
								StatusCode: 200,
							}, nil
						})
					t := time.Now()
					aTeam.StatusCheck()
					result := *aTeam.Challenges
					// Expect(result[0].Histories[0].Id).To(Equal("SOME_UUID")) // Skip this for UUID generation
					Expect((*result[0].Histories)[0].ServiceId).To(Equal("1"))
					Expect((*result[0].Histories)[0].Status).To(Equal("Alive"))
					Expect((*result[0].Histories)[0].Date).To(Equal(t))

				})
			})
			Context("Service is Dead", func() {
				It("should write a new history as Dead", func() {

				})
			})
		})
	})
})

var _ = Describe("Challenge", func() {
	Context("When I have no history for the target service id ", func() {
		It("should be return false as hasHistory", func() {
			challenge := &Challenge{}
			_, hasHistory := challenge.GetLatestHistory("1")
			Expect(hasHistory).To(BeFalse())
			challenge = &Challenge{
				Histories: &[]History{
					History{
						Id: "2",
					},
				},
			}
			_, hasHistory = challenge.GetLatestHistory("1")
			Expect(hasHistory).To(BeFalse())
		})
	})
	Context("When I have a history for the target service id", func() {
		It("should return true as hasHistory and also return the History", func() {
			challenge := &Challenge{
				Histories: &[]History{
					History{
						Id:        "1",
						ServiceId: "1",
					},
				},
			}
			history, hasHistory := challenge.GetLatestHistory("1")
			Expect(hasHistory).To(BeTrue())
			Expect(history.ServiceId).To(Equal("1"))
		})
	})
	Context("When I have several history for the target service id", func() {
		It("should return true as hasHistory and also return the History", func() {
			challenge := &Challenge{
				Histories: &[]History{
					History{
						Id:        "2",
						ServiceId: "1",
						Date:      time.Date(2019, 1, 9, 1, 10, 00, 0, time.Local),
					},
					History{
						Id:        "1",
						ServiceId: "1",
						Date:      time.Date(2019, 1, 9, 1, 10, 10, 0, time.Local),
					},
				},
			}
			history, hasHistory := challenge.GetLatestHistory("1")
			Expect(hasHistory).To(BeTrue())
			Expect(history.Id).To(Equal("1"))
		})
	})
})

var _ = Describe("Service", func() {
	Context("When I receive StatusCode 200", func() {
		It("returns true and set it to the current status", func() {

			parameter := ""
			mockClient := func(uri string) (*http.Response, error) {
				parameter = uri
				return &http.Response{
					StatusCode: 200,
				}, nil
			}
			service := &Service{
				Id:         "1",
				Name:       "Team1Service",
				Uri:        "http://aaa.azure.com/health",
				PokeClient: mockClient,
			}
			Expect(service.HealthCheck()).To(BeTrue())
			Expect(parameter).To(Equal("http://aaa.azure.com/health"))
			Expect(service.CurrentStatus).To(BeTrue())
		})
	})

	Context("When I receive StatusCode 401", func() {
		It("returns false and set currentStatus as false", func() {
			parameter := ""
			mockClient := func(uri string) (*http.Response, error) {
				parameter = uri
				return &http.Response{
					StatusCode: 401,
				}, nil
			}
			service := &Service{
				Id:         "1",
				Name:       "Team1Service",
				Uri:        "http://aaa.azure.com/health",
				PokeClient: mockClient,
			}
			Expect(service.HealthCheck()).To(BeFalse())
			Expect(parameter).To(Equal("http://aaa.azure.com/health"))
			Expect(service.CurrentStatus).To(BeFalse())
		})
	})

	Context("When I've got an error", func() {
		It("fails and set CurrentStatus as False", func() {
			parameter := ""
			mockClient := func(uri string) (*http.Response, error) {
				parameter = uri
				return nil, errors.New("Authorization error")
			}
			service := &Service{
				Id:         "1",
				Name:       "Team1Service",
				Uri:        "http://aaa.azure.com/health",
				PokeClient: mockClient,
			}
			Expect(service.HealthCheck()).To(BeFalse())
			Expect(parameter).To(Equal("http://aaa.azure.com/health"))
			Expect(service.CurrentStatus).To(BeFalse())
		})
	})

})
