package learngohttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Halaman yang kamu cari tidak ada. Sepertinya kamu tersesat :(")
	})

	// test
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestNotFoundHandler
Halaman yang kamu cari tidak ada. Sepertinya kamu tersesat :(
--- PASS: TestNotFoundHandler (0.00s)
PASS
ok      learn-go-httprouter     0.863s
*/
