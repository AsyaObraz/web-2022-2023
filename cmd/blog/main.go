package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
	"github.com/jmoiron/sqlx"
)

const (
	port = ":3000"
	dbDriverName = "mysql"
)

func main() {
	db, err := openDB()  // Открываем соединение к базе данных в самом начале
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDB(db, dbDriverName) // Расширяем стандартный клиент к базе

	mux := http.NewServeMux()
	mux.HandleFunc("/home", index(dbx))  // Передаём клиент к базе данных в ф-ию обработчик запроса
	mux.HandleFunc("/post", post)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Start server " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() (*sql.DB, error) {
	//Здесь прописываем соединение к базе данных
	return sql.Open(dbDriverName, "root:P8j#ns6A!N?9i@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}