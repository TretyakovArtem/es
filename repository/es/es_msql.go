package esRepository

import (
	"fmt"
	"time"

	goqu "gopkg.in/doug-martin/goqu.v5"
)

var db *goqu.Database

const firstDistrict, secondDistrict, thirdDistrict int = 1, 2, 3

type EsRepository struct{}

func (r EsRepository) GenerateSeq() {

	// инициализация хранимой процедуры
	r.CreateProc()

	districtId := 1

	for i := 0; i < 10000; i++ {

		if districtId == firstDistrict {
			_, err2 := db.Exec("INSERT INTO orders_1(id, invoice_id, order_address, created_at) VALUES (seq('orders'), ?, ?, ?)", 791772, "ул. Мира, д. 1", time.Now())

			if err2 != nil {
				fmt.Println(err2.Error())
			}
		}

		if districtId == secondDistrict {
			_, err2 := db.Exec("INSERT INTO orders_2(id, invoice_id, order_address, created_at) VALUES (seq('orders'), ?, ?, ?)", 791772, "ул. Мира, д. 1", time.Now())
			if err2 != nil {
				fmt.Println(err2.Error())
			}
		}

		if districtId == thirdDistrict {
			_, err2 := db.Exec("INSERT INTO orders_3(id, invoice_id, order_address, created_at) VALUES (seq('orders'), ?, ?, ?)", 791772, "ул. Мира, д. 1", time.Now())
			if err2 != nil {
				fmt.Println(err2.Error())
			}
		}

		if districtId == 3 {
			districtId = 1
		} else {
			districtId++
		}
	}
}

func (r EsRepository) CreateProc() {
	_, err := db.Exec("create function seq(seq_name char (20)) returns int begin update seq set val=last_insert_id(val+1) where name=seq_name;return last_insert_id();end;")

	if err != nil {
		fmt.Println(err.Error())
	}
}
