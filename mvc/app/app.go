package app

import (
	"net/http"

	"github.com/gbubemismith/go-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
