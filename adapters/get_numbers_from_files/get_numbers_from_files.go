package get_numbers_from_files

import (
	"fmt"
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
	fmt.Println(wdir)

	for _, file := range ArrayOfFileNamesFromCLI {
		fmt.Println(wdir + "/data/" + file)
		data, err := os.ReadFile(wdir + "/data/" + file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input := string(data)
		arrayOfInputs = append(arrayOfInputs, input)
	}

	tempNumbers = data_converter.DataConverter(strings.Join(arrayOfInputs[:], ","))

	numbers := unique_numbers.UniqueNumbers(tempNumbers)

	return numbers
}
