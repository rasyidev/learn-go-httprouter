package learngohttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (logMiddleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Ada request baru nih")
	logMiddleware.Handler.ServeHTTP(writer, request)
}

func TestLogMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	middleware := LogMiddleware{
		Handler: router,
	}

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestLogMiddleware
Ada request baru nih
Middleware
--- PASS: TestLogMiddleware (0.00s)
PASS
ok      learn-go-httprouter     0.636s
*/
