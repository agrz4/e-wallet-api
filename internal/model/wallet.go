package model

type Wallet struct {
	ID      uint `gorm:"primarykey"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Number  string
	Balance int
}

func (Wallet) TableName() string {
	return "wallets"
}
