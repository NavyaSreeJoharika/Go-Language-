/*
package main

import (
	"fmt"
	"sync"
)

var GIF = 0

func worker(wg *sync.WaitGroup) {
	GIF = GIF + 1

	wg.Done()
}
func main() {

	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w)
	}

	w.Wait()
	fmt.Println("Value of x", GIF)
}
*/
/*
package main

import (
	"fmt"
	"sync"
)

var GFG = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {

	m.Lock()
	GFG = GFG + 1
	m.Unlock()

	wg.Done()
}
func main() {
	var w sync.WaitGroup

	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w, &m)
	}

	w.Wait()
	fmt.Println("Value of x", GFG)
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for raing", input)
}
*/
/*
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "Critical", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./go.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}
*/
/*
package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbcall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal ececution time:%v", time.Since(t0))
	fmt.Printf("\nThe result are %v", results)
}

func dbcall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}
*/
/*
package main

import (
	"fmt"
	"sync"
	"time"
)

// var m = sync.Mutex{}
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbcall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal ececution time:%v", time.Since(t0))
	fmt.Printf("\nThe result are %v", results)
}

func dbcall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()

	wg.Done()
}

func save(result string) {

	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
*/
/*
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4"}

func main() {
	t0 := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go dbcall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal ececution time:%v", time.Since(t0))

}

func dbcall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	wg.Done()
}
*/
/*
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4"}

func main() {
	t0 := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\nTotal ececution time:%v", time.Since(t0))

}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	a := 0

	lock := sync.RWMutex{}

	for i := 1; i < 10; i++ {
		go func(i int) {
			lock.Lock()
			fmt.Printf("Lock: from go routine %d: a = %d\n", i, a)
			time.Sleep(time.Second)
			lock.Unlock()
		}(i)
	}

	b := 0

	for i := 11; i < 20; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("RLock: from go routine %d: b = %d\n", i, b)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	<-time.After(time.Second * 10)
}
