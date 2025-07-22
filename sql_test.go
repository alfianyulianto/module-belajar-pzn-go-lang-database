package _6_7__pzn_go_lang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Mengeksekusi Perintah SQL
func TestExceSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// ExceContext tidak mengembalikan data, jadi kita tidak bisa menggunakan query data seperti SELECT
	// ExecContext dapat digunakan untuk INSERT, UPDATE, DELETE
	_, err := db.ExecContext(ctx, "INSERT INTO customer VALUES ('alfian', 'Alfian')")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

// Query SQL

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// rows ini berisi data
	rows, err := db.QueryContext(ctx, "SELECT * FROM customer")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name) // biasanya pointer, karena set variable id, name dari hasil query
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id, ", Name :", name)

	}

	defer rows.Close() // jika harus menutup rows
}

// Tipe Data Column
func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// rows ini berisi data
	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name string
		var email sql.NullString // Nullable value untuk string; berupa struct{String string, Valid bool); jika valid maka bernilai true dan tidak kosong, jika tidak valid bernilai false atau kosong
		var balance int32
		var rating float64
		var birthDate sql.NullTime // Nullable value untuk time; berupa struct {Time time.Time, Valid bool}; jika valid maka bernilai true dan tidak kosong, jika tidak valid bernilai false atau kosong
		var createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) // biasanya pointer, karena set variable id, name dari hasil query
		if err != nil {
			panic(err)
		}
		fmt.Println("=================")
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthDate.Valid {
			fmt.Println("BirthDate :", birthDate.Time)
		}
		fmt.Println("Married :", married)
		fmt.Println("CreatedAt :", createdAt)
		fmt.Println("=================")
	}

	defer rows.Close() // jika harus menutup rows
}

// SQL Injection
func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "'LIMIT 1"
	fmt.Println(script) // SELECT username FROM user WHERE username = 'admin'; #' AND password = 'salah'LIMIT 1 (ini sql injection)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login :", username)
	} else {
		fmt.Println("Gagal login")
	}
}

// SQL Dengan Parameter
func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//username := "admin"
	username := "admin'; #"
	password := "admin"

	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(sqlQuery) // SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login :", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestExceSQLParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "alfian; DROP TABLE USER'; #"
	password := "alfian"

	// ExceContext tidak mengembalikan data, jadi kita tidak bisa menggunakan query data seperti SELECT
	// ExecContext dapat digunakan untuk INSERT, UPDATE, DELETE
	_, err := db.ExecContext(ctx, "INSERT INTO user VALUES (?, ?)", username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

// Auto Increment
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "alfian@mail.com"
	comment := "Ini belajar golang database"

	// ExceContext tidak mengembalikan data, jadi kita tidak bisa menggunakan query data seperti SELECT
	// ExecContext dapat digunakan untuk INSERT, UPDATE, DELETE
	result, err := db.ExecContext(ctx, "INSERT INTO comments (email, comment) VALUES (?, ?)", email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}

// Prepare Statement

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	// dengan statement kita bisa menggunakan koneksi yang sama ketika proses Exec atau Query
	statment, err := db.PrepareContext(ctx, "INSERT INTO comments (email, comment) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	defer statment.Close()

	for i := 1; i <= 10; i++ {
		email := "alfian" + strconv.Itoa(i) + "@mail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)
		result, err := statment.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id : ", lastInsertId)
	}
}

// Database Trasaction
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// jika menggunakan tx maka koneksi ke database masih sama (seperti Prepare Staetement yang emnggunakan koneksi yang sama)
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// kode program transaksi kita
	// misal insert ke dua buah table

	for i := 1; i <= 10; i++ {
		email := "alfian" + strconv.Itoa(i) + "@mail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, "INSERT INTO comments (email, comment) VALUES (?, ?)", email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id : ", lastInsertId)

		// Menggagalkan transaksi
		if i == 2 {
			tx.Rollback()
			panic("Transaksi gagal")
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
