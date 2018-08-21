package main

import (
	"database/sql"
	"log"
	// TODO What this `_` mean?
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	drop_t_account := "DROP TABLE IF EXISTS t_account"

	create_t_account := `
    CREATE TABLE t_account(
        f_id BIGINT NOT NULL AUTO_INCREMENT,
        f_uid BIGINT NOT NULL,
        f_coin_type INT NOT NULL DEFAULT 3545,
        f_coin_name CHAR(64) NOT NULL DEFAULT 'SEELE',
        f_balance DECIMAL(16,8) NOT NULL DEFAULT 0,
        f_balance_freeze DECIMAL(16,8) NOT NULL DEFAULT 0,
        f_status int NOT NULL DEFAULT 0,
        f_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
        f_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
                     ON UPDATE CURRENT_TIMESTAMP, 
        unique index uid_coin_type(f_uid, f_coin_type),
        PRIMARY KEY (f_id))`

	db, err := sql.Open("mysql", "frank:openmysql@tcp(192.168.1.203:3306)/gotest?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.Exec(drop_t_account)
	stmt, err := db.Prepare(create_t_account)
	if err != nil {
		log.Fatalln(err)
	}
	stmt.Exec()
}
