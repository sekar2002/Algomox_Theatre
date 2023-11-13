package routes

import(
	"github.com/gin-gonic/gin"
	controller "Theatre/controllers"
   "github.com/gocql/gocql"
)
func BookingRoutes(incomingRoutes *gin.Engine, session *gocql.Session){
		//To book  tickets for a show
		incomingRoutes.POST("/users/booking", func(c *gin.Context) {
			controller.Booking(session, c)
		})
		incomingRoutes.GET("/users/getpreviousbookings", func(c *gin.Context) {
			controller.GetPreviousBooking(session, c)
		})
}
