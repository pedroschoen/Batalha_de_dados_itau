package model

import (
	"context"
	"fmt"
	// "github.com/go-sql-driver/mysql"
	"googlemaps.github.io/maps"
)

var apikey = "AIzaSyAQudl5tZE7vA6eautkJfPNax1GMYFMX6w"

// Retorna a Rota de acordo com o Logradouro de Origem
func GetRoutes(origin string) ([]maps.Route, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		return nil, fmt.Errorf("Erro gerando rota: %v", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: "RUA DOUTOR VITAL BRASIL",
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		return nil, fmt.Errorf("Erro gerando directions: %v", err)
	}
	return route, nil
}
