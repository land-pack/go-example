package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:openmysql@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}

	defer db.Close()
	// make sure connection is avaiable
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Person struct {
		Id         int
		First_Name string
		Last_Name  string
	}

	router := gin.Default()
	// Add API handles here

	// GET a person detail
	router.GET("/person/:id", func(c *gin.Context) {
		var (
			person Person
			result gin.H
		)

		id := c.Param("id")
		fmt.Println("The id is ", id)
		row := db.QueryRow("select id, first_name, last_name from person where id = ?;", id)
		err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)
		if err != nil {
			fmt.Println(err)
			// If no results  send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	router.Run(":3000")
}
