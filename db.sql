CREATE DATABASE test_db;

use test_db;

CREATE TABLE `seq` (
  `name` varchar(20) NOT NULL,
  `val` int(10) unsigned NOT NULL,
  PRIMARY KEY  (`name`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8
  COLLATE=utf8_general_ci;


CREATE TABLE test_db.oms_orders_1 (
	id varchar(100) NOT NULL,
	invoice_id varchar(100) NOT NULL,
	order_address varchar(100) NULL,
	created_at DATETIME NOT NULL
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8
  COLLATE=utf8_general_ci;


CREATE TABLE test_db.oms_orders_2 (
	id varchar(100) NOT NULL,
	invoice_id varchar(100) NOT NULL,
	order_address varchar(100) NULL,
	created_at DATETIME NOT NULL
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8
  COLLATE=utf8_general_ci;


CREATE TABLE test_db.oms_orders_3 (
	id varchar(100) NOT NULL,
	invoice_id varchar(100) NOT NULL,
	order_address varchar(100) NULL,
	created_at DATETIME NOT NULL
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8
  COLLATE=utf8_general_ci;

CREATE INDEX oms_orders_1_id_IDX USING BTREE ON test_db.oms_orders_1 (id);
CREATE INDEX oms_orders_1_id_IDX USING BTREE ON test_db.oms_orders_2 (id);
CREATE INDEX oms_orders_1_id_IDX USING BTREE ON test_db.oms_orders_3 (id);