/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var person map[string]interface{}

	jsonData := `{
		"name":"Jake",
		"country":"US",
		"state":"Connecticut"
	}`
	err := json.Unmarshal([]byte(jsonData), &person)

	if err != nil {
		fmt.Println("Error while decoding the data", err.Error())
	}
	fmt.Println("The name of the person is", person["name"], "and he lives in", person["country"])
}

*/
/*

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var persons []map[string]interface{}

	jsonData := `[{
		"name":"Ashok",
		"country":"US",
		"state":"Boston"
	},
	{
		"name":"Jacob",
		"country":"Us",
		"state":"WashingtonDC"
	},
	{
		"name":"James",
		"country":"US",
		"state":"New York"
	},
	{
		"name":"Sri",
		"country":"Japan",
		"state":"Tokyo"
	}
	]`

	err := json.Unmarshal([]byte(jsonData), &persons)

	if err != nil {
		fmt.Println("Error while decoding the data", err.Error())

	}
	for index, person := range persons {
		fmt.Println("Person:", index+1, "Name:", person["name"], "Country:", person["country"], "State:", person["state"])
	}
}
*/
/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {
	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "cosco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}

	}
}
func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal an chicken at %s", <-chickenChannel)
}
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {

	var tofuChannel = make(chan string)
	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "cosco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}
func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}

	}
}
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent:Found a deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent :Found a deal on tofu at %v", website)

	}
}
