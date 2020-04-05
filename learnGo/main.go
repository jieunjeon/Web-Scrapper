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
