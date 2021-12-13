package dto

import uuid "github.com/satori/go.uuid"

type LoginRequest struct {
	Email 	string
	Password string
}

type MemberResponse struct {
	ID 			uuid.UUID 	`json:"id"`
	Name 		string 		`json:"name"`
	Address		string		`json:"address"`
	Email		string		`json:"email"`
	Password 	string		`json:"password"`
}
