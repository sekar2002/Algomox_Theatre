package database
import (
	"fmt"
	"github.com/gocql/gocql"
)
//creating User Table
func CreateTableUser(session *gocql.Session, keyspace string) {
	createUserTableQuery := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s.Users (
            MailId TEXT PRIMARY KEY,
			Name TEXT,
			Password TEXT,
			Location TEXT
        )
    `, keyspace)
	if err := session.Query(createUserTableQuery).Exec(); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
	} else {
		fmt.Printf("Table created successfully in keyspace %s\n", keyspace)
	}
}
 //Creating Theatre table

 func CreateTableTheatre(session *gocql.Session, keyspace string) {
	createTheatreQuery := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s.Theatre (
        TheatreId UUID PRIMARY KEY,
        TotalSeats INT,
		TheatreName TEXT,
        SlotTiming LIST<Text>,
        Rating DOUBLE,
		TheatreOwnerId TEXT
    )`, keyspace)

	if err := session.Query(createTheatreQuery).Exec(); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
	} else {
		fmt.Printf("Table created successfully in keyspace %s\n", keyspace)
	}
}

//Creating Show table
func CreateTableShow(session *gocql.Session, keyspace string){	
	createShowQuery := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s.Show (
        ShowId TEXT PRIMARY KEY,
		Moviename TEXT,
        ShowTiming TEXT,
        TheatreId TEXT,
        DateOfShow TEXT,
        ReservedSeats INT,
        Availability BOOLEAN,
        SeatInTotal SET  frozen<list<frozen<list<boolean>>>>,
		Price INT
		
    )
`, keyspace)
	if err := session.Query(createShowQuery).Exec(); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
	} else {
		fmt.Printf("Table created successfully in keyspace %s\n", keyspace)
	}
}	
//creating Owner Table	
func CreateTableTheatreOwner(session *gocql.Session, keyspace string){
	
	createTheatreOwnerTableQuery := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s.TheatreOwner (
            MailId TEXT PRIMARY KEY,
			Name TEXT,
			Password TEXT,
			Location TEXT
        )
    `, keyspace)


	if err := session.Query(createTheatreOwnerTableQuery).Exec(); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
	} else {
		fmt.Printf("Table created successfully in keyspace %s\n", keyspace)
	}
}

	//Create Booking table

func CreateTableBooking(session *gocql.Session, keyspace string){
       createBookingQuery := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s.Booking (
			BookingId TEXT primary key,
			ShowId TEXT,
			RequiredSeats int,
			ShowTiming TEXT,
			DateOfShow TEXT,
			totalAmount int,
			SeatPositions  frozen<list<frozen<list<boolean>>>>,
			UserId TEXT
		
        )
    `, keyspace)

	
	if err := session.Query(createBookingQuery).Exec(); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
	} else {
		fmt.Printf("Table created successfully in keyspace %s\n", keyspace)
	}
}
//Creating Review Table
func CreateTableReview(session *gocql.Session, keyspace string){

	
	createReviewQuery := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s.Review(
		mailid TEXT,
		theatreid text,
		comments TEXT,
		PRIMARY KEY (mailid, theatreid)
	)
    `, keyspace)

	
	if err := session.Query(createReviewQuery).Exec(); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
	} else {
		fmt.Printf("Table created successfully in keyspace %s\n", keyspace)
	}
}

	


