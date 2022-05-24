package adapters

import (
	"flag"
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

func GetDataFromCLI() []int {

	//get numbers passed to cli
	flag.Var(&ArrayOfNumbersFromCLI, "input-numbers", "pass numbers")

	//get files passed to cli
	flag.Var(&ArrayOfFileNamesFromCLI, "input-file", "pass file name")

	flag.Parse()

	numbers := DataDistributor(ArrayOfFileNamesFromCLI, ArrayOfNumbersFromCLI)

	return numbers
}
