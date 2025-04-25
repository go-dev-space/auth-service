package interfaces

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/auth-service/pkg/logs"
)

type MockHealthcheckUseCase struct{}

func (m *MockHealthcheckUseCase) Execute() (string, error) {
	return "pong", nil
}

func Test_HealthcheckHandler(t *testing.T) {

	tests := []struct {
		name            string
		reqBody         string
		method          string
		expectedStatus  int
		expectedMessage string
	}{
		{"Valid request", ``, http.MethodGet, http.StatusOK, "pong"},
	}

	logger := logs.New()

	for _, test := range tests {

		stream := strings.NewReader(test.reqBody)
		request := httptest.NewRequest(test.method, "/v1/healthcheck", stream)
		response := httptest.NewRecorder()

		handler := NewHealthcheckHandler(logger, &MockHealthcheckUseCase{})
		handler.Handle(response, request)

		t.Run(test.name, func(t *testing.T) {
			if test.expectedStatus == http.StatusOK && response.Result().StatusCode != http.StatusOK {
				t.Errorf("[%s] expected status %d, but got %d", test.name, test.expectedStatus, response.Result().StatusCode)
			}
		})

	}

}
