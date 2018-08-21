package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
    "fmt"
	"net/http"
	// TODO What this `_` mean?
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "frank:openmysql@tcp(192.168.1.203:3306)/test?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// New a web server
	router := gin.Default()

	sql_a := `CREATE TABLE t_person(
                            f_id INT NOT NULL AUTO_INCREMENT,
                            f_first_name VARCHAR(40) NOT NULL DEFAULT '',
                            f_last_name VARCHAR(40) NOT NULL DEFAULT '',
                            PRIMARY KEY(f_id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
    fmt.Println(sql_a)
    stmt, err := db.Prepare(sql_a)
	if err != nil {
		log.Fatalln(err)
	}
	stmt.Exec()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It's Okay!")
	})

	router.Run(":8000")

}
