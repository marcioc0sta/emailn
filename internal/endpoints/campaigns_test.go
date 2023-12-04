package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(campaign contract.NewCampaign) (string, error) {
	args := r.Called(campaign)
	return args.String(0), args.Error(1)
}

func (r *serviceMock) Get() ([]contract.CampaignResponse, error) {
	args := r.Called()
	return nil, args.Error(1)
}

func (r *serviceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := r.Called(id)
	return nil, args.Error(1)
}

func Test_CreateCampaign(t *testing.T) {
	// Arrange
	service := new(serviceMock)
	body := contract.NewCampaign{
		Name:    "test",
		Content: "test",
		Emails:  []string{"teste@teste.com"},
	}
	handler := Handler{CampaignService: service}
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name != body.Name && request.Content != body.Content {
			return false
		}
		return true
	})).Return("123456", nil)
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req := httptest.NewRequest("POST", "/campaigns", &buf)
	res := httptest.NewRecorder()
	// Act
	_, status, err := handler.CreateCampaign(res, req)
	// Assert
	assert.Equal(t, http.StatusCreated, status)
	assert.Nil(t, err)
}

func Test_CreateCampaign_error(t *testing.T) {
	// Arrange
	service := new(serviceMock)
	body := contract.NewCampaign{
		Name:    "test",
		Content: "test",
		Emails:  []string{"teste@teste.com"},
	}
	handler := Handler{CampaignService: service}
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req := httptest.NewRequest("POST", "/campaigns", &buf)
	res := httptest.NewRecorder()
	// Act
	_, _, err := handler.CreateCampaign(res, req)
	// Assert
	assert.NotNil(t, err)
}
