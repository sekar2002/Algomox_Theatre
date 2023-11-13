package routes
import(
	"github.com/gin-gonic/gin"
	controller "Theatre/controllers"
	"github.com/gocql/gocql"
	
)
func AuthenticationRoutes(incomingRoutes *gin.Engine, session *gocql.Session) {
	//Router to signup users
	incomingRoutes.POST("/users/signup", func(c *gin.Context) {
		controller.SignUpUser(session, c)
	})

	//Router to login 
	incomingRoutes.POST("/users/login", func(c *gin.Context) {
		controller.Login(session, c)
	})

	//Router to signup theatreowner
	incomingRoutes.POST("/theatreowner/signup", func(c *gin.Context) {
		controller.SignUpTheatreOwner(session, c)
	})
}
