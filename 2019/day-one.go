package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	total := 0.0
	file, err := os.Open("./input-day1")
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		localTotal := 0.0
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("err")
		}
		total += calcFuel(float64(val), localTotal)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%f", total))
}

func calcFuel(fuel float64, total float64) float64 {
	div := math.Floor(float64(fuel)/3.0) - 2.0
	if div <= 0.0 {
		return total
	} else {
		total += div
		return calcFuel(div, total)
	}
}

func partOne() {
	fuel := 0.0
	file, err := os.Open("./input-day1")
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("err")
		}
		div := math.Floor(float64(val) / 3.0)
		fuel += div - 2.0
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(int(fuel))
}