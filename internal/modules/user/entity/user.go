package entity

type User struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Email string
	Username string
	Password string
	Role string
}