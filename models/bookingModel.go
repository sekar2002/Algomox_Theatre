package models
//structure for booking
type Booking struct{
	BookingId string
	ShowId string
	RequiredSeats int
	ShowTiming string
	DateOfShow string
	TotalAmount int
	SeatPositions [][]bool
	UserId string
}