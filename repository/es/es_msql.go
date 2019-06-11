package esrepository

import (
	"fmt"
	"time"

	goqu "gopkg.in/doug-martin/goqu.v5"
)

const firstDistrict, secondDistrict, thirdDistrict int = 1, 2, 3

type EsRepository struct{}

func (r EsRepository) GenerateSeq(db *goqu.Database) {

	// инициализация хранимой процедуры
	r.CreateProc(db)

	districtId := 1

	for i := 0; i < 10000; i++ {

		if districtId == firstDistrict {
			_, err := db.Exec("INSERT INTO oms_orders_1(id, invoice_id, order_address, created_at) VALUES (seq('orders'), ?, ?, ?)", 791772, "ул. Мира, д. 1", time.Now())

			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if districtId == secondDistrict {
			_, err := db.Exec("INSERT INTO oms_orders_2(id, invoice_id, order_address, created_at) VALUES (seq('orders'), ?, ?, ?)", 791772, "ул. Мира, д. 1", time.Now())
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if districtId == thirdDistrict {
			_, err := db.Exec("INSERT INTO oms_orders_3(id, invoice_id, order_address, created_at) VALUES (seq('orders'), ?, ?, ?)", 791772, "ул. Мира, д. 1", time.Now())
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if districtId == thirdDistrict {
			districtId = firstDistrict
		} else {
			districtId++
		}
	}
}

func (r EsRepository) CreateProc(db *goqu.Database) {
	_, err := db.Exec("create function seq(seq_name char (20)) returns int begin update seq set val=last_insert_id(val+1) where name=seq_name;return last_insert_id();end;")

	if err != nil {
		fmt.Println(err.Error())
	}
}
