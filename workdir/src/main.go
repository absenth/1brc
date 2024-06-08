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
		mintemp := slices.Min(datamap[element]) / 10
		maxtemp := slices.Max(datamap[element]) / 10
		totaltemp := sum(datamap[element])
		count := len(datamap[element])
		meantemp := (totaltemp / count) / 10
		fmt.Printf("%s = %d %d %d \n", element, mintemp, meantemp, maxtemp)

		// TODO: Iterate over the map based on the element, in this loop.
		// TODO: Create the average temp
		// TODO: Stop Doing It Wrong
	}

	// NOTE: This should properly print the output...
	//fmt.Printf("%s = %d %d %d", key, slices.Min(datamap[key]), avg, slices.Max(datamap[key]))

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
