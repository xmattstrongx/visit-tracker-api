package models

type City struct {
	Name   string `json:"name"`
	County string `json:"county"`
	State  string `json:"state"`
}

type CityDetails struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	StateID       string `json:"stateid"`
	County        string `json:"county"`
	Status        string `json:"-"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	DateAdded     string `json:"dateadded"`
	DateTimeAdded string `json:"datetimeadded"`
	LastUpdated   string `json:"lastupdated"`
}
