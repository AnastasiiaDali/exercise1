package main

import (
	"flag"
	"os"

	"exercise1/adapters/data_collector"
	"exercise1/adapters/data_distributor"
	"exercise1/adapters/printer"
	cli_flags "exercise1/cmd"
	"exercise1/domain/calculator"
	"exercise1/domain/formatter"
)

func main() {
	ReadFlags()

	dataCollector := data_collector.New()
	dataDistributor := data_distributor.New(dataCollector)

	// get numbers from input
	numbers := dataDistributor.Distribute(cli_flags.ArrayOfFileNamesFromCLI, cli_flags.ArrayOfNumbersFromCLI)

	//pass this numbers to calculator and get the sum
	calculator := calculator.New()
	sum := calculator.Add(numbers)

	//pass sum to formatter and get desired formatted sum
	formatter := formatter.New()
	formattedSum := formatter.FormatNumbers(sum)

	//print sum
	printer.Printer(os.Stdout, formattedSum)
}

func ReadFlags() {
	//get numbers passed to cli
	flag.Var(&cli_flags.ArrayOfNumbersFromCLI, "input-numbers", "pass numbers")

	//get files passed to cli
	flag.Var(&cli_flags.ArrayOfFileNamesFromCLI, "input-file", "pass file name")

	flag.Parse()
}
