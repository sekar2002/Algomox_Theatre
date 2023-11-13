package database
import
(
	"fmt"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	models "Theatre/models"
	"log"
)
//Function to insert new User
func InsertUsers(session *gocql.Session,user models.User) string {
	//function to store the password in a secured form by encrypting
	 hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	//User password column is of type string
	hashedPasswordString := string(hashedPassword)
	//Inserting a new User in User table
	insertStatement := fmt.Sprintf("INSERT INTO %s.%s ( MailId,Name,Password,Location) VALUES (?, ?, ?, ?)IF NOT EXISTS","ALGOMOX_MULTIPLEX_BOOKING","Users")
	if err := session.Query(insertStatement, user.MailId,user.Name,hashedPasswordString,user.Location,).Exec(); err != nil {
			log.Printf("Error inserting data into table: %v\n", err)
			
			return "User already exist "
		}
	log.Println("Data inserted successfully.")
	return "User added"
}

func GetPasswordByID(session *gocql.Session,MailId string) (string, error) {
	//function to know the password of existing User by giving MailId
	var password string
	query := "SELECT password FROM ALGOMOX_MULTIPLEX_BOOKING.Users WHERE MailId = ?"
	err := session.Query(query, MailId).Consistency(gocql.One).Scan(&password)
	if err != nil {
		query := "SELECT password FROM ALGOMOX_MULTIPLEX_BOOKING.TheatreOwner WHERE MailId = ?"
		err := session.Query(query, MailId).Consistency(gocql.One).Scan(&password)
		if err!=nil{
			return "", err}
		return password,nil	
	}
    
	return password, nil
}
