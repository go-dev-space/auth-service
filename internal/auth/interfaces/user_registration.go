package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/auth-service/internal/auth/application"
	"github.com/auth-service/internal/auth/interfaces/dto"
	"github.com/auth-service/pkg/logs"
)

type RegistrationHandler struct {
	logger       *logs.Logwriter
	registration application.Registrar
}

func NewRegistrationHandler(l *logs.Logwriter, r application.Registrar) *RegistrationHandler {
	return &RegistrationHandler{
		logger:       l,
		registration: r,
	}
}

func (handler *RegistrationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	payload := &dto.Payload{}

	body := http.MaxBytesReader(w, r.Body, int64(1024*1024))
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(payload); err != nil {
		handler.logger.Error.Println("unprocessable entity")
		http.Error(w, "unprocessable entity", http.StatusUnprocessableEntity)
		return
	}

	// registration use case
	data, err := handler.registration.Execute(r.Context(), payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data: []map[string]any{
				{"result": data},
			},
			Error: false,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(dto.Response{
		StatusCode: http.StatusOK,
		Message:    "user created",
		Data: []map[string]any{
			{"result": data},
		},
		Error: false,
	})

}
