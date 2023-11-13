package routes
import(
	"github.com/gin-gonic/gin"
	controller "Theatre/controllers"
	"github.com/gocql/gocql"
	
)
func TheatreOwnerRoutes(incomingRoutes *gin.Engine, session *gocql.Session){
	
	//Router to a theatreOwner to add new theatre
	incomingRoutes.POST("/theatreowner/addtheatres", func(c *gin.Context) {
		controller.AddTheatre(session, c)
	})

	//Router to display all the theatres of a particular theatreOwner
	incomingRoutes.GET("/theatreowner/getalltheatres", func(c *gin.Context) {
		controller.GetAllTheatres(session, c)
	})

	//Router to get particular Theatre
	incomingRoutes.GET("/theatreowner/gettheatrebyid", func(c *gin.Context) {
		controller.GetTheatreDetails(session, c)
	})

	//Router to delete a theatre details
	incomingRoutes.POST("theatreowner/deletetheatre", func(c *gin.Context) {
		controller.DeleteTheatre(session, c)
	})

	//Router to edit theatre details
	incomingRoutes.POST("theatreowner/edittheatre", func(c *gin.Context) {
		controller.EditTheatre(session, c)
	})
	
}