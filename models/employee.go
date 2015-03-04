package models

import "time"

type Employee struct {
	EmployeeNo int
	FirstName  string
	LastName   string
	Gender     string
	BirthDate  time.Time
}
