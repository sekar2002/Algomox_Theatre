# **🎦ALGOMOX MULTIPLEX-THEATER BOOKING APPLICATION**

## Introduction

`Algomox_Multiplex_Booking` is a platform where users can book shows in different theatres available. It provides a good list of available movies to watch, where users can browse different movies. This gives the user great relief for those who wish to have quality time without any last-minute rushes and long waiting queues, as it makes booking pretty easy. Not only for users but also for owners who wish to update their theatre details with very few steps.

## ▶️ER Diagram
Below is the Entity Relationship diagram for my proposed system.

 ![ER Diagram](/ERDiagram.png)
 

## ▶️ Application Flow   
The user flow starts with a successful `login` where their passwords are `encrypted` and stored. Then the shows that are available will be displayed where the user can pick the movie that he wishes to watch. Then he can choose the theatre he wishes to watch. After choosing the booking details like `required seats`, `desired seat position`, booking is confirmed. The user also had an option of searching the movie he wishes to watch.
On the other hand, TheatreOwners can log in with their credentials where they can view their theatres. The owners have an option of performing `CRUD` operations on theatre details. On selecting a particular theatre, the owner can view the shows (movies) in that theatre. Here the `CRUD` Operation can be performed on each show.
Both User and TheatreOwner have `JWT` Authentication, which is valid for `1hr`, after which it expires. 
## ▶️ Program Structure
- The Project has following directories
   - `Controllers` : Contains files to handle requests.
   - `Database` : Contains database queries.
   - `Models` : Contains schema for each table.
   - `Routes` : Contains routing files to route the incoming requests based on url.
   - `Tokens` : Contains files for creating and verifying  `JWT` tokens.
- To ensure `reusability and modularity` I have created separate files for each tables.
- `database/databaseCreation.go` file contains queries for creating tables and key spaces.Table creation takes place only once.
-`database/databaseCreationCalling.go` file contains function for all the table creation queries.
- `main.go` start of the application

| Action                  | Model                     | Controller                 | Routes                       |  CRUD                   |
|-------------------------|---------------------------|----------------------------|------------------------------|-------------------------|
|     SignUpTheatreOwner  | TheatreOwner.go           | authenticationController.go|authenticationRoutes.go       | theatreOwnerQueries.go  |
|     SignUpUser          | user.go                   | authenticationController.go|authenticationRoutes.go       | databaseInsertion.go    |
|     Booking             | bookingModel.go           | bookingController.go       |bookingRoutes.go              | bookingQueries.go       |
|     Show                | showModel.go              | showController.go          |showRouter.go                 | showQueries.go          |
|     Theatre             | theatreModel.go           | theatreOwnerController.go  |theatreOwnerRouter.go         | theatreOwnerQueries.go  |
|     Theatre Owner       | theatrOwnerModel.go       | theatreOwnerController.go  |theatreOwnerRouter.go         | theatreOwnerQueries.go  |
|     Review              | reviewModel.go            | reviewController.go        |reviewRouter.go               | reviewQueries.go        |



## ▶️ List Of API Calls
Below are the api endpoints for the application.

### 🔺Authentication

| API Calls                   | Path                           | Description                                  |
|-----------------------------|--------------------------------|----------------------------------------------|
| SignUp TheatreOwner         | `/theatreowner/signup`         | Register a new Theatre Owner account.        |
| SignUp User                 | `/users/signup`                | Register a new User account.                 |
| Login                       | `/users/login`                 | Authenticate and log in a User.              |

### 🔺TheatreOwner
| API Calls                | Path                                       | Description                                  |
|--------------------------|--------------------------------------------|----------------------------------------------|
| Add Theatres             | `/theatreowner/addtheatres`                | Add new theatres to the system.              |
| GetAllTheatres           | `/theatreowner/getalltheatres`             | Retrieve details of all theatres.            |
| GetTheatreDetails        | `/theatreowner/gettheatrebyid`             | Retrieve details of a specific theatre by ID.|
| EditTheatreDetails       | `/theatreowner/edittheatre`                | Modify details of a theatre.                 |
| DeleteTheatre            | `/theatreowner/deletetheatre`              | Delete a theatre from the system.            |
| Add Show                 | `/theatreowner/theatre/addshows`           | Add new shows to a theatre.                  |
| DisplayAllShows          | `/theatreowner/theatre/getallshows`        | Retrieve details of all shows in a theatre.  |
| GetShowDetails           | `/theatreowner/theatre/getshowbyid`        | Retrieve details of a specific show by ID.   |
| EditShowDetails          | `/theatreowner/theatre/editshow`           | Modify details of a show in a theatre.       |
| DeleteShow               | `/theatreowner/theatre/deleteshow`         | Delete a show from a theatre.                |

