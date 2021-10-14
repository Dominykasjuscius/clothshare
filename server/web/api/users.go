package api

import (
	"time"
)

type JSONUserInput struct {
	Name      string `json:"name"`
	Password  string `json:"pass"`
	Email     string `json:"email"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
	PhotoPath string `json:"photoPath"`
}

type JSONUserOutput struct {
	ID         string    `json:"_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Location   string    `json:"location"`
	Bio        string    `json:"bio"`
	Rating     int       `json:"rating"`
	PhotoPath  string    `json:"photoPath"`
	CreatedAt  time.Time `json:"createdAt"`
	LastSeenAt time.Time `json:"lastSeenAt"`
	UpdatedAt  time.Time `json:"updatedAt"`

	Followers []JSONUserOutput    `json:"followers"`
	Products  []JSONProductOutput `json:"products"`
}
