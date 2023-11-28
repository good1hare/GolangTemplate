package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     int       `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
