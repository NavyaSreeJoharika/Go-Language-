/*package main

import "net/http"

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
}
*/

//Channels
/*
package main

import "fmt"

func main() {

	dataChan := make(chan int)
	dataChan <- 789
	n := <-dataChan
	fmt.Printf("n =%d\n", n)//dead lock
}
*/
/*
package main

import "fmt"

func main() {

	dataChan := make(chan int)
	go func() {
		dataChan <- 789
	}()
	n := <-dataChan
	fmt.Printf("n %d\n", n)
}
*/
/*
package main

import "fmt"

func main() {

	dataChan := make(chan int, 1)
	dataChan <- 789
	n := <-dataChan
	fmt.Printf("n =%d\n", n) //buffer channel

}

*/
/*
package main

import "fmt"

func main() {

	dataChan := make(chan int, 2)
	dataChan <- 789
	dataChan <- 1235
	n := <-dataChan
	fmt.Printf("n =%d\n", n)
	n = <-dataChan
	fmt.Printf("n =%d\n", n)

}

*/

/*
package main

import "fmt"

func main() {

	dataChan := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n=%d\n", n)
	}
}
*/
/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {

	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n=%d\n", n)
	}
}
*/

//comtext

/*
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	imageDate, err := ioutil.ReadAll(res.Body)

	fmt.Printf("download image of size %d\n", len(imageDate))
}
*/
/*
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	imageDate, err := ioutil.ReadAll(res.Body)

	fmt.Printf("download image of size %d\n", len(imageDate))
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
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("IDX from first func:", i)
			//	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("IDX from second func:", i)
			//time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("Done")
}
