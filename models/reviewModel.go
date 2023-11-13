package models

//structure for Review Model
type Review struct {
	MailId    string `cql:"mailid"`
	TheatreId string `cql:"theatreid"`
	Comments  string `cql:"comments"`
}
