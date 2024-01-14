package main

import (
	"bytes"
	"fmt"
	"os"
	"io"
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
}