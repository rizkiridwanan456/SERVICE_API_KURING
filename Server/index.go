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

type Saung struct {
	No        string `json:"no"`
	Kapasitas string `json:"kapasitas"`
	Status    string `json:"status"`
	Untuk     string `json:"untuk"`
}
type Booking struct {
	Nomor  string `json:"nobo"`
	Nama   string `json:"nama"`
	Saung  string `json:"no"`
	Jumlah string `json:"orang"`
}
type SA struct {
	No        string `json:"no"`
	Kapasitas string `json:"kapasitas"`
	Status    string `json:"status"`
}
type pesan struct {
	KodePesan string `json:"kp"`
	Paket     string `json:"nopa"`
	Nama      string `json:"nama"`
	Jumlah    string `json:"jumlah"`
}
type report struct {
	KodePesan string `json:"kp"`
	Paket     string `json:"nopa"`
	Nama      string `json:"nama"`
	Jumlah    string `json:"jumlah"`
}
type makanan struct {
	ID    string `json:"nopa"`
	Paket string `json:"paket"`
	Harga string `json:"harga"`
	Stok  string `json:"stok"`
}

func main() {

	// database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kuring")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// database connection

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hai, Service API!")
	})

	// Kuring
	e.GET("/fb", func(c echo.Context) error {
		res, err := db.Query("CALL fb")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []makanan
		for res.Next() {
			var m makanan
			_ = res.Scan(&m.ID, &m.Paket, &m.Harga, &m.Stok)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})
	e.GET("/saung", func(c echo.Context) error {
		res, err := db.Query("CALL saung")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []SA
		for res.Next() {
			var m SA
			_ = res.Scan(&m.No, &m.Kapasitas, &m.Status)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})
	e.GET("/pesan", func(c echo.Context) error {
		res, err := db.Query("CALL pesan")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []pesan
		for res.Next() {
			var m pesan
			_ = res.Scan(&m.KodePesan, &m.Paket, &m.Nama, &m.Jumlah)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})
	e.GET("/cari/:Nama", func(c echo.Context) error {
		res, err := db.Query("CALL Look", c.Param("Nama"))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Close()

		var mahasiswa []pesan
		for res.Next() {
			var m pesan
			err := res.Scan(&m.KodePesan, &m.Paket, &m.Nama, &m.Jumlah)
			if err != nil {
				log.Fatal(err)
			}
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})
	e.GET("/booking", func(c echo.Context) error {
		res, err := db.Query("CALL booking")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []Booking
		for res.Next() {
			var m Booking
			_ = res.Scan(&m.Nomor, &m.Nama, &m.Saung, &m.Jumlah)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})
	e.GET("/Lihat/:Nama", func(c echo.Context) error {
		res, err := db.Query("CALL Lihat", c.Param("Nama"))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Close()

		var mahasiswa []Booking
		for res.Next() {
			var m Booking
			err := res.Scan(&m.Nomor, &m.Nama, &m.Saung, &m.Jumlah)
			if err != nil {
				log.Fatal(err)
			}
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})
	e.GET("/laporan", func(c echo.Context) error {
		res, err := db.Query("CALL laporan")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []Booking
		for res.Next() {
			var m Booking
			_ = res.Scan(&m.Nomor, &m.Nama, &m.Saung, &m.Jumlah)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})

	e.POST("/order", func(c echo.Context) error {
		var mahasiswa pesan
		c.Bind(&mahasiswa)

		sqlStatement := "CALL insertpesan"
		res, err := db.Query(sqlStatement, mahasiswa.KodePesan, mahasiswa.Paket, mahasiswa.Nama, mahasiswa.Jumlah)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	e.POST("/book", func(c echo.Context) error {
		var mahasiswa Booking
		c.Bind(&mahasiswa)

		sqlStatement := "CALL insertbooking"
		res, err := db.Query(sqlStatement, mahasiswa.Nomor, mahasiswa.Nama, mahasiswa.Saung, mahasiswa.Jumlah)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/saung/:no", func(c echo.Context) error {
		var mahasiswa SA
		c.Bind(&mahasiswa)

		sqlStatement := "CALL updatesaung"
		res, err := db.Query(sqlStatement, c.Param("no"), mahasiswa.Status)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// Mahasiswa

	e.Logger.Fatal(e.Start(":1323"))
}
