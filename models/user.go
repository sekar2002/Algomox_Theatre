package models

//structure for User model
type User struct
{  MailId string `json:"MailId"`
   Name   string `json:"Username"`
   Password string
   Location string
}