package main

import (
	"fmt"
)

func main() {
	// var numbers [4]int

	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"a", "b", "c"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// var fruits = [3]string{"apel", "pisang", "mangga"}
	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fmt.Printf("%#v\n", fruits)

	// var fruits = [3]string{"apple", "banana", "mango"}

	// for i, v := range fruits {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, v)
	// }

	// fmt.Println(strings.Repeat("#", 25))

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	// }

	// balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	// for _, arr := range balances {
	// 	for _, value := range arr {
	// 		fmt.Printf("%d", value)
	// 	}
	// 	fmt.Println()
	// }

	// var fruits = []string{"apple", "banana", "mango"}

	// _ = fruits

	// fmt.Printf("%#v", fruits)

	// var fruits = make([]string, 3)

	// _ = fruits

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fmt.Printf("%#v", fruits)

	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pinaple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pinaple", "starfruit"}

	// nn := copy(fruits1, fruits2)

	// fmt.Println("fruits1:", fruits1)
	// fmt.Println("fruits2:", fruits2)
	// fmt.Println("copied element:", nn)

	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pinaple", "starfruit"}
	// fmt.Println("fruits1:", fruits1)

	// var fruits2 = fruits1[1:4]
	// fmt.Println("fruits2:", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Println("fruits3:", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Println("fruits4:", fruits4)

	// var fruits5 = fruits1[:]
	// fmt.Println("fruits5:", fruits5)

	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pinaple", "starfruit"}

	// fruits1 = append(fruits1[:3], "orange")

	// fmt.Println(fruits1)

	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pinaple", "starfruit"}
	// var fruits2 = fruits1[2:4]
	// fruits2[0] = "rambutan"

	// fmt.Println("fruits1:", fruits1)
	// fmt.Println("fruits2:", fruits2)

	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pinaple", "starfruit"}

	// fmt.Println("fruits1 cap:", cap(fruits1))
	// fmt.Println("fruits1 len:", len(fruits1))

	// fmt.Println(strings.Repeat("#", 25))

	// var fruits2 = fruits1[0:3]

	// fmt.Println("fruits2 cap:", cap(fruits2))
	// fmt.Println("fruits2 len:", len(fruits2))

	// fmt.Println(strings.Repeat("#", 25))

	// var fruits3 = fruits1[1:]

	// fmt.Println("fruits2 cap:", cap(fruits3))
	// fmt.Println("fruits2 len:", len(fruits3))

	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	newCars[0] = "suzuki"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)
}
