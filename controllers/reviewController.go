package controllers
import (
	"github.com/gocql/gocql"
		"github.com/gin-gonic/gin"
		"Theatre/tokens"
		"Theatre/database"
		 models "Theatre/models"
		"net/http"

)


func WriteReviews(session *gocql.Session, c *gin.Context){
	//VerifyToken function for checking token is vaild or not(i.e.tokens/JWT.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
   }
   var review models.Review
   if err := c.ShouldBindJSON(&review); err != nil {
	   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	   return 
   }
   _,err:=database.GiveReviews(session,review)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Response": "Thankyou for your reviews"})

}
func GetAllReviews(session *gocql.Session, c *gin.Context){
	var jwt string
	jwt = c.GetHeader("jwt")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
   }
   TheatreId:=c.PostForm("TheatreId")
   review,err:=database.GetReviews(session,TheatreId)
   if err != nil{
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}
c.JSON(http.StatusOK, gin.H{"Response": review})

}