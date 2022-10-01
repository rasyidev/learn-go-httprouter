package learngohttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHttpRouter(t *testing.T) {
	router := httprouter.New()
	router.HandlerFunc("GET", "/", HomePageHandler)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func GetHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from HttpRouter, Method: GET")
}

func TestGETHttpRouter(t *testing.T) {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", HomePageHandler)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
$ go test -v -run TestGETHttpRouter
=== RUN   TestGETHttpRouter
Hello, this is Home Page
--- PASS: TestGETHttpRouter (0.00s)
PASS
ok      learn-go-httprouter     0.805s
*/

func TestPOSTHttpRouter(t *testing.T) {
	router := httprouter.New()
	request := httptest.NewRequest("POST", "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
$ go test -v -run TestPOSTHttpRouter
=== RUN   TestPOSTHttpRouter
404 page not found

--- PASS: TestPOSTHttpRouter (0.00s)
PASS
ok      learn-go-httprouter     0.812s
*/
