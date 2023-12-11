package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"golang-testing-project/models"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "postgres"
)

func main() {
	e := echo.New()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	e.GET("/order/:order_code", func(c echo.Context) error {
		orderCode := c.Param("order_code")
		row := db.QueryRow("select order_code , sub_amount from orders where  order_code = $1", orderCode)
		var order models.OrderEntity
		if err := row.Scan(&order.OrderCode, &order.SubAmount); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, order)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
