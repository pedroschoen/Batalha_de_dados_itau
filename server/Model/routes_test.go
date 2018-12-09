package model

import (
	"log"
	"testing"
)

func TestRoutes(t *testing.T) {
	err := GetRoutes("AVENIDA PAULISTA")
	if err != nil {
		log.Fatal(err)
	}
}
