package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status int         `json:"status_code"`
	Data   interface{} `json:"data"`
}

// ErrResponse represent error response
type ErrResponse struct {
	Err    string `json:"error"`
	Status int   `json:"status_code"`
}

// Error implement custom error types
func (e ErrResponse) Error() string {
	return fmt.Sprintf("error is %v, and http status code is %d", e.Err, e.Status)
}

func ErrRespondJSON(w *gin.Context, status int, err error) {
	var errRes ErrResponse

	errRes.Status = status
	errRes.Err = err.Error()

	w.JSON(errRes.Status, errRes)
}

func SuccessRespondJSON(w *gin.Context, status int, data interface{}) {
	var successRes SuccessResponse
	successRes.Status = status
	successRes.Data = data

	w.JSON(successRes.Status, successRes)
}
