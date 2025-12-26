package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// Exercise 1
	var score int = 85
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Println("Score:", score, "Grade:", grade)

	// Excercise 2
	if randInt := rand.Intn(100); randInt%2 == 0 {
		fmt.Printf("Number %d is even\n", randInt)
	} else {
		fmt.Printf("Number %d is odd\n", randInt)
	}

	// Excercise 3
	fmt.Println("Count up:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	countDown := 5
	fmt.Println("Count down:")
	for countDown > 0 {
		fmt.Println(countDown)
		countDown--
	}

	counter := 0
	for {
		counter++
		fmt.Println("tick")
		if counter > 2 {
			break
		}
	}

	fmt.Println("Odds:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

	}

	// Exercise 4: Range over collections

	fruits := []string{"apple", "banana", "cherry"}

	for i, fruit := range fruits {
		fmt.Printf("%d: %s ", i, fruit)
	}
	fmt.Println()

	for _, fruit := range fruits {
		fmt.Printf("%s ", fruit)
	}
	fmt.Println()

	switch today := time.Now().Weekday(); today {
	case time.Monday:
		fmt.Printf("Today is %s, time for work :(", today)
	case time.Friday:
		fmt.Printf("Today is %s, almost there", today)
	case time.Saturday, time.Sunday:
		fmt.Printf("Today is %s, weekend!", today)
	default:
		fmt.Printf("Today is %s, nothing special", today)
	}
	fmt.Println()

	var temp int = rand.Intn(100)
	switch {
	case temp < 32:
		fmt.Printf("Temperaure is %d F: Freezing", temp)
	case temp < 62:
		fmt.Printf("Temperaure is %d F: Cold", temp)
	case temp < 79:
		fmt.Printf("Temperaure is %d F: Comfortable", temp)
	default:
		fmt.Printf("Temperaure is %d F: Hot", temp)
	}
	fmt.Println()

	// Exercise 6: Defer Demonstration
	deferOpenFileDemo()

	// Expercise 7: switch type assertion
	switchTypeAssertion("hello")
	switchTypeAssertion(42)
	switchTypeAssertion(true)
	switchTypeAssertion(3.14)
}

func deferOpenFileDemo() error {
	fmt.Println("Function starting")
	file, err := os.Open("exercises/week1-basics/day3/day3.md")

	if err != nil {
		return err
	}

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	defer fmt.Println("Closing resource deferred")
	defer file.Close()

	fmt.Println("Function ending")
	return nil
}

func switchTypeAssertion(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("%s is a string with length %d\n", v, len(v))
	case int:
		fmt.Printf("%d is an integer\n", v)
	case bool:
		fmt.Printf("%t is a bool\n", v)
	default:
		fmt.Printf("%v is an unknown type: %T\n", v, v)
	}
}
