package controllers
import (
	
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	database "Theatre/database"
	models "Theatre/models"
	"Theatre/tokens"
)
//SignUp for User
func SignUpUser(session *gocql.Session, c *gin.Context) {
	var newUser models.User
  
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Added return statement to exit the function on error
	}
	res := database.InsertUsers(session, newUser)
	c.JSON(http.StatusOK, gin.H{"Response": res})
}
//SignUp for TheatreOwners
func SignUpTheatreOwner(session *gocql.Session, c *gin.Context) {
	var newUser models.TheatreOwner
	
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Added return statement to exit the function on error
	}
    
	res := database.InsertTheatreOwner(session, newUser)
	c.JSON(http.StatusOK, gin.H{"Response": res})
}
//Login function
func Login(session *gocql.Session, c *gin.Context) {
	mailID := c.PostForm("MailId")
	userEnteredPassword:=c.PostForm("Password")

	Password,_:= database.GetPasswordByID(session, mailID)
	err := bcrypt.CompareHashAndPassword([]byte(Password), []byte(userEnteredPassword))
	if err == nil {
		//Login Successfull,then Function call to JWT Token Generation(i.e.tokens/JWT.go)
		jwttoken,_:=tokens.GenerateToken(mailID)
        
		c.JSON(http.StatusOK, gin.H{"JWT": jwttoken})
	} else {
		// If password doesnt matches,then rise error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Username or Password"})
	}
	
}



