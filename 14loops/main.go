package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", i, day)
	// }

	for _, day := range days {
		fmt.Printf("index is i and value is %v\n", day)
	}

lco:
	fmt.Println("goto hit")

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 5 {
			break
		}

		if rougeValue == 3 {
			rougeValue++
			continue
		}

		if rougeValue == 2 {
			goto lco
		}

		fmt.Println("value is ", rougeValue)
		rougeValue++
	}
}
