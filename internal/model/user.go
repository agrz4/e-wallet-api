package model

type User struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Email    string
	Password string
}

func (User) TableName() string {
	return "users"
}
