//go:generate moq -out internal/mocks/data_collector_moq.go -pkg=mocks . DataCollector
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

type DataCollector struct {
}

func New() *DataCollector {
	return &DataCollector{}
}

func (dc *DataCollector) ExtractAndDeduplicateNumbers(ArrayOfNumbersFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfNumbers []string

	for _, str := range ArrayOfNumbersFromCLI {
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

func (dc *DataCollector) GetNumbersFromFile(ArrayOfFileNamesFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfInputs []string

	wdir := "/Users/anastasiia.dalakishvili/github/personal/exercise1"

	for _, file := range ArrayOfFileNamesFromCLI {
		filePath := wdir + "/data/" + file

		d, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Failed to open file: %s\n", filePath)
			return nil
		}

		data, err := readFile(d)
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

func readFile(reader io.Reader) (string, error) {
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
