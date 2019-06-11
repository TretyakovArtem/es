package controllers

import (
	"fmt"
	"net/http"

	esRepository "esequence/repository/es"

	goqu "gopkg.in/doug-martin/goqu.v5"
)

type Controller struct{}

func (c Controller) GenerateSeq(db *goqu.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		esRepo := esRepository.EsRepository{}

		esRepo.GenerateSeq(db, book)
		fmt.Println("done")
	}
}
