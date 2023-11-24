package endpoints

import (
	"emailn/internal/contract"
	internalErrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var req contract.NewCampaign
	render.DecodeJSON(r.Body, &req)
	id, err := h.CampaignService.Create(req)
	if err != nil {
		if errors.Is(err, internalErrors.ErrInternal) {
			return nil, http.StatusInternalServerError, err
		}
		return nil, http.StatusBadRequest, err
	}

	return map[string]string{"id": id}, http.StatusCreated, nil
}

func (h *Handler) GetCampaigns(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Repository.Get()
	if err != nil {
		if errors.Is(err, internalErrors.ErrInternal) {
			return nil, http.StatusInternalServerError, err
		}
		return nil, http.StatusBadRequest, err
	}

	return campaigns, http.StatusOK, nil
}
