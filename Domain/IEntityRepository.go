package Domain

import "tu.project/echoTest/Model"

type Repository interface {
	Create(user Model.IEntity) error
	Delete(id string, entity Model.IEntity) error
	Find(entities [] Model.IEntity) error
	FindOne(id string, entity Model.IEntity) error
	InitialTable() error
	InitialRepo(entity Model.IEntity, migrate string)
}