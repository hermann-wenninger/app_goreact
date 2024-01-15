package main
import(
	"wennhmann.de/lectionTwo/fill"
	 xyz "wennhmann.de/lectionTwo/init"
	"fmt"
	"time"
)

func main(){
	fill.Fill()
	fmt.Println(xyz.Init)
	go fill.PrintNum_0()
	go fill.PrintNum_1()
	time.Sleep(1 * time.Minute) 
}