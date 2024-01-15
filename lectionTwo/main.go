package main
import(
	"wennhmann.de/lectionTwo/fill"
	 xyz "wennhmann.de/lectionTwo/init"
	"fmt"
)

func main(){
	fill.Fill()
	fmt.Println(xyz.Init)
	 go fill.PrintNum_0()
	 fill.PrintNum_1()
}