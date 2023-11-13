package controllers

import
(
	"github.com/gocql/gocql"
"github.com/gin-gonic/gin"
"Theatre/tokens"
"Theatre/database"
models "Theatre/models"
"net/http"
"fmt"
)
//Controller for booking
func Booking(session *gocql.Session, c *gin.Context){
	//VerifyToken function for checking token is vaild or not(i.e.tokens/JWT.go)
	var jwt string
	var username string
	jwt = c.GetHeader("jwt")
	if details,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}else{		
		username= details.Username
	}
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	booking.UserId = username
	//function to create booking(database/bookingQueries.go)
	bookingId,err:=database.CreateBooking(session,booking)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_,err=database.SendConfirmation(session,booking.UserId,bookingId)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{"Response": bookingId})
	
}
func GetPreviousBooking(session *gocql.Session, c *gin.Context){
		//VerifyToken function for checking token is vaild or not(i.e.tokens/JWT.go)
		var jwt string
	    jwt = c.GetHeader("jwt")
	    var username string 
	    if details,err:= tokens.VerifyToken(jwt);err!=nil{
		    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
	}else{
		username= details.Username
	}
	fmt.Printf(username)
	//function to know the booking that the user had made
		booking,err:=database.PreviousHistory(session,username)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Response": booking})
}
