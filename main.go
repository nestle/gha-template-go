package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	whoToGreet := os.Getenv("INPUT_WHO-TO-GREET")

	fmt.Printf("Hello %s\n", whoToGreet)
	time := time.Now()

	fmt.Printf("::set-output name=time::%s\n", time)
}