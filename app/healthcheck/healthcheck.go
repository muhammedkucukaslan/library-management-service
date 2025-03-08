package healthcheck

import "context"

type HealthcheckRequest struct{}

type HealthcheckResponse struct {
	Status string `json:"status"`
}

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Handle(ctx context.Context, req *HealthcheckRequest) (*HealthcheckResponse, error) {
	return &HealthcheckResponse{
		Status: "OK",
	}, nil
}
