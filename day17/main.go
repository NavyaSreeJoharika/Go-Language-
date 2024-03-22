/*
package main

import (

	"fmt"

)

func main() {

		a := 10
		b := 30
		c := a + b

		fmt.Println(c)
		count := 0

		for {

			if count == 10 {
				break

			}
			count++
			fmt.Println(count)
		}
		fmt.Println(count)
	}
*/
//buffer
/*
package main

import (
	"bytes"
	"fmt"
)

// importing the bytes package so that buffer can be used
func main() {
	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Ram")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is", strBuffer.String())
}

*/
/*
package main

import (
	"bytes"
	"fmt"
)

//importing the bytes package so that buffer can be used

func main() {
	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")
	strBuffer.WriteString("[DONE]")
	fmt.Println("The string buffer output is", strBuffer.String())
}

*/
/*
package main

import (
	"bytes"
	"fmt"
	"os"
)

//importing the bytes and os package so that buffer can be used on the os based

func main() {
	//Creating buffer variable to hold and manage the string data
	var byteString bytes.Buffer
	byteString.Write([]byte("Hi "))
	fmt.Fprintf(&byteString, "Hello! how are you")
	byteString.WriteTo(os.Stdout)
}
*/
/*
package main

import (
	"bytes"
	"fmt"
)

//importing the bytes package so that buffer can be used

func main() {
	//Creating buffer variable to hold and manage the string data
	var strByyte bytes.Buffer
	strByyte.Grow(64)
	strByytestrByyte := strByyte.Bytes()
	strByyte.Write([]byte("It is a 64 byte"))
	fmt.Printf("%b", strByytestrByyte[:strByyte.Len()])
}
*/
/*
package main

import (
	"bytes"
	"fmt"
)

// importing the bytes package so that buffer can be used
func main() {
	//Creating buffer variable to hold and manage the string data
	var strByyte bytes.Buffer
	strByyte.Grow(64)
	strByyte.Write([]byte("navya"))
	fmt.Printf("%b", strByyte.Len())
}
*/
/*
package main

import (
	"bytes"
	"fmt"
)

// importing the bytes package so that buffer can be used
func main() {
	//Creating buffer variable to hold and manage the string data
	var strByyte = []byte("Ram, Kumar")
	strByyte = bytes.TrimPrefix(strByyte, []byte("Hello,  "))
	strByyte = bytes.TrimPrefix(strByyte, []byte("Good we will meet again,"))
	fmt.Printf("Hello%s", strByyte)
}
*/

//end  of buffer

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		go dbCall(i)
	}
	fmt.Printf("\nTotal execution time : %v", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is :", dbData[i])
}
