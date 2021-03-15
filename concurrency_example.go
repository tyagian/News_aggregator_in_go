//same to practice concurrency before applying to our crawler
package main
import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i <3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	go say("Hey")
	go say("There")
	time.Sleep(time.Second)
}