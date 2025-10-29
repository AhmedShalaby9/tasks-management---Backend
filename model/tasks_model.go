package models

type Task struct {
	ID          int      `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description *string  `json:"description"`
	Done        bool     `json:"done"`
	CategoryID  *int     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}
