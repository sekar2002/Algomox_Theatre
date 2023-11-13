package routes

import(
	"github.com/gin-gonic/gin"
	 controller "Theatre/controllers"
	"github.com/gocql/gocql"
)
func UserRoutes(incomingRoutes *gin.Engine, session *gocql.Session){
	
	//Routing to get all available shows when user login
	incomingRoutes.GET("/users/getallavailableshows", func(c *gin.Context) {
		controller.GetAllAvailableShows(session, c)
	})

	//User Routing to search for a movie 
	incomingRoutes.GET("/users/searchmoviebyname", func(c *gin.Context) {
		controller.GetMovieByName(session, c)
	})

	

}