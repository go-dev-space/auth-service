package interfaces

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/auth-service/internal/auth/interfaces/dto"
	"github.com/auth-service/pkg/logs"
)

type MockRegistrationUserUseCase struct{}

func (m *MockRegistrationUserUseCase) Execute(ctx context.Context, p *dto.Payload) (map[string]string, error) {
	return map[string]string{"result": "ok"}, nil
}

func Test_RegistrationHandler(t *testing.T) {

	tests := []struct {
		name            string
		reqBody         string
		method          string
		expectedStatus  int
		expectedMessage string
	}{
		{"Valid request", `{"email":"user@domain.com","password":"a1b2"}`, http.MethodPost, http.StatusOK, "user created"},
	}

	logger := &logs.Logwriter{}
	for _, test := range tests {

		stream := strings.NewReader(test.reqBody)
		request := httptest.NewRequest(test.method, "/v1/auth/user/create", stream)
		response := httptest.NewRecorder()

		handler := NewRegistrationHandler(logger, &MockRegistrationUserUseCase{})
		handler.Handle(response, request)

		t.Run(test.name, func(t *testing.T) {
			if test.expectedStatus == http.StatusOK && response.Result().StatusCode != http.StatusOK {
				t.Errorf("[%s] expected status %d, but got %d", test.name, test.expectedStatus, response.Result().StatusCode)
			}
		})

	}

}
