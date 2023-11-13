package database
import(
	"fmt"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	models "Theatre/models"
	"github.com/google/uuid"
	"log")

//Function to insert new TheatreOwner in TheatreOwner Table
func InsertTheatreOwner(session *gocql.Session,user models.TheatreOwner) string {
	//function to store the password in a secured form by encrypting
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	hashedPasswordString := string(hashedPassword)
	//Inserting a new TheatreOwner in  table
    insertStatement := fmt.Sprintf("INSERT INTO %s.%s ( MailId,Name,Password,Location) VALUES (?, ?, ?, ?)IF NOT EXISTS","ALGOMOX_MULTIPLEX_BOOKING","TheatreOwner")
	if err := session.Query(insertStatement, user.MailId,user.Name,hashedPasswordString,user.Location,).Exec(); err != nil {
		log.Printf("Error inserting data into table: %v\n", err)
		return "Theatre Owner already exist "
	}
	log.Println("Data inserted successfully.")
	return "Theatre Owner  added"
}

//Function to insert theatre
func InsertTheatre(session *gocql.Session,theatre models.Theatre)(string,error){
	insertQuery := "INSERT INTO ALGOMOX_MULTIPLEX_BOOKING.Theatre (TheatreId,TheatreName, TotalSeats, SlotTiming, Rating,TheatreOwnerId) VALUES (?, ?, ?, ?, ?,?)"
    id := uuid.New()
    TheatreId:=id.String()
	if err := session.Query(
		insertQuery,
		TheatreId,
		theatre.TheatreName,
		theatre.TotalSeats,
		theatre.SlotTiming,
		0.0,
		theatre.TheatreOwnerId,
	).Exec(); err != nil {
		log.Fatalf("Error inserting data: %v\n", err)
		return "Error inserting theatre ",err
	} else {
		fmt.Println("Data inserted successfully.")
	}
	return "Theatre added",nil
}

//Function to delete Theatre
func DeleterTheatreById(session *gocql.Session,TheatreId string)(string,error){
	deleteQuery:="DELETE FROM algomox_multiplex_booking.Theatre WHERE TheatreId = ?"
	if err := session.Query(
		deleteQuery,
		TheatreId,
	).Exec(); err != nil {
		log.Fatalf("Error deleting theatre: %v\n", err)
		return "error",err
	}else{
		fmt.Println("Data deleted successfully.")
	}
	return "Theatre deleted",nil
}
