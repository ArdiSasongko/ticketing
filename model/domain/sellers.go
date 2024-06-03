package domain

type Seller struct {
	ID       int       `gorm:"column:id;primaryKey"`
	Name     string    `gorm:"column:name"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password"`
	Events   []Event `gorm:"foreignKey:SellerID"`

}