### 🔺User
| API Calls                      | Path                                       | Description                              |
|--------------------------------|--------------------------------------------|------------------------------------------|
| DisplayAllAvailableShows       | `/users/getallavailableshows`              | Retrieve details of all available shows. |
| Search                         | `/users/searchmoviebyname`                 | Search for movies by name.               |
| Booking                        | `/users/booking`                           | Make a booking for a show.               |
| PreviousBooking                | `/users/getpreviousbookings`               | Retrieve details of previous bookings.   |
| WriteReviews                   | `/users/review`                            | Submit reviews for movies or shows.      |
| DisplayReviews                 | `/users/getallreview`                      | Retrieve reviews submitted by the user.  |


## ▶️ Pre-requisite
To run the application below pre-requisites has to be met.
   - Go (Version `1.21.3`)
   - Cassandra(Version `3.11.16`) - To run, needs Python and Java
   - Python(Version `2.7`) installed
   - Java (Version `1.8`)installed    
## ▶️ How To Run
### 🔺Running Cassandra
   - Navigate  to the `bin` folder in Cassandra.
   - Open `command prompt`, type `cassandra` (Its service starts running).
   - Open another `command prompt`, type `cqlsh` (To ensure Cassandra service).
### 🔺Running go file
- Run the following command in the terminal.
```go
go mod init Theatre
go run main.go
``` 
### ▶️ Database Schema
- Creating a cluster.
- KeySpace Creation-`ALGOMOX_MULTIPLEX_BOOKING`.
- Table Creation        
   1. Users
   2. TheatreOwner
   3. Theatre
   4. Show
   5. Booking
   6. Review

### 🔺Users Schema
| Column Name | Data Type    | Description                      |
|-------------|--------------|----------------------------------|
| MailId      | String       | Unique identifier (mailId)       |
| Name        | String       | Name of the User                 |
| Password    | String       | UserPassword(encrypted)          |
| Location    | String       | Location                         |

### 🔺TheatreOwner Schema
| Column Name | Data Type    | Description                      |
|-------------|--------------|----------------------------------|
| MailId      | String       | Unique identifier (mailId)       |
| Name        | String       | Name of the theatreOwner         |
| Password    | String       | UserPassword(encrypted)          |
| Location    | String       | Location                         |

### 🔺Theatre Schema
| Column Name  | Data Type    | Description                      |
|--------------|--------------|----------------------------------|
| TheatreId    | String       | Id generated for theatre         |
| TheatreName  | String       | Name of the theatre              |
| TotalSeats   | Int          | TotalSeats                       |
| SlotTiming   | []string     | Timings                          |
| Rating       | Int          |  Rating to theatre               |
|TheatreOwnerId| String       |   Id of Owner                    |
### 🔺Show Schema
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
### 🔺Booking Schema
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

### 🔺Review Schema
| Column Name| Data Type   | Cassandra Column Name  |
|------------|-------------|------------------------|
| MailId     | string      | mailid                 |
| TheatreId  | string      | theatreid              |
| Comments   | string      | comments               |

## ▶️  Testing
### 🔺Tool Used-PostMan
   - The file named Theatre.postman_collection.json(`Theatre/Theatre.postman_collection.json`) contains test results which can seen by exporting it via postman.

## ▶️  Future Improvements
- Frontend development.
- Implement functionality to cancel booking.
- Functionality for the Owners to reply to the reviews.
- Email notification to confirm booking.( I have tried several approches, but could not achieve it before deadline.)
## ▶️  Conclusion
In conclusion, the Algomox Multiplex Theater Booking Application strives to provide a streamlined and efficient platform for both users and theater owners. Its well-designed features, robust authentication, and thoughtful database schema contribute to a reliable and enjoyable theater booking experience.Kindly contact the below email for any queries.
[sowmiyaa.20cse@kongu.edu](sowmiyaa.20cse@kongu.edu)

![Thank You gif](https://media.tenor.com/nu9XqB4pQGsAAAAi/line-sticker-thank-you-sticker.gif)
