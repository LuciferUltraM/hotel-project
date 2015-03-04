package models

import "time"

type Customer struct {
	FirstName        string
	LastName         string
	Tel              string
	Nationality      string
	BirthDate        time.Time
	Gender           string
	IDCardPassportNo string
}
