package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gbubemismith/go-microservices/mvc/services"
	"github.com/gbubemismith/go-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")

	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, userErr := services.GetUser(userId)
	if userErr != nil {
		resp.WriteHeader(userErr.StatusCode)
		resp.Write([]byte(userErr.Message))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
