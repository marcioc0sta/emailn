package endpoints

import (
	internalErrors "emailn/internal/internal-errors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrorHandler_internal_error(t *testing.T) {
	// Arrange
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalErrors.ErrInternal
	}
	handlerFn := ErrorHandler(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	// Act
	handlerFn.ServeHTTP(res, req)
	// Assert
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Contains(t, res.Body.String(), internalErrors.ErrInternal.Error())
}

func Test_ErrorHandler_bad_request(t *testing.T) {
	// Arrange
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("bad request")
	}
	handlerFn := ErrorHandler(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	// Act
	handlerFn.ServeHTTP(res, req)
	// Assert
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Contains(t, res.Body.String(), "bad request")
}

func Test_ErrorHandler_no_error(t *testing.T) {
	// Arrange
	type bodyForTest struct {
		Name string `json:"name"`
	}
	expectedObject := bodyForTest{Name: "test"}
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return expectedObject, http.StatusOK, nil
	}
	handlerFn := ErrorHandler(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	// Act
	handlerFn.ServeHTTP(res, req)
	returnedObject := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &returnedObject)
	// Assert
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, returnedObject, returnedObject)
}
