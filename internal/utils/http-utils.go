package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/messages"
)

func SendResponse(w http.ResponseWriter, data *models.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)

	return json.NewEncoder(w).Encode(data)
}

func GenSuccessResponse(entity string, messageCode int, data interface{}) *models.Response {
	httpCode := messageCode
	if httpCode == http.StatusNoContent || httpCode == http.StatusAccepted || httpCode == codes.LOGIN_SUCCESS {
		httpCode = http.StatusOK
	}

	if httpCode == codes.ROLE_IN_USE {
		httpCode = http.StatusConflict
	}

	return &models.Response{
		Success: true,
		Message: messages.Success(entity, messageCode),
		Data:    data,
		Code:    httpCode,
	}
}

func GenErrorResponse(entity string, statusCode int, err error) *models.Response {
	appErr := errors.New(entity, statusCode, err)
	return &models.Response{
		Success: false,
		Message: appErr.UserMessage,
		Code:    statusCode,
	}
}

func GenAuthResponse(authCode, statusCode int) *models.Response {
	return &models.Response{
		Success: false,
		Message: messages.Error(entities.AUTHORIZATION, authCode),
		Code:    statusCode,
	}
}
