package service

type EventServiceImpl struct {
	repository   repository.UserRepository
	tokenUseCase helper.TokenUseCase
}

func (service *EventServiceImpl) GetEventList() ([]entity.UserEntity, error) {
	getEventList, errGetEventList := service.repository.GetUsers()

	if errGetEventList != nil {
		return []entity.UserEntity{}, errGetEventList
	}

	return entity.ToUserListEntity(getEventList), nil
}
