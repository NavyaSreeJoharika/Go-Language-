/*

package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
	//"github.com/blackflagsoftware/debug_tut/util"
	"database/sql"

    _ "github.com/lib/pq"
)

func main() {
	//addTwoNumbers(10, 5)


	//fmt.Println("What is your name:")
	//name := util.ParseInput()
	//sayHello(name)
	//fmt.Println("Pick a number between 1 -100: ")
	//idxStr := util.ParseInput()
	//idx, _ := strconv.Atoi(idxStr)
	//random := pickRandom(idx - 1)
	//fmt.Println("your random number is :", random)
}

func addTwoNumbers(a, b int) int {
	one := a
	two := b
	total := one + two
	return total
}

func sayHello(name string) (msgPrinted bool) {
	msg := fmt.Sprintf("Hello %s,how are you today ", name)
	fmt.Println(msg)
	msgPrinted = true
	return
}

func pickRandom(idx int) int {
	rand.Seed(time.Now().Unix())
	ints := []int{}
	for i := 0; i < 100; i++ {
		ints = append(ints, rand.Intn(100))
	}
	return ints[idx]
}
*/
/*
package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(2, 9))
}
*/

/*
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	count := 0

	for {
		if count == 10 {
			break
		}
		count++
	}
	fmt.Println(count)
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Avenger represents a single hero
type Avenger struct {
	RealName string `json:"real_name"`
	HeroName string `json:"hero_name"`
	Planet   string `json:"planet"`
	Alive    bool   `json:"alive"`
}

func (a *Avenger) isAlive() {
	a.Alive = true
}

func main() {
	avengers := []Avenger{
		{
			RealName: "Dr. Bruce Banner",
			HeroName: "Hulk",
			Planet:   "Midgard",
		},
		{
			RealName: "Tony Stark",
			HeroName: "Iron Man",
			Planet:   "Midgard",
		},
		{
			RealName: "Thor Odinson",
			HeroName: "Thor",
			Planet:   "Midgard",
		},
	}

	avengers[1].isAlive()

	jsonBytes, err := json.Marshal(avengers)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("This is an error")
	fmt.Println(err)

	err2 := fmt.Errorf("This error wraps the first one:%w", err)
	fmt.Println(err2)
}
