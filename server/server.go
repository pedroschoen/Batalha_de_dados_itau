package main

import (
	controller "./Controller"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {

}

func main() {
	ctrl := controller.NewController()
	h := httprouter.New()
	h.GET("/generate_route", ctrl.GenerateRoute)
	fmt.Println("Servindo...")
	http.ListenAndServe(":4000", h)
}
