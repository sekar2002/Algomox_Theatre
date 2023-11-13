package routes

import(
	"github.com/gin-gonic/gin"
	controller "Theatre/controllers"
	"github.com/gocql/gocql"
)
func ReviewRoutes(incomingRoutes *gin.Engine, session *gocql.Session){
		incomingRoutes.POST("/users/review", func(c *gin.Context) {
				controller.WriteReviews(session, c)})
		incomingRoutes.GET("/users/getallreview", func(c *gin.Context) {
					controller.GetAllReviews(session, c)})		

}