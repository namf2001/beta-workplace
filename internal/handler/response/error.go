package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error represents an HTTP web error tailored for JSON representation
type Error struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Desc   string `json:"desc"`
}

func (e *Error) Error() string { return e.Desc }

// HandleError evaluates an error and responds with the appropriate JSON payload
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	var resErr *Error
	if errors.As(err, &resErr) {
		c.JSON(resErr.Status, resErr)
	} else {
		// Fallback for unknown errors
		c.JSON(http.StatusInternalServerError, Error{
			Status: http.StatusInternalServerError,
			Code:   "internal_error",
			Desc:   err.Error(),
		})
	}
}
