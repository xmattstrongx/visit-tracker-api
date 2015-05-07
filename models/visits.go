package models

type Visit struct {
	City         string `json:"city"`
	Abbreviation string `json:"state"`
}

type VisitDetails struct {
	ID            string `json:"id"`
	UserID        string `json:"userid"`
	City          string `json:"city"`
	Abbreviation  string `json:"state"`
	DateTimeAdded string `json:"datetimeadded"`
}
