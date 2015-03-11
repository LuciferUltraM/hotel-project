package models

type Room struct {
	RoomNo   string
	Floor    string
	Status   string
	RoomType *RoomType
}
