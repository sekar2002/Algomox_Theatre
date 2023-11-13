package controllers
import("github.com/gocql/gocql"
		"github.com/gin-gonic/gin"
		"Theatre/tokens"
		"Theatre/database"
		 models "Theatre/models"
		"net/http"
		

)
//Controller to get all shows available
func GetAllAvailableShows(session *gocql.Session, c *gin.Context){
	//VerifyToken function for checking token is vaild or not(i.e.tokens/JWT.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//Function call to display shows(database/showQueries)
	shows,_:=database.GetAllShows(session)
	var availableShows[] models.Show
	for _, show := range shows{
        if show.Availability{
			availableShows=append(availableShows,show)
		}
	}
	c.JSON(http.StatusOK, gin.H{"Response": availableShows})

}
//Search by movie name
func GetMovieByName(session *gocql.Session, c *gin.Context){
	//VerifyToken function for checking token is vaild or not(i.e.tokens/JWT.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//Getting movie name for searching
	movieToSearch:=c.PostForm("MovieName")
	//Function to Get the show details (i.e.showqueries.go)
	show,err:=database.GetMovieDetails(session,movieToSearch)
	if(err!=nil){
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"Response": show})
}

