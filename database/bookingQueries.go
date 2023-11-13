package database
import(
	"github.com/gocql/gocql"
	"Theatre/models"
	"fmt"
	"log"
	"github.com/google/uuid"
	
)
func CreateBooking(session *gocql.Session,booking models.Booking)(string,error){
	ShowId:=booking.ShowId
	//Gets current show
	show,err := GetShowDetails(session,ShowId)
	booking.TotalAmount = show.Price * booking.RequiredSeats
	fmt.Printf("%v\n",show)
	if err != nil {
		fmt.Printf("%v",err)
        return "",err
    }
	// Update seat positions to unavailable
	for i, rows := range booking.SeatPositions {
        for j,value:= range rows{
			if !value{
				show.SeatInTotal[i][j] = false
			}
		}
    }
	// Update the show details
	_,err= DeleteShowById(session,show.ShowId)
	if(err!=nil){
		
		return " ",err
		}else{			
			_,err := InsertShow(session, *show)
			if(err!=nil){
					return " ",err
				}
			
		}
	// Create new booking
	id := uuid.New()
    BookingId:=id.String()
	insertStatement := fmt.Sprintf("INSERT INTO %s.%s ( BookingId,DateOfShow,RequiredSeats,SeatPositions,ShowId,ShowTiming,TotalAmount,UserId) VALUES (?, ?, ?, ?, ?,?,?,?)IF NOT EXISTS","ALGOMOX_MULTIPLEX_BOOKING","Booking")
	if err := session.Query(
		insertStatement,
		BookingId,
		booking.DateOfShow,
		booking.RequiredSeats,
		booking.SeatPositions,
		booking.ShowId,
		booking.ShowTiming,
		booking.TotalAmount,
		booking.UserId,
		).Exec(); err != nil {
			log.Printf("Error inserting data into table: %v\n", err)
			
			return " ",err
		}
	log.Println("Data inserted successfully.")
	return BookingId,nil
}
func PreviousHistory(session *gocql.Session,username string)([]models.Booking,error){
	query := "SELECT  *  FROM ALGOMOX_MULTIPLEX_BOOKING.Booking WHERE 	UserId = ? ALLOW FILTERING"
    iter := session.Query(query,username).Iter()
    var booking []models.Booking
    var book models.Booking
    for iter.Scan(
		&book.BookingId,
		&book.DateOfShow,
		&book.RequiredSeats,
		&book.SeatPositions,
		&book.ShowId,
		&book.ShowTiming,
		&book.TotalAmount,
		&book.UserId,
		) {
        booking = append(booking, book)
        book = models.Booking{} 
    }

    if err := iter.Close(); err != nil {
        return nil, err
    }
    return booking, nil
}
