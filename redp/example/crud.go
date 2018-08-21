package main

import (
    "database/sql"
    "log"
    _"github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
 //   "encoding/json"
)

type Account struct {
    Balance string  `json:"balance" form:"balance"`
    Balance_freeze string `json:"balance_freeze" form:"balance_freeze"`
    Coin_type int `json:"coin_type" form:"coin_type"`
    Coin_name string `json:"coin_name" form:"coin_name"`
}


type Response struct {
   Code    string      `json:"code"`
   Message string      `json:"msg"`
   Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code string, message string, data interface{}) {
   //code, message := errors.DecodeErr(err)

   c.JSON(http.StatusOK, Response{
      Code:    code,
      Message: message,
      Data:    data,
   })
}

func main() {
    router := gin.Default()
    db, err := sql.Open("mysql", "frank:openmysql@tcp(192.168.1.203:3306)/test?parseTime=true")
    if err != nil {
        log.Fatalln(err)
    }
    defer db.Close()

    router.POST("/api/v1/accounts", func(c *gin.Context){
    accounts := make([]Account, 0)

    fetch_accounts_sql := `
        SELECT f_balance as balance, f_balance_freeze as balance_freeze,
            f_coin_type as coin_type, f_coin_name as coin_name
        FROM t_account
        WHERE f_uid=? AND f_status=0`
    uid := 55555
    rows, err := db.Query(fetch_accounts_sql, uid)
    if err != nil {
        log.Fatalln(err)
    }
    for rows.Next() {
        var account Account
        rows.Scan(&account.Balance, &account.Balance_freeze, &account.Coin_type, &account.Coin_name)
        accounts = append(accounts, account)
    }
    fmt.Println(accounts)
    if err = rows.Err(); err != nil {
        log.Fatalln(err)
    }

    //b, err := json.Marshal(accounts)
    //if err != nil {
    //    log.Fatalln(err)
    // }
    //code :=  "200"
    SendResponse(c, "200", "OK", accounts)
  })
  router.Run(":8000")
}
