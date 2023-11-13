package routes

import(
	"github.com/gin-gonic/gin"
	 controller "Theatre/controllers"
	"github.com/gocql/gocql"
	
)
func ShowRoutes(incomingRoutes *gin.Engine, session *gocql.Session){
        
		//Router to add show in a theatre
		 incomingRoutes.POST("/theatreowner/theatre/addshows",func(c *gin.Context){
	        controller.AddShow(session,c)
         })

		 //Router to display all the shows
	     incomingRoutes.GET("/theatreowner/theatre/getallshows",func(c *gin.Context){
			controller.GetAllShows(session,c)
		 })

		 //Router to get a particular show detail
		 incomingRoutes.GET("/theatreowner/theatre/getshowbyid",func(c *gin.Context){
			controller.GetShowById(session,c)
		 })

		 //Router to delete a show
		 incomingRoutes.POST("/theatreowner/theatre/deleteshow",func(c *gin.Context){
	        controller.DeleteShow(session,c)
         })

		 //Router to edit the show details
		 incomingRoutes.POST("/theatreowner/theatre/editshow",func(c *gin.Context){
	        controller.EditShow(session,c)
         })

		
}
