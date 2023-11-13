package database
import("fmt"
"log"
"github.com/gocql/gocql"
"Theatre/models"
)
func GiveReviews(session *gocql.Session,review models.Review )(string,error){
	insertQuery := "INSERT INTO ALGOMOX_MULTIPLEX_BOOKING.Review (MailId,TheatreId,Comments) VALUES (?, ?, ?)"
	if err := session.Query(
		insertQuery,
		review.MailId,
		review.TheatreId,
		review.Comments,
	).Exec(); err != nil {
		log.Fatalf("Error inserting data: %v\n", err)
		return "Error inserting theatre ",err
	} else {
		fmt.Println("Data inserted successfully.")
	}
	return "Review added",nil
}
func GetReviews(session *gocql.Session,TheatreId string)([]models.Review,error){
	
	query := "SELECT  * FROM ALGOMOX_MULTIPLEX_BOOKING.Review where TheatreId=? ALLOW FILTERING"
    iter := session.Query(query,TheatreId).Iter()
    var reviews []models.Review
    var review models.Review
    for iter.Scan(
		&review.MailId,
		&review.Comments,
		&review.TheatreId,
		) {
        reviews = append(reviews, review)
        review = models.Review{} 
    }
	if err := iter.Close(); err != nil {
        return nil, err
    }
    return reviews, nil
}