package controller

import (
	model "../Model"
	"bytes"
	"io/ioutil"
	// "database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//Controller - Type
type Controller struct{}

//NewController Cria um novo controllers
func NewController() *Controller {
	return &Controller{}
}

//GenerateRoute Retorna a Route para a requisição
func (c *Controller) GenerateRoute(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	route, err := model.GetRoutes("AVENIDA PAULISTA")
	if err != nil {
		http.Error(w, "Erro gerando rotas", http.StatusBadRequest)
		return
	}
	uj, err := json.Marshal(route)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
	}
	fmt.Fprintf(w, string(uj))
}

func (c *Controller) GetRecomendationRoute(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro gerando rota", http.StatusBadGateway)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "localhost:5000/recommend_route", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Erro gerando rota", http.StatusBadGateway)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error gerando rota", http.StatusBadGateway)
		return
	}
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro gerando rota", http.StatusBadGateway)
		return
	}
	fmt.Fprintf(w, string(rbody))
}
