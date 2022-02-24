package service

import (
	contract "dfa/contract-simulation"
	"dfa/entity"
	"dfa/general"
	"time"
)

type BehaviorService interface {
	Register() (entity.Behavior, error)
	RecordAccess(l entity.LogInfo) (entity.Behavior, error)
	Login(pristr, pubstr string) (entity.Behavior, error)
}

type behaviorService struct {
	behavior entity.Behavior
	contract contract.BehaviorContract
}

func NewBehavior() BehaviorService {
	return &behaviorService{
		behavior: *new(entity.Behavior),
		contract: contract.NewBehaviorChain(),
	}
}

func (service *behaviorService) Register() (entity.Behavior, error) {
	b, err := service.contract.Register()
	if err != nil {
		// log.Fatal(err)
		return entity.Behavior{}, err
	}
	service.behavior = b
	return b, nil
}

func (service *behaviorService) RecordAccess(l entity.LogInfo) (entity.Behavior, error) {
	l.TimeStamp = time.Now()
	service.behavior.Logs = append(service.behavior.Logs, l)
	return service.behavior, nil
}

func (service *behaviorService) Login(pristr, pubstr string) (entity.Behavior, error) {
	priBytes, err := general.StrToBytes(pristr)
	if err != nil {
		return entity.Behavior{}, err
	}
	signature, err := general.MakeSignature(priBytes, priBytes)
	if err != nil {
		// log.Fatal(err)
		return entity.Behavior{}, err
	}
	b, err := service.contract.Login(pubstr, signature)
	if err != nil {
		// log.Fatal(err)
		return entity.Behavior{}, err
	}
	service.behavior = b
	return b, nil
}