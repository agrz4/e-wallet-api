package model

type SourceOfFund struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

func (SourceOfFund) TableName() string {
	return "source_of_funds"
}
