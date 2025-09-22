package main

import "fmt"

func greet() {
	// String variable declaration and assignment

	var name string = "Aravinda"
	var nametwo = "Kumar"
	var namethree string
	fmt.Println("Greetings from Go!", name, nametwo, namethree)

	namethree = "Welcome!"
	fmt.Println("Greetings from Go!", name, nametwo, namethree)

	namefour := "to the Go programming language." // We can not use := outside a function
	fmt.Println("Greetings from Go!", name, nametwo, namethree, namefour)

	// Integer variable declaration and assignment
	var age int = 25
	var ageTwo = 30
	var ageThree int
	fmt.Println("Age:", age, ageTwo, ageThree)

	ageThree = 35
	fmt.Println("Age:", age, ageTwo, ageThree)

	ageFour := 40
	fmt.Println("Age:", age, ageTwo, ageThree, ageFour)

	// Float variable declaration and assignment
	var height float32 = 5.9
	var heightTwo = 6.1
	var heightThree float32
	fmt.Println("Height:", height, heightTwo, heightThree)

	heightThree = 5.11
	fmt.Println("Height:", height, heightTwo, heightThree)

	heightFour := 6.0
	fmt.Println("Height:", height, heightTwo, heightThree, heightFour)

	// Boolean variable declaration and assignment
	var isStudent bool = true
	var isEmployed = false
	var isMarried bool
	fmt.Println("Is Student:", isStudent, isEmployed, isMarried)

	isMarried = true
	fmt.Println("Is Student:", isStudent, isEmployed, isMarried)

	isSingle := false
	fmt.Println("Is Student:", isStudent, isEmployed, isMarried, isSingle)

	// bits and memory size
	var smallNumber int8 = 127                      // 8 bits
	var mediumNumber int16 = 32767                  // 16 bits
	var largeNumber int32 = 2147483647              // 32 bits
	var veryLargeNumber int64 = 9223372036854775807 // 64 bits

	fmt.Println("Small Number:", smallNumber)
	fmt.Println("Medium Number:", mediumNumber)
	fmt.Println("Large Number:", largeNumber)
	fmt.Println("Very Large Number:", veryLargeNumber)

	fmt.Printf("Memory Size of smallNumber: %d bytes\n", 8/8)
	fmt.Printf("Memory Size of mediumNumber: %d bytes\n", 16/8)
	fmt.Printf("Memory Size of largeNumber: %d bytes\n", 32/8)
	fmt.Printf("Memory Size of veryLargeNumber: %d bytes\n", 64/8)

	var unsignedNumber uint = 42
	fmt.Println("Unsigned Number:", unsignedNumber)
}

func lecture4() {
	// Print
	fmt.Print("Hello, World! ")

	// Print with newline
	fmt.Println("Hello, World!") // Print with newline

	// Formatted Print
	name := "Aravinda"
	age := 25
	height := 5.9
	isStudent := true

	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Is Student: %t\n", name, age, height, isStudent)

	// Save the formatted string to a variable
	formattedString := fmt.Sprintf("Name: %s, Age: %d, Height: %.1f, Is Student: %t", name, age, height, isStudent)
	fmt.Println(formattedString)
}

func main() {
	lecture4()
}
