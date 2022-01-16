package app

import (
	"net/http"

	"github.com/gbubemismith/go-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	//if initialization; condition {
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
