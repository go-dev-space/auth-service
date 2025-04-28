package interfaces

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/auth-service/pkg/logs"
)

func Test_Router(t *testing.T) {

	tests := []struct {
		name           string
		reqBody        string
		route          string
		method         string
		expectedStatus int
	}{
		{"Healthcheck route", ``, "/healthcheck", http.MethodGet, http.StatusOK},
		{"Registration route", `{"email":"user@domain.com","password":"a1b2"}`, "/auth/user/create", http.MethodPost, http.StatusOK},
	}

	logger := &logs.Logwriter{}

	accessHeader := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	registerHandler := NewRegistrationHandler(logger, &MockRegistrationUserUseCase{})
	healthcheckHandler := NewHealthcheckHandler(logger, &MockHealthcheckUseCase{})

	authRouter := NewRouter(*healthcheckHandler, *registerHandler, accessHeader)

	for _, test := range tests {

		stream := strings.NewReader(test.reqBody)
		request := httptest.NewRequest(test.method, test.route, stream)
		request.Header.Add("X-Access-Header", os.Getenv("ACCESS_HEADER"))
		// log.Println(os.Getenv("ACCESS_HEADER"))
		response := httptest.NewRecorder()

		authRouter.ServeHTTP(response, request)

		t.Run(test.name, func(t *testing.T) {
			if test.expectedStatus == http.StatusOK && response.Result().StatusCode != http.StatusOK {
				t.Errorf("[%s] expected status %d, but got %d", test.name, test.expectedStatus, response.Result().StatusCode)
			}
		})

	}

}
