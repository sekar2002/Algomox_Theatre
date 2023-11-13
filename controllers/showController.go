
package controllers
import("github.com/gocql/gocql"
"github.com/gin-gonic/gin"
"Theatre/tokens"
"Theatre/database"
models "Theatre/models"
"net/http"


 )
func AddShow(session *gocql.Session, c *gin.Context){ 
	var newShow models.Show
	var jwt string
	
	if err := c.ShouldBindJSON(&newShow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	jwt = c.GetHeader("jwt")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Function call to insert Show(i.e.database/showqueries.go)
	res,err := database.InsertShow(session, newShow)
	if(err!=nil){
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	    return 

	}
	c.JSON(http.StatusOK, gin.H{"Response": res})
}
func GetAllShows(session *gocql.Session, c *gin.Context){
	
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	TheatreId:=c.PostForm("TheatreId")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	//Function call to display Show((i.e.database/showqueries.go)
	res,err := database.DisplayShow(session,TheatreId)
	if(err!=nil){
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	    return

	}
	c.JSON(http.StatusOK, gin.H{"Response": res})	
}
func GetShowById(session *gocql.Session, c *gin.Context){
	
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	ShowId := c.PostForm("ShowId")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//Function call to display particular Show((i.e.database/showqueries.go)
	show,err:=database.GetShowDetails(session,ShowId)
	if(err!=nil){
	c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	return
	}
    c.JSON(http.StatusOK, gin.H{"Response": show})	

}
func DeleteShow(session *gocql.Session, c *gin.Context){
	
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	jwt = c.GetHeader("jwt")
	ShowId := c.PostForm("ShowId")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//Function call to delete Show((i.e.database/showqueries.go)
	res,err:=database.DeleteShowById(session,ShowId)
	if(err!=nil){
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
		}
		c.JSON(http.StatusOK, gin.H{"Response": res})	
}
func EditShow(session *gocql.Session, c *gin.Context){
	//VerifyToken function for checking token is vaild or not(tokens/jwt.go)
	var jwt string
	var newShow models.Show
	jwt = c.GetHeader("jwt")
	if _,err:= tokens.VerifyToken(jwt);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if err := c.ShouldBindJSON(&newShow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	//Function call to delete Show((i.e.database/showqueries.go)
	_,err:=database.DeleteShowById(session,newShow.ShowId)
	if(err!=nil){
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
		}else{			
			//Function call to insert Show(i.e.database/showqueries.go)
			_,err := database.InsertShow(session, newShow)
			if(err!=nil){
					c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
					return
				}
			c.JSON(http.StatusOK, gin.H{"Response": "Show Updated :)"})
		}
		
}
