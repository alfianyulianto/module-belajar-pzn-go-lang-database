package _6_7__pzn_go_lang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Database Pooling
func GetConnection() *sql.DB {
	// parseTime untuk konversi tipe data DATE, DATETIME, TIMESTAMP
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // jumlah koneksi minimum
	db.SetMaxOpenConns(100)                 // jumlah koneksi maksimum
	db.SetConnMaxIdleTime(5 * time.Minute)  // berapa lama koneksi tidak bekerja, jika setelah 5 menit maka langsung dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // berapa lama koneksi digunakan, kemudian akan diperbarui dengan koneksi baru

	return db
}
