package database
import(
	"github.com/gocql/gocql"
)

//Calling  table creation function
func CreateTable(session *gocql.Session, keyspace string){
CreateTableUser(session, keyspace)
CreateTableTheatre(session, keyspace)
CreateTableShow(session, keyspace)
CreateTableTheatreOwner(session, keyspace)
CreateTableBooking(session, keyspace)
CreateTableReview(session, keyspace)

}