/*
package main

import (

	"fmt"
	"sync"

)

	func main() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			fmt.Println("do some async thing")
			wg.Done()

		}()
		wg.Wait()
	}
*/
/*
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the goroutine completes
	fmt.Printf("Worker %d starting\n", id)
	// Simulate doing some work
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d working...\n", id)
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers have finished")
}
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done()
	}()
	fmt.Println("This is synchronous")
	wg.Wait()
}

// wait group
/*
package main

import (
	"fmt"
)

func main() {
	go greeter("hello")
	greeter("world")

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}
*/
/*
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
*/
/*
package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am first runner")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go runner1(wg)
	go runner2(wg)

	wg.Wait()
}

func main() {
	execute()
}
*/
/*
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {

		result := job * 2
		results <- result
	}
}

func main() {
	numJobs := 10
	numWorkers := 4

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	var wg sync.WaitGroup
	wg.Add(numJobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
		wg.Done()
	}
}
*/
/*
package main

import (
	"fmt"
	"sync"
)

var GFG = 0

func worker(wg *sync.WaitGroup) {
	GFG = GFG + 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w)
	}

	w.Wait()
	fmt.Println("Value of x :", GFG)
}
*/
/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d stating\n", id)

	time.Sleep(time.Millisecond)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
*/
