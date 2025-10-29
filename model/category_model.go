package models

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title" gorm:"unique;not null"`
}
