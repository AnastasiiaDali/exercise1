package data_distributor

type DataCollector interface {
	ExtractAndDeduplicateNumbersFromCLI([]string) []int
	ExtractAndDeduplicateNumbersFromFiles([]string) []int
}

type DataDistributor struct {
	dataCollector DataCollector
}

func New(dataCollector DataCollector) *DataDistributor {
	return &DataDistributor{
		dataCollector: dataCollector,
	}
}

func (dd *DataDistributor) Distribute(fileNamesFromCLI []string, numbersFromCLI []string) []int {
	var numbers []int

	if len(numbersFromCLI) != 0 && len(fileNamesFromCLI) != 0 {
		return nil
	} else if len(fileNamesFromCLI) != 0 {
		numbers = dd.dataCollector.ExtractAndDeduplicateNumbersFromFiles(fileNamesFromCLI)
	} else if len(numbersFromCLI) != 0 {
		numbers = dd.dataCollector.ExtractAndDeduplicateNumbersFromCLI(numbersFromCLI)
	}

	return numbers
}
