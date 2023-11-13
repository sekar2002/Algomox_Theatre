package controllers
import("github.com/gocql/gocql"
"github.com/gin-gonic/gin"
"Theatre/tokens"
"Theatre/database"
models "Theatre/models"
"net/http"
"fmt"
 )
func AddTheatre(session *gocql.Session, c *gin.Context){   	
	var newTheatre models.Theatre
	var jwt string
  
	if err := c.ShouldBindJSON(&newTheatre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	jwt = c.GetHeader("jwt")
	
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//function call to insert Theatre(i.e.database/theatreownerqueries)
	if res,err:= database.InsertTheatre(session,newTheatre);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{"Responce":res})
	}

}
func GetAllTheatres(session *gocql.Session, c *gin.Context) ([]models.Theatre, error) {
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	var username string 
	if details,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil,nil
	}else{
		username= details.Username
	}
	//	Query to display theatres
    fmt.Println(username)
	query := "SELECT  TheatreId, TheatreName, TotalSeats, SlotTiming, Rating, TheatreOwnerId FROM ALGOMOX_MULTIPLEX_BOOKING.Theatre WHERE TheatreOwnerId = ? ALLOW FILTERING"
    iter := session.Query(query,username).Iter()
    var theatres []models.Theatre
    var theatre models.Theatre

    for iter.Scan(
		&theatre.TheatreId,
		&theatre.TheatreName, 
		&theatre.TotalSeats, 
		&theatre.SlotTiming, 
		&theatre.Rating, 
		&theatre.TheatreOwnerId) {
        theatres = append(theatres, theatre)
        theatre = models.Theatre{} // Reset the theatre variable for the next iteration
    }

    if err := iter.Close(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return nil, err
    }
	c.JSON(http.StatusOK, gin.H{"data": theatres})
    return theatres, nil
}
func GetTheatreDetails(session *gocql.Session, c *gin.Context){
	
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	var username string 
	var theatre models.Theatre
	if details,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}else{
		username= details.Username
	}
    fmt.Println(username)
	//Query to display particular theatre details
	TheatreId:= c.PostForm("TheatreId")
    fmt.Println(TheatreId)
    query:="SELECT * FROM ALGOMOX_MULTIPLEX_BOOKING.THEATRE WHERE TheatreId=?"
	err := session.Query(query,TheatreId).Consistency(gocql.One).Scan(
		&theatre.TheatreId,
		&theatre.Rating,
		&theatre.SlotTiming,
		&theatre.TheatreName,
		&theatre.TheatreOwnerId,
		&theatre.TotalSeats)
		
   
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
        return 
    }
	c.JSON(http.StatusOK, gin.H{"data": theatre})
    return 
}
func DeleteTheatre(session *gocql.Session, c *gin.Context){
	
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	TheatreId:=c.PostForm("TheatreId")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Function call to Delete Theatre(i.e.database/theatreOwnerQueries)
	if _,err:=database.DeleterTheatreById(session,TheatreId);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"Response": "Theatre Deleted"})
	}
}
func EditTheatre(session *gocql.Session, c *gin.Context){
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	var newtheatre models.Theatre
	jwt = c.GetHeader("jwt")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
	if err := c.ShouldBindJSON(&newtheatre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//Function call to Delete Theatre(i.e.database/theatreOwnerQueries)
	if _,err:=database.DeleterTheatreById(session,newtheatre.TheatreId);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}else{
		//Function call to Insert Theatre(i.e.database/theatreOwnerQueries)
		if _,err:= database.InsertTheatre(session,newtheatre);err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}else{
			c.JSON(http.StatusOK,gin.H{"Response":"Theatre Edited"})
		}
	}
}