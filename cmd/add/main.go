package main

import (
	"flag"
	"os"

	"exercise1/adapters/data_distributor"
	"exercise1/adapters/printer"
	"exercise1/domain"
)

//files
type arrayOfFlags []string

func (f *arrayOfFlags) String() string {
	return ""
}

func (f *arrayOfFlags) Set(flag string) error {
	*f = append(*f, flag)
	return nil
}

var ArrayOfFileNamesFromCLI arrayOfFlags

//numbers
type arrayOfNumbers []string

func (n *arrayOfNumbers) String() string {
	return ""
}

func (n *arrayOfNumbers) Set(flag string) error {
	*n = append(*n, flag)
	return nil
}

var ArrayOfNumbersFromCLI arrayOfNumbers

func main() {
	//extract the numbers pass by user either from the file or just an array of integers

	//get numbers passed to cli
	flag.Var(&ArrayOfNumbersFromCLI, "input-numbers", "pass numbers")

	//get files passed to cli
	flag.Var(&ArrayOfFileNamesFromCLI, "input-file", "pass file name")

	flag.Parse()

	numbers := data_distributor.DataDistributor(ArrayOfFileNamesFromCLI, ArrayOfNumbersFromCLI)

	//pass this numbers ro calculator and get the sum
	sum := domain.Add(numbers)

	//pass sum to formatter and get desired formatted sum
	formattedSum := domain.FormatNumber(sum)

	//print sum
	printer.Printer(os.Stdout, formattedSum)
}
