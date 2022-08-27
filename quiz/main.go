package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func parseData(data [][]string) []problem {
	ret := make([]problem, len(data))
	for i, v := range data {
		ret[i] = problem{
			q: v[0],
			a: v[1],
		}
	}
	return ret

}

func startTest(problems []problem, timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	var count int
	for _, v := range problems {
		fmt.Printf("what is %v sir ?", v.q)
		inputCh := make(chan string)
		go func() {
			var input string
			_, err := fmt.Scanln(&input)
			if err != nil {
				log.Fatal(err)
			}
			inputCh <- input
		}()

		select {
		case <-timer.C:
			fmt.Printf("you got %d of %d", count, len(problems))
			return
		case input := <-inputCh:
			if input == v.a {
				count++
			}

		}
	}
	fmt.Printf("you got %d of %d", count, len(problems))
}

func main() {

	fileName := flag.String("filename", "problems.csv", "Provide problem file")
	timeLimit := flag.Int("limit", 20, "provide time limti for quiz")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	problems := parseData(data)
	startTest(problems, *timeLimit)

}
