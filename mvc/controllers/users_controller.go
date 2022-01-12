package controllers

import (
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return
	}
	
}
