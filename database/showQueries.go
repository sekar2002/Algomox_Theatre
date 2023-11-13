package database
import(
	"fmt"
	"github.com/gocql/gocql"
	models "Theatre/models"
	"github.com/google/uuid"
	"log")

	//Function to insert to new show
func InsertShow(session *gocql.Session,show models.Show)(string,error){
	fmt.Printf("%v\n",show)
		insertQuery:="INSERT INTO ALGOMOX_MULTIPLEX_BOOKING.Show(Moviename,Price,ShowTiming,TheatreId,DateOfShow,ReservedSeats,Availability,SeatInTotal,ShowId) VALUES (?, ?, ?, ?, ?,?,?,?,?)"
		if(len(show.ShowId)==0){
			id:=uuid.New()
			show.ShowId =id.String()
		}
		
		if err := session.Query(
			insertQuery,
			show.Moviename,
			show.Price,
			show.ShowTiming,
			show.TheatreId,
			show.DateOfShow,
			show.ReservedSeats,
			show.Availability,
			show.SeatInTotal,
			show.ShowId,).Exec(); err != nil {
				log.Fatalf("Error inserting data: %v\n", err)
				return "ERROR",err
			} else {
				fmt.Println("Data inserted successfully.")
			}
			return "Show added",nil	
}

//Function to display all the shows in a theatre
func DisplayShow(session *gocql.Session,TheatreId string)([]models.Show,error){
	query := "SELECT  *  FROM ALGOMOX_MULTIPLEX_BOOKING.SHOW WHERE TheatreId = ? ALLOW FILTERING"
    iter := session.Query(query,TheatreId).Iter()
    var shows []models.Show
    var show models.Show
    for iter.Scan(
		&show.ShowId,
		&show.Availability,
		&show.DateOfShow,
		&show.Moviename,
		&show.Price,
		&show.ReservedSeats,
		&show.SeatInTotal,
		&show.ShowTiming,
		&show.TheatreId,
		) {
        shows = append(shows, show)
        show = models.Show{} 
    }
    if err := iter.Close(); err != nil {
        return nil, err
    }
    return shows, nil
}

//Function to display a particular show details
func GetShowDetails(session *gocql.Session,ShowId string)(*models.Show,error){
	var show models.Show
	query:="SELECT * FROM ALGOMOX_MULTIPLEX_BOOKING.SHOW WHERE ShowId=?"
	if err := session.Query(query,ShowId).Consistency(gocql.One).Scan(
		&show.ShowId,
		&show.Availability,
		&show.DateOfShow,
		&show.Moviename,
		&show.Price,
		&show.ReservedSeats,
		&show.SeatInTotal,
		&show.ShowTiming,
		&show.TheatreId,
		);err != nil{
			return nil,err
		}else{
			return &show,nil
		}
}

//Function to delete a partilcular show
func DeleteShowById(session *gocql.Session,ShowId string)(string,error){
	deleteQuery:="DELETE FROM algomox_multiplex_booking.Show WHERE ShowId = ?"
	if err := session.Query(
		deleteQuery,
		ShowId,
	).Exec(); err != nil {
		log.Fatalf("Error deleting Show: %v\n", err)
		return "error",err
	}else{
		fmt.Println("Data deleted successfully.")
	}
	return "Show deleted",nil
}

//Function to display only available shows for the users
func GetAllShows(session *gocql.Session)([]models.Show,error){
  
   query := "SELECT  * FROM ALGOMOX_MULTIPLEX_BOOKING.Show"
    iter := session.Query(query).Iter()
    var shows []models.Show
    var show models.Show
    for iter.Scan(
		&show.ShowId,
		&show.Availability,
		&show.DateOfShow,
		&show.Moviename,
		&show.Price,
		&show.ReservedSeats,
		&show.SeatInTotal,
		&show.ShowTiming,
		&show.TheatreId,
		) {
        shows = append(shows, show)
        show = models.Show{} 
    }
	if err := iter.Close(); err != nil {
        return nil, err
    }
    return shows, nil


}
//Function to search for a movie for a user
func GetMovieDetails(session *gocql.Session,movieToSearch string)([]models.Show,error){
    query := "SELECT  *  FROM ALGOMOX_MULTIPLEX_BOOKING.SHOW WHERE Moviename>=? and Moviename<? ALLOW FILTERING"
    iter := session.Query(query,movieToSearch,movieToSearch+"z").Iter()
    var shows []models.Show
    var show models.Show
    for iter.Scan(
		&show.ShowId,
		&show.Availability,
		&show.DateOfShow,
		&show.Moviename,
		&show.Price,
		&show.ReservedSeats,
		&show.SeatInTotal,
		&show.ShowTiming,
		&show.TheatreId,
		) {
        shows = append(shows, show)
        show = models.Show{} 
    }
    if err := iter.Close(); err != nil {
        return nil, err
    }
    return shows, nil
}