package main

import "fmt"

// Pi is number Pi https://github.com/golang/go/wiki/CodeReviewComments#doc-comments
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
