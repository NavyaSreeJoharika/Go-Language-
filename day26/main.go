/*
package main

import (

	"encoding/json"
	"fmt"

)

	type Human struct {
		Name    string
		Age     int
		Address string
	}

	func main() {
		human1 := Human{"Aklia", 24, "Ap"}
		human_enc, err := json.Marshal(human1)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(human_enc))

		human2 := []Human{
			{Name: "Rose", Age: 13, Address: "kolkata"},
			{Name: "Lisa", Age: 26, Address: "New Delhi"},
			{Name: "Jiso", Age: 28, Address: "Mumbai"},
			{Name: "Jenni", Age: 29, Address: "Chenni"},
		}

		human2_enc, err := json.Marshal(human2)

		if err != nil {
			fmt.Println(err)

		}
		fmt.Println()
		fmt.Println(string(human2_enc))
	}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var human1 Human

	Data := []byte(`{
		"Name":"Devi",
		"Age":25,
		"Address":"Kakinada"
	}`)

	err := json.Unmarshal(Data, &human1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Struct is:", human1)
	fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address)

	var human2 []Human
	Data2 := []byte(`[
		{"Name":"Gowtham","Age":25,"Address":"Banglore"},
		{"Name":"Datta","Age":24,"Address":"Kakinada"},
		{"Name":"Srinivas","Age":23,"Address":"Japan"}
	]`)

	err2 := json.Unmarshal(Data2, &human2)

	if err2 != nil {
		fmt.Println(err2)
	}

	for i := range human2 {
		fmt.Println(human2[i])
	}
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

type Country struct {
	Name      string
	Capital   string
	Continent string
}

func main() {
	var count1 Country
	Data := []byte(`{
		"Name":"India",
		"Capital":"New Delhi",
		"Continent":"Asia"
	}`)

	err := json.Unmarshal(Data, &count1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Struct is:", count1)
	fmt.Printf("%s's captial is %s and it is in %s.\n", count1.Name, count1.Capital, count1.Continent)
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

type Country struct {
	Name    string
	Capital string
}

func main() {
	var country []Country
	Data := []byte(`[

		{"Name": "Japan", "Capital": "Tokyo"},
		{"Name": "Germany", "Capital": "Berlin"},
		{"Name": "India", "Capital": "New Delhi"},
		{"Name": "Greece", "Capital": "Athens"},
		{"Name": "Israel", "Capital": "Jerusalem"},
		{"Name": "South Korea", "Capital": "Seoul"},
		{"Name": "Thailand", "Capital": "Bangkok(Krung Thep Mahanakhon Amon Rattanakosin Mahinthara Ayuthaya Mahadilok Phop Noppharat Ratchathani Burirom Udomratchaniwet Mahasathan Amon Piman Awatan Sathit Sakkathattiya Witsanukam Prasit)"},
		{"Name": "Philippines", "Capital": "Manila"},
		{"Name": "Vietnam", "Capital": " Hanoi"},
		{"Name": "Malaysia ", "Capital": "Kuala Lumpur"},
		{"Name": "Singapore", "Capital": "Singapore"},
		{"Name": "Taiwan", "Capital": "Taipei City"},
		{"Name": "China", "Capital": "Beijing"}
]`)

	err := json.Unmarshal(Data, &country)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(country)

	for i := range country {
		fmt.Println(country[i].Name + "-" + country[i].Capital)
	}
	human2 := []Country{
		{Name: "India", Capital: "New Delhi"},
		{Name: "Japan", Capital: "Tokyo"},
		{Name: "Germany", Capital: "Berlin"},
		{Name: "Greece", Capital: "Athens"},
		{Name: "Israel", Capital: "Jerusalem"},
		{Name: "South Korea", Capital: "Seoul"},
	}

	human2_enc, err := json.Marshal(human2)

	if err != nil {
		fmt.Println(err)

	}
	fmt.Println()
	fmt.Println(string(human2_enc))
}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Address string
}

func main() {

	human1 := Human{"Ankit", 23, "New Delhi"}

	human_enc, err := json.Marshal(human1)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(string(human_enc))

	human2 := []Human{
		{Name: "Rahul", Age: 23, Address: "New Delhi"},
		{Name: "Priyanshi", Age: 20, Address: "Pune"},
		{Name: "Shivam", Age: 24, Address: "Bangalore"},
	}

	human2_enc, err := json.Marshal(human2)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println(string(human2_enc))

	for i := range human2 {

		fmt.Println(human2[i])
	}
}
