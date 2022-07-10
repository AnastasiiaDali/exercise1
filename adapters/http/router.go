package temphttp

import (
	"encoding/json"
	"net/http"

	"exercise1/domain/formatter"
	"exercise1/helpers/string_to_int_converter"
)

type TempCalculator interface {
	Add(numbers []int) int
}

type server struct {
	tempCalculator TempCalculator
}

func NewRouter(calculatorService TempCalculator) http.Handler {
	svr := server{tempCalculator: calculatorService}

	mux := http.NewServeMux()
	mux.HandleFunc("/add", svr.calculator)

	return mux
}

func (s server) calculator(w http.ResponseWriter, r *http.Request) {
	var numbersStr []string

	formatter := formatter.New()

	//get numbers from url
	num := r.URL.Query()
	for _, v := range num {
		numbersStr = v
	}

	//convert number into array of integers
	numbers := string_to_int_converter.StringToIntConverter(numbersStr)

	sum := s.tempCalculator.Add(numbers)
	formatterSum := formatter.FormatNumbers(sum)

	res, err := json.Marshal(formatterSum)

	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

//func MathServer() {
//	r := mux.NewRouter()
//	r.HandleFunc("/", HealthCheckHandler).Methods(http.MethodGet)
//	r.HandleFunc("/add", MathHandler).Methods(http.MethodPost)
//
//	fmt.Printf("Starting server on port 8081...\n")
//	http.ListenAndServe(":8081", r)
//}
