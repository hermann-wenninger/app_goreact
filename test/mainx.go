package main

import (
	"bytes"
	"fmt"
	"os"
	"io"
	"database/sql"
)

func main() {
    s := [][]string{
        {"age", "genre", "name"},
        {"23", "M", "Hendrick"},
        {"65", "F", "Stephany"},
    }
    // Open the file or create it
    f, err := os.Create("myFile.csv")
    defer f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    var buffer bytes.Buffer
    // iterate over the slice
    for _, data := range s {
        buffer.WriteString(fmt.Sprintf("%s,%s,%s\n", data[0], data[1], data[2]))
    }
    n, err := f.Write(buffer.Bytes())
    fmt.Printf("%d bytes written\n", n)
    if err != nil {
        fmt.Println(err)
        return
    }
	x, errr := os.Open("teacher.sql")
if errr != nil {
    fmt.Println(errr)
    return
}
b, err := io.ReadAll(x)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(b)

res, err := db.Exec(string(b))
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)

func createTeacher(firstname string, lastname string, db *sql.DB) (int64, error) {
    res, err := db.Exec("INSERT INTO `teacher` (`create_time`, `firstname`, `lastname`) VALUES (NOW(), ?, ?)", firstname, lastname)
    if err != nil {
        return 0, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return 0, err
    }
    return id, nil
}
xy:= createTeacher("hermann","podgorni")
fmt.Println(xy)

func teachers(db *sql.DB) (*[]Teacher, error) {
    rows, err := db.Query("SELECT id, firstname, lastname FROM teacher ")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    teachers := make([]Teacher, 0)
    for rows.Next() {
        teacher := Teacher{}
        if err := rows.Scan(&teacher.id, &teacher.firstname, &teacher.lastname); err != nil {
            return nil, err
        }
        teachers = append(teachers, teacher)
    }
    return &teachers, nil
}
teacher()
}