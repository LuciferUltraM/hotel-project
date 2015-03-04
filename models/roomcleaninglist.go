package models

import "time"

type RoomCleaningList struct {
	ID         int
	Status     string
	Detail     string
	CreateDate time.Time
	UpdateDate time.Time
	AssignBy   *Receptionist
	CleanUpBy  *CleaningStaff
}
