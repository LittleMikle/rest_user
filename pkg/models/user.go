package models

type User struct {
	ID    uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}
