package controllers

import (
	"fmt"
	"net/http"

	esRepository "github.com/es/repository/es"

	goqu "gopkg.in/doug-martin/goqu.v5"
)

type Controller struct{}

func (c Controller) GenerateSeq(db *goqu.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		esRepo := esRepository.EsRepository{}

		esRepo.GenerateSeq(db)
		fmt.Println("done")
	}
}
