package main
import(
	"wennhmann.de/lectionTwo/fill"
	 xyz "wennhmann.de/lectionTwo/init"
	"fmt"
	"time"
	"log"
)

func main(){
	fill.Fill()
	fmt.Println(xyz.Init)
	go fill.PrintNum_0()
	go fill.PrintNum_1()
	go fill.Nummmer()
	

	//chan
	var received int
    ch5 := make(chan int, 12)

    ch5 <- fill.Nummmer()
	received, ok := <-ch5
	if !ok {
    log.Println("channel is empty or closed")
}
  
    log.Println("interterend #####", received,ok)
	time.Sleep(1000 * time.Minute) 
	
}