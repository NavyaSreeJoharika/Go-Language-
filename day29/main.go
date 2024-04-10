/*
package main

import (

	"net/http"

)

	func main() {
		http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		})
		http.ListenAndServe(":8080", nil)
	}

//   http://localhost:8080/
//output === 404 page not found
//   http://localhost:8080/hello-world
//output === hello world

*/

/*
package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Educative Home!")

}

func handleReq() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8200", nil))
}

func main() {
	handleReq()
}

//  http://localhost:8200/
// Output === Welcome to Educative Home!

*/
/*

package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Educative Home!")
}

func return_contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Email:support@educative.io")
}

func handleReq() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", return_contact)

	log.Fatal(http.ListenAndServe(":8200", nil))
}

func main() {
	handleReq()
}

//  http://localhost:8200/contact
// output ==== Email:support@educative.io

*/

package main

import (
	"fmt"
	"net/http"
)

const url = "https://api.weatherapi.com/"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type:%T\n", response)

	response.Body.Close()
}

// output ==== Web Request
//			   Response is of type:*http.Response

/*
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.weatherapi.com/"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type:%T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	constant := string(databytes)
	fmt.Println(constant)

}
*/
//output
/*  Web Request
Response is of type:*http.Response
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title></title>
</head>
<body>
    All OK
</body>
</html>

*/
