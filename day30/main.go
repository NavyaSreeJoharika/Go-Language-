/*
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "World", "your name")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}

// output ==== Hello, World
*/
/*
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
}
*/
/*
output=====
PS C:\golang\day30>go run main.go
[C:\Users\user\AppData\Local\Temp\go-build908362433\b001\exe\main.exe]
PS C:\golang\day30> go run main.go -age 59
[C:\Users\user\AppData\Local\Temp\go-build91003092\b001\exe\main.exe -age 59]
*/
/*
package main

import (
	"flag"
	"fmt"
)

func main() {
	var age = flag.Int("age", 20, "An age integer")

	flag.Parse()

	fmt.Println(age)

}
*/
/*output===
PS C:\golang\day30> go run main.go
0xc000016088
PS C:\golang\day30> go run main.go -age 59
0xc0000160b8
*/
/*
package main

import (
	"flag"
	"fmt"
)

func main() {
	var age = flag.Int("age", 20, "An age integer")

	flag.Parse()

	fmt.Println(*age)

}
*/
/* output ====
PS C:\golang\day30> go run main.go
20
PS C:\golang\day30> go run main.go -age 50
50
*/
/*
package main

import (
	"flag"
	"fmt"
)

func main() {
	var age int
	flag.IntVar(&age, "age", 20, "An age integer")

	flag.Parse()

	fmt.Println(age)
}
*/
/* output ====
PS C:\golang\day30> go run main.go
20
PS C:\golang\day30> go run main.go -age 30
30
*/
/*
package main

import (
	"flag"
	"fmt"
)

type ServerConfig struct {
	port    int
	workers int
	env     string
}

func main() {
	var serverConfig ServerConfig

	flag.IntVar(&serverConfig.port, "port", 8000, "Server port")
	flag.IntVar(&serverConfig.workers, "workers", 4, "Number of worker processes")
	flag.StringVar(&serverConfig.env, "env", "dev", "Environment")

	flag.Parse()

	fmt.Printf("%+v\n", serverConfig)
}
*/

/* output ====
PS C:\golang\day30> go run main.go
{port:8000 workers:4 env:dev}

PS C:\golang\day30> go run main.go -port 4000
{port:4000 workers:4 env:dev}

PS C:\golang\day30> go run main.go -workers 10
{port:8000 workers:10 env:dev}

PS C:\golang\day30> go run main.go -env prod
{port:8000 workers:4 env:prod}

PS C:\golang\day30> go run main.go -env prod -workers 20
{port:8000 workers:20 env:prod}

PS C:\golang\day30> go run main.go -env prod -port 2000 -workers 5
{port:2000 workers:5 env:prod}
*/

package main

import (
	"flag"
	"fmt"
)

func main() {

	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/* output ====
PS C:\golang\day30> go run main.go
word: foo
numb: 42
fork: false
svar: bar
tail: []

PS C:\golang\day30> go run main.go -word optimistic -numb 10
word: optimistic
numb: 10
fork: false
svar: bar
tail: []

PS C:\golang\day30> go run main.go -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

PS C:\golang\day30> go run main.go -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

PS C:\golang\day30> go run main.go -word=opt a1 a2 a3
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3]

PS C:\golang\day30> go run main.go -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

PS C:\golang\day30> go run main.go -h
Usage of C:\Users\user\AppData\Local\Temp\go-build668941171\b001\exe\main.exe:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")

PS C:\golang\day30> go run main.go -wat
flag provided but not defined: -wat
Usage of C:\Users\user\AppData\Local\Temp\go-build1899557449\b001\exe\main.exe:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
exit status 2
*/
