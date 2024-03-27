/*
package main

import (

	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

)

	func main() {
		runtime.GOMAXPROCS(4)

		var counter uint64
		var wg sync.WaitGroup

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for c := 0; c < 1000; c++ {
					atomic.AddUint64(&counter, 1)
				}
			}()
		}
		wg.Wait()
		fmt.Println("counter:", counter)
	}
*/
/*
package main

import "fmt"

func main() {

	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)
}
*/
/*
package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}
*/
/*
package main

import "fmt"

func myfunc(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "Hello"
	}
	close(mychnl)
}

func main() {
	c := make(chan string)

	go myfunc(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close", ok)
			break
		}
		fmt.Println("Channel Open", res, ok)
	}
}
*/
/*
package main

import (
	"fmt"
)

func main() {

	mychnl := make(chan string)

	go func() {
		mychnl <- "Hello"
		mychnl <- "Welcome"
		mychnl <- "How"
		mychnl <- "Are"
		mychnl <- "You"
		close(mychnl)
	}()

	for res := range mychnl {
		fmt.Println(res)
	}
}
*/
/*
package main

import (
	"fmt"
)

func main() {

	mychnl := make(chan string, 4)

	mychnl <- "Hello"
	mychnl <- "Welcome"
	mychnl <- "How"
	mychnl <- "Are"
	mychnl <- "You"

	fmt.Println("Length of the channel is:", len(mychnl))
}
*/
/*
package main

import (
	"fmt"
)

func main() {

	mychnl := make(chan string, 5)

	mychnl <- "Hello"
	mychnl <- "Welcome"
	mychnl <- "How"
	mychnl <- "Are"
	mychnl <- "You"

	fmt.Println("Length of the channel is:", len(mychnl))
}
*/
/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)

	go func() {
		ch <- "the message"
	}()

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}
*/
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch2 <- "message to channel 2"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	time.Sleep(10 * time.Millisecond)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
*/

package main

func main() {
}

func Add(l, r int) int {
	return l + r
}
