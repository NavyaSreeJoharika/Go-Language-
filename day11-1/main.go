/*package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := map[string]string{
		"foo": "baba",
	}
	data, _ := json.Marshal(x)
	fmt.Println(string(data))
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type person struct {
		Name        string `json:"name"`
		age         int
		Description string `json:"descr,omitempty"`
		Secret      string `json:"secret"`
	}

	x := person{
		Name:   "Bob",
		age:    23,
		Secret: "Shhhhhu",
	}
	fmt.Println(x)

	data, _ := json.Marshal(x)
	fmt.Println(string(data))
}
*/
/*
package main

import (
	"encoding/json"
)

func main() {
	data := []byte(`{"foo":"baba"}`)

	var x interface{}
	_ = json.Unmarshal(data, &x)
	spew.Dump(x)
}
*/
package main

import (
	"fmt"
)

func main() {
	hitesh := User{"Hitesh", "hi@go.com", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %v\n ", hitesh)
	fmt.Printf("Name is %v and email is %v .\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMain()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}
func (u User) NewMain() {
	u.Email = "test@go.com"
	fmt.Println("Email of the user is :", u.Email)
}
