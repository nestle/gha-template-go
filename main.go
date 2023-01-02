package main

import (
	"fmt"
	"os"
	"time"
)

func appendStringToFile(appendString string, toFile string) {

	file, err := os.OpenFile(toFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed opening file: %s", err)
	}

	defer file.Close()

	_, err2 := file.WriteString(appendString)
	if err2 != nil {
		fmt.Printf("Failed writing to file: %s", err2)
	}

}

func main() {
	whoToGreet := os.Getenv("INPUT_WHO-TO-GREET")

	fmt.Printf("Hello %s\n", whoToGreet)
	currentTime := time.Now()
	stringToAppend := "CURRENT_TIME=" + currentTime.String()

	appendStringToFile(stringToAppend, os.Getenv("GITHUB_OUTPUT"))
}