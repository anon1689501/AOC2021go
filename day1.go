package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	day1, err := os.Open("inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer day1.Close()

	scanner := bufio.NewScanner(day1)
	total := 0
	for scanner.Scan() {
		// myNum, err2 := scanner.Text()
		myNum, err2 := strconv.Atoi(scanner.Text())

		total = total + fuel(myNum)
		fmt.Println(total)
		if err2 != nil {
			log.Fatal(err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(total)
}

func fuel(x int) int {
	x = (x / 3) - 2
	if x < 0 {
		return 0
	}
	return x + fuel(x)

}
