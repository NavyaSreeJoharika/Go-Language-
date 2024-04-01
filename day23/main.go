/*
package main

import (
	"fmt"
)

func add(z, y int) (c int) {
	c = z + y
	return
}

func main() {
	fmt.Println(add(23, 55))

	x := "a"
	switch x {
	case "a":
		fmt.Println("Hi")

	case "b":
		fmt.Println("Hello")

	case "c":
		fmt.Println(" Welcome")

	}

	var apple = 98

	switch apple {
	case 1, 3, 5, 7, 9:
		fmt.Println("apple are odd number")

	case 2, 4, 6, 8:
		fmt.Println("apples are even")

	case 0:
		fmt.Println("no apple")

	default:
		fmt.Println("more than 9 apples")
	}

}
*/
/*
//waitgroup
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	words := []string{"foo", "bar", "baz"}

	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			time.Sleep(1 * time.Second)
			defer wg.Done()
			fmt.Println(word)
		}(word)
	}

	wg.Wait()
}
*/
/*
//channel
package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{"foo", "bar", "baz"}
	done := make(chan bool)

	for _, word := range words {

		go func(word string) {
			time.Sleep(1 * time.Second)
			fmt.Println(word)
			done <- true
		}(word)
	}

	for range words {
		<-done
	}
}
*/
/*
//WAITGROUP
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	words := []string{"foo", "bar", "baz"}
	var wg sync.WaitGroup
	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond) // a job
			fmt.Println(word)
		}(word)
	}
	wg.Wait()
}
*/
/*
//channels can be buffered
package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{"foo", "bar", "baz"}
	done := make(chan struct{}, len(words))
	for _, word := range words {
		go func(word string) {
			time.Sleep(100 * time.Millisecond) // a job
			fmt.Println(word)
			done <- struct{}{} // not blocking
		}(word)
	}
	for range words {
		<-done
	}
}
*/
/*
//channels can be unbuffered

package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{"foo", "bar", "baz"}
	done := make(chan struct{})
	for _, word := range words {
		go func(word string) {
			time.Sleep(100 * time.Millisecond) // a job
			fmt.Println(word)
			done <- struct{}{} // blocking
		}(word)
	}
	for range words {
		<-done
	}
}
*/
/*
// You may limit the number of concurrent jobs with buffered channel capacity
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t0 := time.Now()
	var wg sync.WaitGroup
	words := []string{"foo", "bar", "baz"}
	done := make(chan struct{}, 1) // set the number of concurrent job here
	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			done <- struct{}{}
			time.Sleep(100 * time.Millisecond) // job
			fmt.Println(word, time.Since(t0))
			<-done
			wg.Done()
			}(word)
		}
		wg.Wait()
	}

*/

//message using channel

package main

import "fmt"

func main() {
	done := make(chan string)
	go func() {
		for _, word := range []string{"foo", "bar", "baz"} {
			done <- word
		}
		close(done)
	}()
	for word := range done {
		fmt.Println(word)
	}
}
