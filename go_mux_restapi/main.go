// main.go

package main

//import "os"

func main() {
	a := App{}
	a.Initialize("postgres","23hermann75@#*","postgres")

	a.Run(":8080")
}
