package temphttp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"exercise1/adapters/http/auth"
	"exercise1/domain/formatter"
	"exercise1/helpers/string_to_int_converter"
	mux2 "github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type TempCalculator interface {
	Add(numbers []int) int
}

type server struct {
	tempCalculator TempCalculator
}

func NewRouter(calculatorService TempCalculator) http.Handler {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	keys := os.Getenv("AUTH_KEYS")

	authKeys := strings.Split(keys, ",")

	var authorisedUsers []string

	for _, key := range authKeys {
		authorisedUsers = append(authorisedUsers, key)
	}

	svr := server{tempCalculator: calculatorService}

	authMiddleware := auth.NewAuthMiddleware(authorisedUsers)
	mux := mux2.NewRouter()

	mux.Use(authMiddleware.AuthHandler)

	mux.HandleFunc("/add", svr.calculator).Methods(http.MethodPost)

	return mux
}

func (s server) calculator(w http.ResponseWriter, r *http.Request) {
	var t struct {
		Nums []string
	}

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

	//format sum
	formatter := formatter.New()
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
