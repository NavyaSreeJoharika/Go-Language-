package main

import "fmt"

func main() {
	/*
		var apple = 21

		switch apple {
		case 1:
			fmt.Println("Apple 1")
		case 2:
			fmt.Println("Apple 2")
		case 3:
			fmt.Println("Apple 3")
		default:
			fmt.Println("Orange")
		}
	*/
	/*
		day := "a"
		switch day {
		case "a":
			fmt.Println("bbb")
		case "b":
			fmt.Println("333")
		}
	*/
	/*
		day := 4
		switch day {
		case 1, 5, 3:
			fmt.Println("odd day")
		case 2, 4:
			fmt.Println("even day")
		case 6, 7:
			fmt.Println("weekend")
		default:
			fmt.Println("not a week day")

		}
	*/

	//For LOOP
	/*
		for i := 1; i < 5; i++ {
			fmt.Print(i)
		}
	*/
	/*
					for i := 0; i < 51; i += 5 {
						fmt.Println(i)
					}

					for i := 10; i >= 1 ; i-- {
						fmt.Println(i)
					}


				sr1 := []string{"1"}
				//aar := sr1[1:3]
				fmt.Println(len(sr1))
				fmt.Println(sr1)
				//fmt.Println(aar)

			for i := 10; i >= 0; i-- {
				fmt.Println(i)

			}

		for i := 20; i <= 30; i++ {
			fmt.Println(i)
		}

		for i := 80; i <= 100; i++ {
			if i == 90 {
				continue
			}
			fmt.Println(i)
		}
	*/
	/*

			var a int = 10

			// do loop execution
			for a <= 20 {
				if a == 15 {
					// skip the iteration
					a = a + 1
					continue

				}
				fmt.Printf("value of a: %d\n", a)
				a++
			}

		for i := 10; i < 20; i++ {
			if i == 15 {
				continue
			}
			if i == 18 {
				continue
			}
			fmt.Printf("value of i: %d\n", i)
		}
	*/
	/*
					for i := 10; i >= 1; i-- {
						if i == 5 {
							break
						}
						if i == 8 {
							continue
						}
						fmt.Println(i)
					}


				var word string = "goa"
				for j := 0; j < len(word); j++ {
					fmt.Print(string(word[j]))

				}

			var word string = "goa"
			for j := 0; j < len(word); j++ {
				fmt.Println(string(word[j]))

			}

			var wor string = "goa"
			for j := 0; j < len(wor); j++ {
				fmt.Println(j, string(wor[j]))

			}

		var car = "bang"
		for i := 0; i < len(car); i++ {
			fmt.Print(string(car[i]))
		}
	*/
	/*
			adj := [2]string{"I like", "I hate"}
			sub := [3]string{"flower", "cake", "pasta"}

			for i := 0; i < len(adj); i++ {
				for j := 0; j < (len(sub)); j++ {
					fmt.Println(adj[i], sub[j])
				}
			}
			sag := []string{"pp", "gg", "ff"}
			for i := 0; i < len(sag); i++ {
				fmt.Println(i, string(sag[i]))
			}

		fruits := [3]string{"apple", "orange", "banana"}
		for idx, val := range fruits {
			fmt.Printf("%v\t %v\n", idx, val)
		}

		fruit := []string{"123", "11234", "1112345"}
		for _, val := range fruit {
			fmt.Printf("%v\n", val)
		}
	*/
	/*
		apples := [5]string{"ki", "ku", "ke", "ka", "ko"}
		for idx, val := range apples {
			fmt.Printf("%v\t%v\n", idx, val)
		}

		oranges := [3]string{"123", "456", "789"}
		for _, val := range oranges {
			fmt.Printf("%v \n", val)
		}

		var carrot = [2]string{"hi", "hello"}
		for idx, _ := range carrot {
			fmt.Printf("%v\n", idx)
		}

		var fruits = [3]string{"apple", "orange", "banana"}

		for idx, _ := range fruits {
			fmt.Printf("%v\n", idx)
		}

	*/
	/*
		var arr1 = [4]int{1, 3, 5, 7}
		slice := arr1[0:2]
		fmt.Println(slice)
		fmt.Println(cap(slice))
	*/

	for i := 4; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 20; i += 2 {
		if i == 10 {
			continue
		}
		if i == 16 {
			continue
		}
		if i == 16 {
			break
		}
		fmt.Println(i)
	}
}
