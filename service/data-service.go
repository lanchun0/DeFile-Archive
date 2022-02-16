package service

import (
	"dfa/entity"
	"fmt"
	"time"
)

type DataService interface {
	FindAll() ([]entity.Data, error)
	Save(entity.Data) (entity.Data, error)
}

type dataService struct {
	data []entity.Data
}

func NewData() DataService {
	return &dataService{
		data: []entity.Data{},
	}
}

func (service *dataService) FindAll() ([]entity.Data, error) {
	if len(service.data) <= 0 {
		err := fmt.Errorf("Error: %s", "no available data")
		return service.data, err
	} else {
		return service.data, nil
	}
}

func (service *dataService) Save(d entity.Data) (entity.Data, error) {
	d.MeteData.TimeStamp = time.Now()
	service.data = append(service.data, d)
	return d, nil
}
