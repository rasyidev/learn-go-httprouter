package learngohttprouter

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// panic handler, untuk menangkap dan mengirim error di HTTP Response
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Error: ", i)
	}

	router.GET("/error", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := errors.New("halaman ini masih dalam maintenance")
		panic(err.Error())
	})

	// test
	request := httptest.NewRequest(http.MethodGet, "/error", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

/*
=== RUN   TestPanicHandler
Error: halaman ini masih dalam maintenance
--- PASS: TestPanicHandler (0.00s)
PASS
ok      learn-go-httprouter     0.865s
*/
