# **🎦ALGOMOX MULTIPLEX-THEATER BOOKING APPLICATION**

## **Introduction**

`Algomox_Multiplex_Booking` is a platform where users can book shows in different theatres available. This helps us to book tickets with just a few clicks. It provides a good list of available movies to watch, where users can browse different movies. This gives the user great relief for those who wish to have quality time without any last-minute rushes and long waiting queues, as it makes booking pretty easy. Not only for users but also for owners who wish to update their theatre details with very few steps.

## ER Diagram
 ![ER Diagram](/ERDiagram.png)

## List Of API Calls
### Authentication

| API Calls                   | Path                           | Description                                  |
|-----------------------------|--------------------------------|----------------------------------------------|
| SignUp TheatreOwner         | `/theatreowner/signup`         | Register a new Theatre Owner account.        |
| SignUp User                 | `/users/signup`                | Register a new User account.                 |
| Login                       | `/users/login`                 | Authenticate and log in a User.               |

### TheatreOwner
| API Calls                | Path                                       | Description                                          |
|--------------------------|--------------------------------------------|------------------------------------------------------|
| Add Theatres             | `/theatreowner/addtheatres`                | Add new theatres to the system.                      |
| GetAllTheatres           | `/theatreowner/getalltheatres`             | Retrieve details of all theatres.                    |
| GetTheatreDetails        | `/theatreowner/gettheatrebyid`             | Retrieve details of a specific theatre by ID.        |
| EditTheatreDetails       | `/theatreowner/edittheatre`                | Modify details of a theatre.                         |
| DeleteTheatre            | `/theatreowner/deletetheatre`              | Delete a theatre from the system.                    |
| Add Show                 | `/theatreowner/theatre/addshows`           | Add new shows to a theatre.                          |
| DisplayAllShows          | `/theatreowner/theatre/getallshows`        | Retrieve details of all shows in a theatre.          |
| GetShowDetails           | `/theatreowner/theatre/getshowbyid`        | Retrieve details of a specific show by ID.           |
| EditShowDetails          | `/theatreowner/theatre/editshow`           | Modify details of a show in a theatre.               |
| DeleteShow               | `/theatreowner/theatre/deleteshow`         | Delete a show from a theatre.                        |

### User
| API Calls                      | Path                                       | Description                                             |
|--------------------------------|--------------------------------------------|---------------------------------------------------------|
| DisplayAllAvailableShows       | `/users/getallavailableshows`              | Retrieve details of all available shows.                |
| Search                         | `/users/searchmoviebyname`                 | Search for movies by name.                              |
| Booking                        | `/users/booking`                           | Make a booking for a show.                              |
| PreviousBooking                | `/users/getpreviousbookings`               | Retrieve details of previous bookings.                  |
| WriteReviews                   | `/users/review`                            | Submit reviews for movies or shows.                     |
| DisplayReviews                 | `/users/getallreview`                      | Retrieve details of all reviews submitted by the user.  |


## How does Control flow   
The user flow starts with a successful login where their passwords are encrypted and stored. Then the shows that are available will be displayed where the user can pick the movie that he wishes to watch. Then he can choose the theatre he wishes to watch. After choosing the booking details like required seats, desired seat position, booking is confirmed. The user also has an option of searching the movie he wishes to watch.
On the other hand, TheatreOwners can log in with their credentials where they can view their theatres. The owners have an option of performing CRUD operations on theatre details. On selecting a particular theatre, the owner can view the shows (movies) in that theatre. Here the CRUD Operation can be performed on each show.
Both User and TheatreOwner have JWT Authentication, which has a validity of 1hr, after which it expires. 

## Pre-requisite
   1. Go
   2. Cassandra - To run, needs Python and Java
   3. Python installed
   4. Java installed    
## How To Run
### Running Cassandra
    1. Navigate up to the bin folder in Cassandra.
    2. Open cmd, type cassandra (Its service starts running).
    3. Open cmd, type cqlsh (To ensure Cassandra service).
   ### Running go file
   1. Type go mod init projectname.
   2. Run go run main.go.
### Database Schema
1. Creating a cluster.
2. KeySpace Creation-`ALGOMOX_MULTIPLEX_BOOKING`.
3. Table Creation
        1. Users
        2. TheatreOwners
        3. Theatres
        4. Shows
        5. Booking
        6. Reviews

## Table Schema
### User Schema
| Column Name | Data Type    | Description                      |
|-------------|--------------|----------------------------------|
| MailId      | String       | Unique identifier (mailId)       |
| Name        | String       | Name of the User                 |
| Password    | String       | UserPassword(encrypted)          |
| Location    | String       | Location                         |

### TheatreOwner Schema
| Column Name | Data Type    | Description                      |
|-------------|--------------|----------------------------------|
| MailId      | String       | Unique identifier (mailId)       |
| Name        | String       | Name of the theatreOwner         |
| Password    | String       | UserPassword(encrypted)          |
| Location    | String       | Location                         |

### Theatre Schema
| Column Name  | Data Type    | Description                      |
|--------------|--------------|----------------------------------|
| TheatreId    | String       | Id generated for theatre         |
| TheatreName  | String       | Name of the theatre              |
| TotalSeats   | Int          | TotalSeats                       |
| SlotTiming   | []string     | Timings                          |
| Rating       | Int          |  Rating to theatre               |
|TheatreOwnerId| String       |   Id of Owner                    |
### Show Schema
| Column Name    | Data Type| Description                          |
|----------------|----------|--------------------------------------|
| Moviename      | string   | Name of the movie                    |
| ShowTiming     | string   | Timing of the show                   |
| TheatreId      | string   | Identifier for the theatre           |
| DateOfShow     | string   | Date of the show                     |
| ReservedSeats  | int      | Number of reserved seats             |
| Availability   | bool     | Availability status                  |
| SeatInTotal    | [][]bool | 2D array representing seat status    |
| ShowId         | string   | Identifier for the show              |
| Price          | int      | Price of the ticket                  |
### Booking Schema
| Column Name    | Data Type| Description                       |
|----------------|----------|-----------------------------------|
| BookingId      | string   | Unique identifier for the booking |
| ShowId         | string   | Identifier for the show           |
| RequiredSeats  | int      | Number of seats required          |
| ShowTiming     | string   | Timing of the show                |
| DateOfShow     | string   | Date of the show                  |
| TotalAmount    | int      | Total amount for the booking      |
| SeatPositions  | [][]bool | 2D array representing seat status |
| UserId         | string   | Identifier for the user           |

### Review Schema
| Column Name| Data Type   | Cassandra Column Name  |
|------------|-------------|------------------------|
| MailId     | string      | mailid                 |
| TheatreId  | string      | theatreid              |
| Comments   | string      | comments               |

## Conclusion
In conclusion, the Algomox Multiplex Theater Booking Application strives to provide a streamlined and efficient platform for both users and theater owners. Its well-designed features, robust authentication, and thoughtful database schema contribute to a reliable and enjoyable theater booking experience.

