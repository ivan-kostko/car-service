package services

import (
	"./../../models"
	"errors"
	log "./../../pkg/logger"
)

const(
	INVALID_ID_ERR := errors.NewError("Car id should be in range 1-2^64")
)

type carByIdGetter interface {
	GetCarById(int) (*models.CarEntity, error)
}

type CarService struct {
	carByIdGetter
	log.
}

func NewService() *CarService {
	return &CarService{}
}

(cs *CarService) func GetCarById(id int) (*models.CarEntity, error) {
	
	if id <= 0 || id > 2^64 {
		return nil, INVALID_ID_ERR
	}
	return cs.carByIdGetter.GetCarById(id)
}