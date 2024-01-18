package main
import (
"fmt"
"net/http"
"strings"
"log"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
r.ParseForm() // parse arguments, you have to call this by yourself
fmt.Println("formc",r.Form) // print form information in server side
fmt.Println("pathc", r.URL.Path)
fmt.Println("schemec", r.URL.Scheme)
fmt.Println("the key of form rrr",r.Form["url_long"])
for k, v := range r.Form {
fmt.Println("key:", k)
fmt.Println("val:", strings.Join(v, ""))
}
fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}
func main() {
http.HandleFunc("/", sayhelloName) // set router
err := http.ListenAndServe(":9090", nil) // set listen port
if err != nil {
log.Fatal("ListenAndServe: ", err)
}
}
