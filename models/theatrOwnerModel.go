package models

//structure for TheatreOwner
type TheatreOwner struct{
	MailId string `json:"MailId"`
   Name   string `json:"Username"`
   Password string
   Location string
}