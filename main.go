package main

import (
	"fmt"
	"time"
)

func getRange(tol float64, val float64) (float64, float64) {
	low := (float64(val) * (1 - tol))
	high := (float64(val) * (1 + tol))
	return low, high
}

func main() {
	for {
		var tol float64
		fmt.Print("Enter the tolerance (0.05, 0.1, etc): ")
		_, err := fmt.Scan(&tol)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number for tolerance.")
			continue
		}

		for {
			var val float64
			fmt.Print("Enter value: ")
			_, err := fmt.Scan(&val)
			if err != nil {
				fmt.Println("Invalid input. Please enter an integer value.")
				continue
			}

			low, high := getRange(tol, val)
			fmt.Printf("Acceptable Range: (%.3f, %.3f)\n", low, high)

			var yn string
			fmt.Print("New value with same tolerance? [y/n]: ")
			fmt.Scan(&yn)
			if yn == "n" {
				break
			}
		}

		var yn string
		fmt.Print("New tolerance? [y/n]: ")
		fmt.Scan(&yn)
		if yn == "n" {
			fmt.Println("🧚 Thank you katekiiiii 🧚")
			time.Sleep(3 * time.Second)
			return
		}
	}
}
