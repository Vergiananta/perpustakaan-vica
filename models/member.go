package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type Member struct {
	ID 			uuid.UUID 	`json:"id"`
	Name 		string 		`json:"name"`
	Address		string		`json:"address"`
	Email 		string		`json:"email"`
	Password 	string		`json:"password"`
	CreatedAt   time.Time  	`json:"created_at"`
	UpdatedAt   *time.Time  	`json:"updated_at"`
	DeletedAt   *time.Time 	`json:"deleted_at" sql:"index"`
}

func (m *Member) Prepare() error {
	m.ID = uuid.NewV4()
	hashedPassword, err := Hash(m.Password)
	if err != nil {
		return err
	}
	m.Password = string(hashedPassword)
	m.Email = html.EscapeString(strings.TrimSpace(m.Email))
	m.CreatedAt = time.Now()
	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}