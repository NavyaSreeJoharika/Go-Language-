package main

import (
	"fmt"
)

/*
	func myLearning() {
		fmt.Println("Hi There")
	}

func main() {

		myLearning()
	}
*/
/*
func lee() {
	fmt.Println("Hello Everyone")

}

func main() {
	lee()
	lee()
	lee()
}
*/

/*
func myFamily(fname string) {
	fmt.Println("Hi", fname, "Gompa")

}

func main() {
	myFamily("Navya")
	myFamily("Pramodh")
	myFamily("Madhuri")
	myFamily("Abhi")
}
*/
/*
func myAge(age int, name string) {
	fmt.Println("Hey", name, "you are", age, "years old")
}

func main() {
	myAge(20, "Navya")
	myAge(17, "Pramodh")
	myAge(22, "Madhuri")
	myAge(20, "Abhi")
}
*/
/*
func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println("x + y =", add(1, 7))
}
*/
/*
func myFunction(i int, j int) (result int) {
	result = i * j
	return
}
func main() {
	fmt.Println(myFunction(1, 6))
}
*/
/*
func MyFunction(a string, b string) (output string) {
	output = a + " " + b
	return
}
func main() {
	fmt.Println(MyFunction("Hi", "Hello"))
}
*/
/*
func me(x string, y string) (repe string) {
	repe = x + y
	return repe
}
func main() {
	fmt.Println(me("me", "go"))
}
*/
/*
func myy(i int, j string) (result int, txt string) {
	result = i
	txt = j + " " + "Hello"
	return
}
func main() {
	fmt.Println(myy(7, "Hlo"))
}
*/
/*
func myFunc(i int, j string) (result int, output string) {
	result = i + i
	output = j + "hi"
	return
}
func main() {
	a, b := myFunc(3, "Jo")
	fmt.Println(a, b)
}

*/
/*
func myFunc(i int, j int) (ads int, sbu int) {
	ads = i + j
	sbu = i - j
	return
}
func main() {
	total, sub := myFunc(1, 5)

	fmt.Println(total, sub)

}
*/
/*
func addMain(i int, j int) (result int) {
	result = i + j
	return
}
func main() {
	total := addMain(4, 986)
	fmt.Println("the total of a,b is ", total)
}
*/
/*
func myFunc(i int, j string) (result int, txt string) {
	result = i * 3
	txt = j + "Hello Everyone"
	return
}
func main() {
	fmt.Println(myFunc(5, "BTS here"))
}
*/ /*
func myFunc(i int, j string) (result int, txt string) {
	result = i + 1
	txt = "Hi" + " " + j
	return
}
func main() {
	a, b := myFunc(5, "Guys")
	fmt.Println(a, b)
}
*/ /*
func myFunc(i int, j string) (result int, txt string) {
	result = i + 1
	txt = "Hi" + " " + j
	return
}
func main() {
	_, b := myFunc(5, "Guys")
	fmt.Println(b)
}
*/
func myFunc(i int, j string) (result int, txt string) {
	result = i + 1
	txt = "Hi" + " " + j
	return
}
func main() {
	a, _ := myFunc(5, "Guys")
	fmt.Println(a)
}
