/*
package main

import "fmt"

func main() {
	a := 12
	b := &a
	c := *b
	d := *&c

	fmt.Println(b)
	fmt.Println(c)

	*b = 87434
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	a = 78899
	fmt.Println(a)

	b = new(int) //only to pointer
	fmt.Println(b, *b)
}
*/

/* output-
PS C:\golang\day32> go run main.go
0xc000016088
12
87434
12
12
78899
0xc0000160f8 0
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
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}

*/
/*
PS C:\golang\day32> go run main.go
What would you like me to scream?
go is really cool
GO IS REALLY COOL!
*/

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":4000", nil)
}
func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}

/*
output ====
PS C:\golang\day32> go run main.go
exit status 0xc000013a
http://localhost:4000/

Coffee
Espresso
Cappuccino

Lemon Tea
Hot Tea
Chai Latte
Chai Tea


Hot Chocolate
Macha
*/
