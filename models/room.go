package models

type Room struct {
	RoomNo    string
	RoomFloor string
	Status    string
	RoomType  *RoomType
}
