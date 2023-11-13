package models

//structure for show model
type Show struct
{   Moviename string
	ShowTiming string 
	TheatreId string
	DateOfShow string
	ReservedSeats int 
	Availability bool
	SeatInTotal [][]bool
	ShowId string
	Price int
}