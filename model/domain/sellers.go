package domain

type Sellers struct {
	SellerID int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (Sellers) TableName() string {
	return "seller"
}
