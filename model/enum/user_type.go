package enum

type UserType string

const (
	UserTypeBuyer  UserType = "buyer"
	UserTypeSeller UserType = "seller"
	UserTypeAdmin  UserType = "admin"
)
