package api

import "time"

type JSONProductOutput struct {
	ID            string    `json:"_id"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Description   string    `json:"description"`
	Condition     string    `json:"condition"`
	Size          string    `json:"size"`
	Color         string    `json:"color"`
	ViewCount     int64     `json:"viewCount"`
	Brand         string    `json:"brand"`
	Category      string    `json:"category"`
	Location      string    `json:"location"`
	ImageFilePath string    `json:"imgPath"`
	Price         float64   `json:"price"`
	Tags          []string  `json:"tags"`
	Author        string    `json:"author"`
	CreatedAt     time.Time `json:"createdAt"`
}

type JSONProductInput struct {
	ImgPath []byte `json:"imgPath"`
	Name    string `json:"name"`
}
