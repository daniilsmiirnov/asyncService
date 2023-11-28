// package main

// import "fmt"

//	func main() {
//		fmt.Println("Hello!")
//	}
// package main

// import "fmt"

//	func main() {
//		var a int = 5 // объявление переменной типа int
//		b := 10       // автоматическое определение типа через :=
//		var (
//			c = 15        // сокращенное объявление переменной
//			d = "example" // строковая переменная
//		)
//		fmt.Println(a, b, c, d)
//	}
package main

import "fmt"

func test() {
	x := 10

	// Условный оператор if
	if x > 5 {
		fmt.Println("x больше 5")
	} else {
		fmt.Println("x меньше или равен 5")
	}

	// Циклы: for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Циклы: for range (для итерации по коллекциям)
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Индекс: %d, Значение: %d\n", index, value)
	}
}
