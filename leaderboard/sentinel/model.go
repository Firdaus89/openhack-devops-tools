package main

import "time"

type Team struct {
	Id       string
	Name     string
	Score    int
	Services []Service
}

type Challenge struct {
	Id        string
	StartDate time.Time
	EndDate   time.Time
	Histories []History
}

type History struct {
	Id        string
	ServiceId string
	Status    string
	Date      time.Time
}

type Service struct {
	Id   string
	Name string
	Uri  string
}
