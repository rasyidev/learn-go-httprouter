package learngohttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	router := httprouter.New()
	// :id adalah parameter dinamis
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		fmt.Fprint(w, "Product "+id)
	})

	request := httptest.NewRequest(http.MethodGet, "/products/888", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Product 888", string(body))
}

/*
$ go test -v -run TestParam
=== RUN   TestParam
--- PASS: TestParam (0.00s)
PASS
ok      learn-go-httprouter     0.831s
*/
