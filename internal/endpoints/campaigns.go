package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var req contract.NewCampaign
	render.DecodeJSON(r.Body, &req)
	id, err := h.CampaignService.Create(req)
	return map[string]string{"id": id}, http.StatusCreated, err
}

func (h *Handler) GetCampaigns(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	// campaigns, err := h.CampaignService.Get()
	return nil, http.StatusOK, nil
}
