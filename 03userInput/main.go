package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello, welcome to user user Input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating of our pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("type of our guy %T", input)
}
