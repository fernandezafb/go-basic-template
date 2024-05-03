package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Item struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Quantity uint64 `json:"quantity"`
	SellerId string `json:"sellerId"`
	Price    int64  `json:"price"`
}

var (
	items       = map[int64]*Item{}
	seq   int64 = 1
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		status := struct {
			Status string `json:"status"`
		}{Status: "healthy"}

		return c.JSON(http.StatusOK, status)
	})

	e.POST("/items", createItem)
	e.GET("/items/:id", getItem)
	e.PUT("/items/:id", updateItem)
	e.DELETE("/items/:id", deleteItem)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

//----------
// Handlers
//----------

func createItem(c echo.Context) error {
	item := &Item{
		Id: seq,
	}

	if err := c.Bind(item); err != nil {
		return err
	}

	items[item.Id] = item
	seq++

	return c.JSON(http.StatusCreated, item)
}

func getItem(c echo.Context) error {
	// Item Id from path `items/:id`
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	return c.JSON(http.StatusOK, items[id])
}

func updateItem(c echo.Context) error {
	u := new(Item)

	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	items[id].Name = u.Name
	items[id].Quantity = u.Quantity
	items[id].SellerId = u.SellerId
	items[id].Price = u.Price

	return c.JSON(http.StatusOK, items[id])
}

func deleteItem(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	delete(items, id)

	return c.NoContent(http.StatusNoContent)
}
