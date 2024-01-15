package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=23hermann75@#* sslmode=disable")
    if err != nil {
        fmt.Println(err)
        return
    }
    err = db.Ping()
    if err != nil {
        fmt.Println(err)
        return
    }
}