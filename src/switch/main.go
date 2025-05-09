package main

import "fmt"

func main() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool.")
		case int:
			fmt.Println("I am int.")
		default:
			fmt.Printf("Don`t know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(255)
	whatAmI("")
	whatAmI(struct{}{})
}
