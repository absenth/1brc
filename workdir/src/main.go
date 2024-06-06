package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"sort"
	"strconv"
	"strings"
)

func main() {

	datamap := make(map[string][]int64)

	f, err := os.Open("data/sample.txt")
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
		fmt.Printf("%s %d \n", city, temp)

		/*
			next steps:
			 * find min, max, and mean after all lines are read
			 * divide min/max/mean by 10 to replace my mising deicmal
		*/
		datamap[city] = append(datamap[city], temp)
	}
	// NOTE: This sorts the cities into a slice and alphabatizes them
	keys := make([]string, 0, len(datamap))

	for key := range datamap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// NOTE: "keys" are an alphabetically sorted list of cities from the datamap

	// avgtemp := 0

	for index, element := range keys {
		fmt.Printf("%d %s \n", index, element)

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
