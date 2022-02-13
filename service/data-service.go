package service

import "dfa/entity"

type DataService interface {
	FindAll() ([]entity.Data, error)
	FindOne(string) (entity.Data, error)
	Save(entity.Data) (entity.Data, error)
	Autorise(string)
}
