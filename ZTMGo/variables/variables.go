package main

import "fmt"

func main() {
	var myName = "Chrisian"
	fmt.Println("my name is ", myName, myName)

	var name string = "Kathy"
	fmt.Println("name = ", name)

	username := "admin"
	fmt.Println("username = ", username)

	var sum int
	fmt.Println("The sum is ", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is ", part1, "other is ", other)

	part2, other := 2, 0
	fmt.Println("part 2 is ", part2, "other is ", other)

	sum = part1 + part2
	fmt.Println("sum = ", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName = ", lessonName)
	fmt.Println("lessonType = ", lessonType)

	//neccsito un cambio
	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
