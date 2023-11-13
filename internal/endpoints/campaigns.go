package endpoints

import (
	"emailn/internal/contract"
	internalErrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var req contract.NewCampaign
	render.DecodeJSON(r.Body, &req)
	id, err := h.CampaignService.Create(req)
	if err != nil {
		if errors.Is(err, internalErrors.ErrInternal) {
			render.Status(r, http.StatusInternalServerError)
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": id})
}

func (h *Handler) GetCampaigns(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, h.CampaignService.Repository.Get())
}
