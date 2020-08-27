package learnGoMain

import (
	"fmt"
	"strings"
)

func varAndConstant() {
	var name string = "jieun"
	shortName := "jieun"
	const name2 string = "jieun2"

	fmt.Println(name, shortName)
	fmt.Println("Hello World!")

	/** variables and constants **/
	totalLenght, _ := lenAndUpper("jieun")
	fmt.Println(totalLenght)
}

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

func canIDrink_switchLoop(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

}

func goSpecific_pointer() {
	a := 2
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	/* 0xc000014078
	2 */
	a = 5
	fmt.Println(b)
	fmt.Println(*b)
	/* 0xc000014078
	5 */
	*b = 20
	fmt.Println(a, *b)
	/* 20 20 */
}

func go_Arrays_slice() {
	names := []string{"a", "b", "c"}
	namesWithLength := [5]string{"1", "2", "3"}

	names = append(names, "d") // append(0) returns new slice for the array
	fmt.Println(names)
	namesWithLength[3] = "333"
	fmt.Println(namesWithLength)
}

func go_map() {
	newMap := map[string]string{"name": "jieun", "age": "25"} // key and value, following the types
	fmt.Println(newMap)

	for key, value := range newMap {
		fmt.Println(key, value)
	}
}

func go_struct() {
	type person struct {
		name    string
		age     int
		favFood []string
	}
	favFood := []string{"ramen", "padthai"}
	whistle := person{name: "whistle", age: 25, favFood: favFood}
	fmt.Println(whistle)
}

func main() {
	goSpecific_pointer()
	go_Arrays_slice()
	go_map()
	go_struct()

}
