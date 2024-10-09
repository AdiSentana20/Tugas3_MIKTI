package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
    // Inisialisasi database
    initDB()
    defer db.Close()

    // Inisialisasi Echo
    e := echo.New()

    // Route endpoints CRUD
    e.POST("/items", buatItem)
    e.GET("/items", getItems)
    e.GET("/items/:id", getItem)
    e.PUT("/items/:id", updateItem)
    e.DELETE("/items/:id", hapusItem)

    // Jalankan server pada port 8000
    e.Logger.Fatal(e.Start(":8000"))
}