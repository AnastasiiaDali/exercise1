package get_numbers_from_files

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"exercise1/helpers/data_converter"
	"exercise1/helpers/unique_numbers"
)

func GetNumbersFromFile(ArrayOfFileNamesFromCLI []string) []int {
	var tempNumbers []int
	var arrayOfInputs []string

	wdir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	for _, file := range ArrayOfFileNamesFromCLI {
		filePath := wdir + "/data/" + file

		d, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Failed to open file: %s", filePath)
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
