package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // run after the function returned.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd_loop(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
		/*
			0 4
			1 4
			2 3
			3 5
		*/
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func main() {

	/** variables and constants **/
	var name string = "jieun"
	shortName := "jieun"
	const name2 string = "jieun2"

	fmt.Println(name, shortName)
	fmt.Println("Hello World!")

	/** variables and constants **/
	totalLenght, _ := lenAndUpper("jieun")
	fmt.Println(totalLenght)
}
