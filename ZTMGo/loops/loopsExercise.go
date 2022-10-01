package main

import "fmt"

//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

func main() {
	//* Print integers 1 to 50, except:
	for i := 0; i <= 50; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
			//  - Print "Buzz" if the integer is divisible by 5
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		} else {
			fmt.Println(i)
		}
	}
}
