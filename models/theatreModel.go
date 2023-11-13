package models

//structure for Theatre 
type Theatre struct
{
	TheatreId string
	TheatreName string
	TotalSeats int
	SlotTiming []  string
	Rating float64
	TheatreOwnerId string
}