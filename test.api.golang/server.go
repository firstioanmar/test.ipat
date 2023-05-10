package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Produk struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

func main() {

	// database connection
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_ipat")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// database connection

	e := echo.New()

	// Enable CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Service API!")
	})

	// Produk
	e.GET("/produk", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM produk")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var produk []Produk
		for res.Next() {
			var m Produk
			_ = res.Scan(&m.ID, &m.Nama, &m.Deskripsi)
			produk = append(produk, m)
		}

		return c.JSON(http.StatusOK, produk)
	})

	e.POST("/produk", func(c echo.Context) error {
		var produk Produk
		c.Bind(&produk)

		sqlStatement := "INSERT INTO produk (id,nama, deskripsi)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, produk.ID, produk.Nama, produk.Deskripsi)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, produk)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/produk/:id", func(c echo.Context) error {
		var produk Produk
		c.Bind(&produk)

		sqlStatement := "UPDATE produk SET nama=?,deskripsi=? WHERE id=?"
		res, err := db.Query(sqlStatement, produk.Nama, produk.Deskripsi, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, produk)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/produk/:id", func(c echo.Context) error {
		var produk Produk
		c.Bind(&produk)

		sqlStatement := "DELETE FROM produk WHERE id=?"
		res, err := db.Query(sqlStatement, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, produk)
		}
		return c.String(http.StatusOK, "ok")
	})
	// Produk

	e.Logger.Fatal(e.Start(":8100"))
}
