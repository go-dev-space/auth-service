package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/auth-service/internal/auth/application"
	"github.com/auth-service/internal/auth/interfaces/dto"
	"github.com/auth-service/pkg/logs"
)

type HealthcheckHandler struct {
	logger      *logs.Logwriter
	healthcheck application.Healthchecker
}

func NewHealthcheckHandler(l *logs.Logwriter, h application.Healthchecker) *HealthcheckHandler {
	return &HealthcheckHandler{
		logger:      l,
		healthcheck: h,
	}
}

// Healthcheck Handler
//
//	@Summary		Healthcheck for api (Kubernetes)
//	@Description	check the api health
//	@Tags			healthcheck
//	@Accept			json
//	@Produce		json
//	@Param			X-Access-Header	header		string	true	"Access Header for health route"
//	@Success		200				{object}	dto.Response
//	@Failure		400				{object}	dto.Response
//	@Router			/auth/healthcheck [get]
func (handler *HealthcheckHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// set header content type
	w.Header().Set("Content-Type", "application/json")

	// healthcheck use case
	result, err := handler.healthcheck.Execute()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
			Error:      true,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(dto.Response{
		StatusCode: http.StatusOK,
		Message:    result,
		Data:       nil,
		Error:      false,
	})
}
