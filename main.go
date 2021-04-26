package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	polyline "github.com/twpayne/go-polyline"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var location [][]float64

	for scanner.Scan() {
		re2 := regexp.MustCompile(`Location\sarrived\s\<(.+)\>`)
		match := re2.FindStringSubmatch(scanner.Text())
		if len(match) > 1 {
			splitedmatch := strings.Split(match[1], ",")

			if lat, err0 := strconv.ParseFloat(splitedmatch[0], 64); err0 == nil {
				if long, err1 := strconv.ParseFloat(splitedmatch[1], 64); err1 == nil {
					location = append(location, []float64{lat, long})
				}
			}
		}

	}

	fmt.Println(string(polyline.EncodeCoords(location)))
}
