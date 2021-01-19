package model

import "time"

type UserCredentials struct {
	Timestamp time.Time
	Users []User
}

var UserCredentialsManagement UserCredentials