package learngohttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestMethodNotAllowedHandler(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Methodnya salah, Bree. Cek lagi deh!")
	})
	router.POST("/submit", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Halaman submit, cuma terima Method POST aja!")
	})

	// test
	request := httptest.NewRequest(http.MethodGet, "/submit", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestMethodNotAllowedHandler
Methodnya salah, Bree. Cek lagi deh!
--- PASS: TestMethodNotAllowedHandler (0.00s)
PASS
ok      learn-go-httprouter     0.872s
*/
