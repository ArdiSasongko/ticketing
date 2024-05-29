package service

type EventService interface {
	GetUserList() ([]entity.UserEntity, error)
	DeleteEvent(request web.UserEventServiceRequest, pathId int) (map[string]interface{}, error)
}
