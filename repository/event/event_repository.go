package repository

type EventRepository interface {
	GetUser(Id int) (domain.User, error)
	GetUsers() ([]domain.User, error)
}
