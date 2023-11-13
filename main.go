package main
import (
    
        "fmt"
        "Theatre/database"
        "log"
        "github.com/gin-gonic/gin"
        "Theatre/routes"
)
func main() {
    //Connection to Cassandra DataBase i.e. (database.databaseConnection.go)
    session, err := database.OpenDB()
    if err != nil {
        fmt.Println("Failed to open Cassandra session:", err)
        return
    }
    defer session.Close()
    //Table creation i.e.(database.databaseCreation.go)
    database.CreateTable(session, "ALGOMOX_MULTIPLEX_BOOKING")
    if err != nil {
        log.Fatal("Error creating table:", err)
        return
    }
    router:=gin.New()
    router.Use(gin.Logger())
    routes.AuthenticationRoutes(router,session)
    routes.TheatreOwnerRoutes(router,session)
    routes.ShowRoutes(router,session)
    routes.UserRoutes(router,session)
    routes.ReviewRoutes(router,session)
    routes.BookingRoutes(router,session)
    //Connecting to the port 8000
    router.Run(":8000")

}


