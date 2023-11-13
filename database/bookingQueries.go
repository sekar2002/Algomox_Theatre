package database
import(
	"github.com/gocql/gocql"
	"Theatre/models"
	"fmt"
	"log"
	"github.com/google/uuid"
	//"gopkg.in/gomail.v2"
	//"crypto/tls"
	"net/smtp"
	"github.com/jordan-wright/email"
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
func SendConfirmation(ession *gocql.Session,username string,BookingId string)(string,error){
	// m := gomail.NewMessage()
	// m.SetHeader("From", "sowmiyaa.20cse@kongu.edu") // replace with your email
	// m.SetHeader("To", username)
	// m.SetHeader("Subject", "Booking Confirmation")
	// body := fmt.Sprintf("Your booking with ID %s is confirmed!",BookingId)
	// m.SetBody("Have a nice time :)", body)
	// d := gomail.NewDialer("sowmiyasekar000@gmail.com", 587, "vimalkumars.20cse@kongu.edu", "9003874129") 
	// if err := d.DialAndSend(m); err != nil {
	// 	return "",err
	// }
	// return "success",nil
	/*emailApPassword := "Luffy@2002:)"
	yourMail := "selvakumarvimalkumar@gmail.com"
	recipient := username
	hostAddress := "smtp.gmail.com"
	hostPort := "465"
	mailSubject := "Booking Confirmation"
	mailBody := fmt.Sprintf("Your booking with ID %s is confirmed!",BookingId)
	fullServerAddress := hostAddress + ":" + hostPort
	headerMap := make(map[string]string)
        headerMap["From"] = yourMail
        headerMap["To"] = recipient
        headerMap["Subject"] = mailSubject
        mailMessage := ""
        for k, v := range headerMap {
                mailMessage += fmt.Sprintf("%s: %s\\r", k, v)
        }
        mailMessage += "\\r" + mailBody
		authenticate := smtp.PlainAuth("", yourMail, emailApPassword, hostAddress)
        tlsConfigurations := &tls.Config{
                InsecureSkipVerify: true,
                ServerName: hostAddress,
        }
        conn, err := tls.Dial("tcp", fullServerAddress, tlsConfigurations)
        if err != nil {
                log.Panic(err)
				return "",err
        }
newClient, err := smtp.NewClient(conn, hostAddress)
        if err != nil {
                log.Panic(err)
				return "",err
        }
        // Auth
        if err = newClient.Auth(authenticate); err != nil {
                log.Panic(err)
				return "",err
        }
        // To && From
        if err = newClient.Mail(yourMail); err != nil {
                log.Panic(err)
				return "",err
        }
        if err = newClient.Rcpt(headerMap["To"]); err != nil {
                log.Panic(err)
				return "",err
        }
		 // Data
		 writer, err := newClient.Data()
		 if err != nil { 
				 log.Panic(err)
				 return "",err 
		 }
 
		 _, err = writer.Write([]byte(mailMessage))
 
		 if err != nil { 
				 log.Panic(err)
				 return "",err 
		 }
		 err = writer.Close()
		 if err != nil { 
				 log.Panic(err) 
				 return "",err
		 } 
		 err = newClient.Quit() 
		 if err != nil { 
				return "",err 
		 }
		 fmt.Println("Successful, the mail was sent!")
         return "mail sent",nil*/

	emailApPassword := "Luffy@2002:)"
	yourMail := "vimalkumars.20cse@kongu.edu"
	recipient := username
	hostAddress := "smtp.gmail.com"
	hostPort := "465"
	emailInstance := email.NewEmail()

	emailInstance.From = yourMail
        emailInstance.To = []string{recipient}
        emailInstance.Subject = "Booking Confirmation"
        emailInstance.Text = []byte("Your booking with ID "+BookingId+" is confirmed!")

	err := emailInstance.Send(fmt.Sprintf("%s:%s", hostAddress, hostPort), smtp.PlainAuth("", yourMail, emailApPassword, hostAddress))
	if err != nil {
			fmt.Println("There was an error sending the mail: ",err)
			return "", err
	}
	return "Mail sent :)",nil
}