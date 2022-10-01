package learngohttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HomePageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello, this is Home Page")
}

func TestHomePageHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomePageHandler)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func TestGETHomePage(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HomePageHandler(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
$ go test -v -run TestGETHomePage
=== RUN   TestGETHomePage
Hello, this is Home Page
--- PASS: TestGETHomePage (0.00s)
PASS
ok      learn-go-httprouter     0.792s
*/

func TestPOSTHomePage(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()

	HomePageHandler(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
$ go test -v -run TestPOSTHomePage
=== RUN   TestPOSTHomePage
Hello, this is Home Page
--- PASS: TestPOSTHomePage (0.00s)
PASS
ok      learn-go-httprouter     0.800s
*/

// Nah loh, secara default ServeMux bisa nerima semua jenis method
