/*
package main

import (

	"fmt"
	"io"
	"io/ioutil"
	"os"

)

	func main() {
		fmt.Println("Hi")
		content := "This need to go in file - LearnCodeLine.in"

		file, err := os.Create("./mylcogofile.txt")
		if err != nil {
			panic(err)

			length, err := io.WriteString(file, content)
			if err != nil {
				panic(err)
			}
			fmt.Println("Length is :", length)
			defer file.Close()
			readFile("./mylcogofile.txt")
		}
	}

	func readFile(filename string) {
		databytes, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)

		}
		fmt.Println("Text data inside the file is \n", databytes)
	}

	func checkNillErr(err error) {
		if err != nil {
			panic(err)
		}
	}
*/
/*
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(W http.ResponseWriter, r *http.Request) {
		fmt.Fprint(W, "web server are easy in go ")

	})

	http.HandleFunc("/home", func(W http.ResponseWriter, r *http.Request) {
		http.ServeFile(W, r, "./home.html")

	})
	http.ListenAndServe(":3000", nil)
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you do ")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}
*/
package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}
func Handler(W http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")

	io.Copy(W, f)
}
