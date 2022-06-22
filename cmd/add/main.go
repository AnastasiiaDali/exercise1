package main

import (
	"flag"
	"os"

	"exercise1/adapters/data_distributor"
	"exercise1/adapters/printer"
	cli_flags "exercise1/cmd"
	"exercise1/domain/calculator"
	"exercise1/domain/formatter"
)

func main() {
	//get numbers passed to cli
	flag.Var(&cli_flags.ArrayOfNumbersFromCLI, "input-numbers", "pass numbers")

	//get files passed to cli
	flag.Var(&cli_flags.ArrayOfFileNamesFromCLI, "input-file", "pass file name")

	flag.Parse()

	numbers := data_distributor.DataDistributor(cli_flags.ArrayOfFileNamesFromCLI, cli_flags.ArrayOfNumbersFromCLI)

	//pass this numbers ro calculator and get the sum
	sum := calculator.Add(numbers)

	//pass sum to formatter and get desired formatted sum
	formattedSum := formatter.FormatNumber(sum)

	//print sum
	printer.Printer(os.Stdout, formattedSum)
}
