package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {

	datamap := make(map[string][]int)

	f, err := os.Open("data/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var city, stemp string
		splitrecord := strings.SplitN(scanner.Text(), ";", 2)
		city, stemp = splitrecord[0], splitrecord[1]
		itemp := strings.Replace(stemp, ".", "", 1)
		temp, _ := strconv.ParseInt(itemp, 0, 32)
		datamap[city] = append(datamap[city], int(temp))
	}

	keys := make([]string, 0, len(datamap))

	for key := range datamap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, element := range keys {
		mintemp := float32(slices.Min(datamap[element])) / 10
		maxtemp := float32(slices.Max(datamap[element])) / 10
		totaltemp := sum(datamap[element])
		count := len(datamap[element])
		meantemp := float32((totaltemp / count)) / 10
		fmt.Printf("%s=%.1f/%.1f/%.1f, ", element, mintemp, meantemp, maxtemp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sum(arr []int) int {

	var sum int
	for idx, _ := range arr {
		sum += arr[idx]
	}
	return sum
}
