package data_collector

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"exercise1/helpers/data_converter"
	"exercise1/helpers/unique_numbers"
)

type DataExtractor struct {
}

func New() *DataExtractor {
	return &DataExtractor{}
}

func (dc *DataExtractor) ExtractAndDeduplicateNumbersFromCLI(numbersFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfNumbers []string

	for _, str := range numbersFromCLI {
		stringWithoutSpaces := strings.ReplaceAll(str, " ", "")
		arrayOfNumbers = strings.Split(stringWithoutSpaces, ",")
	}

	for _, i := range arrayOfNumbers {
		j, err := strconv.Atoi(i)
		if err != nil {
			j = 0
			fmt.Printf("not able to convert the %v that is not an integer. zero value has been assign instead\n", i)
		}
		tempNumbers = append(tempNumbers, j)
	}
	numbers := unique_numbers.UniqueNumbers(tempNumbers)
	return numbers
}

func (dc *DataExtractor) ExtractAndDeduplicateNumbersFromFiles(fileNamesFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfInputs []string

	wdir := "/Users/anastasiia.dalakishvili/github/personal/exercise1"

	for _, file := range fileNamesFromCLI {
		filePath := wdir + "/data/" + file

		d, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Failed to open file: %s\n", filePath)
			return nil
		}

		data, err := ReadFile(d)
		if err != nil {
			fmt.Printf("Failed to read file: %s", file)
			return nil
		}
		arrayOfInputs = append(arrayOfInputs, data)
		d.Close()
	}

	tempNumbers = data_converter.DataConverter(strings.Join(arrayOfInputs[:], ","))

	numbers := unique_numbers.UniqueNumbers(tempNumbers)

	return numbers

}

func ReadFile(reader io.Reader) (string, error) {
	var lines []string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return strings.Join(lines, "\n"), nil
}
