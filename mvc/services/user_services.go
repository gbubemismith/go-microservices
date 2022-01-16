package services

import (
	"github.com/gbubemismith/go-microservices/mvc/domain"
	"github.com/gbubemismith/go-microservices/mvc/utils"
)

func GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userid)
}
