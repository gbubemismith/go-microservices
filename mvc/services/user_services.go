package services

import (
	"fmt"

	"github.com/gbubemismith/go-microservices/mvc/domain"
)

func GetUser(userid int64) (*domain.User, error) {
	fmt.Println("userid", userid)
	return domain.GetUser(userid)
}
