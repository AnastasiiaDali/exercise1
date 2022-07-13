package temphttp

import (
	"encoding/json"
	"io"
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
	var t struct {
		Nums []string
	}

	formatter := formatter.New()

	//get numbers from url
	num := r.URL.Query()

	if !num.Has("num") {
		r.ParseForm()
		for _, value := range r.Form {
			t.Nums = value
		}
	} else if num.Has("num") {
		for _, v := range num {
			t.Nums = v
		}
	}

	resBytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if len(resBytes) != 0 {
		err = json.Unmarshal(resBytes, &t)
		if err != nil {
			panic(err)
		}
	}

	//convert number into array of integers
	numbers := string_to_int_converter.StringToIntConverter(t.Nums)

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
