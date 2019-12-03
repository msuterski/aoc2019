package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const dataLink = "https://adventofcode.com/2019/day/1/input"

func main() {

	data, err := getData(os.Getenv("AOC_SESSION"))
	if err != nil {
		log.Fatal(err)
	}

	var sum int

	for _, d := range data {
		sum += calculateFuel(d)
	}

	fmt.Println("Sum:  ", sum)
}

func calculateFuel(mass int) int {

	fuel := mass/3 - 2

	if fuel <= 0 {
		return 0
	}

	fuel += calculateFuel(fuel)

	return fuel
}

func getData(sessionToken string) ([]int, error) {
	req, err := http.NewRequest("GET", dataLink, nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionToken,
	})
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Error pulling data")
		return nil, err
	}

	defer resp.Body.Close()

	csvData := csv.NewReader(resp.Body)

	data, err := csvData.ReadAll()
	if err != nil {
		log.Println("CSV parsing error")
		return nil, err
	}

	var intData []int

	for _, d := range data {
		i, err := strconv.Atoi(d[0])
		if err != nil {
			log.Println("Error converting string to int: ", d[0])
			return nil, err
		}

		intData = append(intData, i)
	}

	return intData, nil
}
