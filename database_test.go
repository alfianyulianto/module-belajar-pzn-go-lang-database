package _6_7__pzn_go_lang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // tujuan import ini memanggil init function di driver
	"testing"
)

func TestEmpty(t *testing.T) {

}

// Membuat Koneksi ke Database

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close() // untuk mematikan koneksi database

	// gunakan database

}
