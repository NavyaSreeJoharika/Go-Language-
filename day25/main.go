/*
package main

import "fmt"

func main() {

		//declaration of a string with double quotes
		value := "Irasshaimase"
		var value2 string
		value2 = "konnichiwa"

		fmt.Println("String1:", value)
		fmt.Println(value2)

		value3 := "Mise e \n yokoso"
		value4 := `Mise e \n yokoso`

		fmt.Println(value3)
		fmt.Println(value4)

		for index, s := range "Hello" {
			fmt.Printf("index of %c is %d \n", s, index)
		}

		str := "Welcome to GeeksforGeeks"

		// Accessing the bytes of the given string
		for c := 0; c < len(str); c++ {
			fmt.Printf("\nCharacter = %v Bytes = %x", str, str)
		}
	}
*/
/*
package main

import (
	"fmt"
)

/*func updateName(x string) {
	x = "wedge"

}

func main() {
	var name string
	updateName(name)
	fmt.Println("Updated name:", name)
	//name := "tifa"

	updateName(name)

	fmt.Println(name)
}
*/
/*
func updateName(x *string) {
	*x = "wedge"
}

func main() {
	var name string
	updateName(&name)
	fmt.Println("Updated name:", name)
}
*/
/*
func updateName(x string) {
	x = ""
}

func main() {
	var name string
	name = "tina"
	updateName(name)

	fmt.Println("name:", name)
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Prince   int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("welcome")
	EncodeJson()
}

func EncodeJson() {
	locCourses := []course{
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "abc123", []string{"web-dev", "js"}},
		{"ReactsJs Bootcamp", 199, "LearnCodeOnLine.in", "abca1a2a3", []string{"full-stak", "js"}},
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "hit123", nil},
	}
	finalJson, err := json.Marshal(locCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Prince   int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("welcome")
	EncodeJson()
}

func EncodeJson() {
	locCourses := []course{
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "abc123", []string{"web-dev", "js"}},
		{"ReactsJs Bootcamp", 199, "LearnCodeOnLine.in", "abca1a2a3", []string{"full-stak", "js"}},
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "hit123", nil},
	}
	finalJson, err := json.MarshalIndent(locCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}
*/
/*

package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Prince   int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome")
	EncodeJson()
}

func EncodeJson() {
	locCourses := []course{
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "abc123", []string{"web-dev", "js"}},
		{"ReactsJs Bootcamp", 199, "LearnCodeOnLine.in", "abca1a2a3", []string{"full-stak", "js"}},
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "hit123", nil},
	}
	finalJson, err := json.MarshalIndent(locCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}
*/
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Prince   int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome")
	DecodeJson()
}

func EncodeJson() {
	locCourses := []course{
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "abc123", []string{"web-dev", "js"}},
		{"ReactsJs Bootcamp", 199, "LearnCodeOnLine.in", "abca1a2a3", []string{"full-stak", "js"}},
		{"ReactsJs Bootcamp", 299, "LearnCodeOnLine.in", "hit123", nil},
	}
	finalJson, err := json.MarshalIndent(locCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataFormWeb := []byte(`
        {
                "coursename": "ReactsJs Bootcamp",
                "Prince": 299,
                "website": "LearnCodeOnLine.in",
                "tags": [ "web-dev","js"]
        }
	`)

	var locCourses course

	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFormWeb, &locCourses)
		fmt.Printf("%#v\n", locCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	// some cases where we just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFormWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is:%T\n", k, v, v)
	}
}
