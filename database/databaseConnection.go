package database
import (
	"log"
	"github.com/gocql/gocql"
)
var cluster *gocql.ClusterConfig// init function for Cluster Creation
func init() { 
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Port = 9042
}

func OpenDB() (*gocql.Session, error) {
	//function for Connection
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	keyspace := "ALGOMOX_MULTIPLEX_BOOKING"//Name of the KeySpace
	if err := session.Query("USE " + keyspace).Exec(); 
	  err != nil {
		//If keyspace already exist then open it,else creating a new one
		createKeyspaceQuery := `
            CREATE KEYSPACE IF NOT EXISTS ` + keyspace + `
            WITH replication = {
                'class': 'SimpleStrategy',
                'replication_factor': 1
            }`

		if err := session.Query(createKeyspaceQuery).Exec(); err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return session, nil
}

