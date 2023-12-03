// Reads calibration document
// Document contains lines of text, each made of alphanumeric characters
// Calibration value is made by combining the first and last digit of each line
// Example: "uen2hnk34nh15" -> 25, "1n2n3n4n5n" -> 15, "1n2n3n4n5n6n" -> 16, "nthkkeo3hnk" -> 33
// Sum of all calibration values is returned

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataStr := string(data)
	lines := strings.Split(dataStr, "\n")
	sum := int64(0)
	// TODO: how to handle shared letters? "twone" -> "2" and "1", "eightwo" -> "8" and "2"
	re := regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)
	valuesDict := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	// loop over each line
	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		if len(matches) == 0 {
			continue
		}

		if len(matches) == 1 {
			valueStr := fmt.Sprintf("%s%s", valuesDict[matches[0]], valuesDict[matches[0]])
			log.Println(valueStr)
			value, err := strconv.ParseInt(valueStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			sum += value
			continue
		}

		valueStr := fmt.Sprintf("%s%s", valuesDict[matches[0]], valuesDict[matches[len(matches)-1]])
		log.Println(valueStr)
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		sum += value
	}

	fmt.Println(sum)
}
