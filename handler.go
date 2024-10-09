package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func buatItem(c echo.Context) error {
    item := new(Item)
    if err := c.Bind(item); err != nil {
        return err
    }
    
    result, err := db.Exec("INSERT INTO items (nama, harga) VALUES (?, ?)", item.Nama, item.Harga)
    if err != nil {
        return err
    }
    
    id, _ := result.LastInsertId()
    item.ID = int(id)
    
    return c.JSON(http.StatusCreated, item)
}

func getItems(c echo.Context) error {
    rows, err := db.Query("SELECT id, nama, harga FROM items")
    if err != nil {
        return err
    }
    defer rows.Close()

    var items []Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ID, &item.Nama, &item.Harga); err != nil {
            return err
        }
        items = append(items, item)
    }

    return c.JSON(http.StatusOK, items)
}

func getItem(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    
    var item Item
    err := db.QueryRow("SELECT id, nama, harga FROM items WHERE id = ?", id).Scan(&item.ID, &item.Nama, &item.Harga)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "Item tidak ditemukan")
        }
        return err
    }
    
    return c.JSON(http.StatusOK, item)
}

func updateItem(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    
    item := new(Item)
    if err := c.Bind(item); err != nil {
        return err
    }
    
    _, err := db.Exec("UPDATE items SET nama = ?, harga = ? WHERE id = ?", item.Nama, item.Harga, id)
    if err != nil {
        return err
    }
    
    item.ID = id
    return c.JSON(http.StatusOK, item)
}

func hapusItem(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    
    result, err := db.Exec("DELETE FROM items WHERE id = ?", id)
    if err != nil {
        return err
    }
    
    affected, _ := result.RowsAffected()
    if affected == 0 {
        return c.JSON(http.StatusNotFound, "Item tidak ditemukan")
    }
    
    return c.NoContent(http.StatusNoContent)
}