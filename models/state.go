package models

type State struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type StateDetails struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Abbreviation  string `json:"abbreviation"`
	DateAdded     string `json:"dateadded"`
	DateTimeAdded string `json:"datetimeadded"`
	LastUpdated   string `json:"lastupdated"`
}
