package learngohttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestRouterParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/carts/:cartId/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := fmt.Sprintf("CartID: %s | ProductID: %s", p.ByName("cartId"), p.ByName("id"))
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:9090/carts/1/products/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
$ go test -v -run TestRouterParameter
=== RUN   TestRouterParameter
CartID: 1 | ProductID: 2
--- PASS: TestRouterParameter (0.00s)
PASS
ok      learn-go-httprouter     0.845s
*/

func TestRouterParameterCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Image: " + p.ByName("image")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "/images/image/Rasyidev/Pro.jpg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestRouterParameterCatchAll
Image: /image/Rasyidev/Pro.jpg
--- PASS: TestRouterParameterCatchAll (0.00s)
PASS
ok      learn-go-httprouter     0.911s
*/
