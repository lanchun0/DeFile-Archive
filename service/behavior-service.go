package service

import "dfa/entity"

type BehaviorService interface {
	Register() (entity.Behavior, bool)
	RecordAccess(entity.LogInfo) bool
}

type behaviorService struct {
	behavior entity.Behavior
}

func NewBehavior() BehaviorService {
	return &behaviorService{
		behavior: *new(entity.Behavior),
	}
}

func (service *behaviorService) Register() (entity.Behavior, bool) {
	b := entity.Behavior{
		PublicKey:  "PublicKey",
		PrivateKey: "PrivateKey",
		Asset:      []string{},
		Logs:       []entity.LogInfo{},
	}
	service.behavior = b
	return b, true
}

func (service *behaviorService) RecordAccess(l entity.LogInfo) bool {
	service.behavior.Logs = append(service.behavior.Logs, l)
	return true
}
