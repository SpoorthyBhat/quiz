package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var numCorrect int
var numTotal int

func timeUp() {
	fmt.Printf("Time up: You got %d correct out of %d\n", numCorrect, numTotal)
	os.Exit(0)
}

func main() {

	timeLimit := flag.Int("timeLimit", 10, "an integer")
	flag.Parse()

	//Open file
	problems, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	lines, err := csv.NewReader(problems).ReadAll()

	//User input
	var input int
	numTotal = len(lines)
	numCorrect = 0

	//Start timer
	timer := time.AfterFunc(time.Duration(*timeLimit)*time.Second, timeUp)

	for _, record := range lines {

		fmt.Printf("%s = ", record[0])
		fmt.Scanln(&input)
		answer, err := strconv.Atoi(record[1])

		if err != nil {
			log.Fatalln("Error parsing input file", err)
		}
		if answer == input {
			numCorrect++
		}

	}
	timer.Stop()
	fmt.Printf("You got %d correct out of %d\n", numCorrect, numTotal)
}
