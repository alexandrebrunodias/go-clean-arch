package common

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HandleError(ctx echo.Context, err error) {
	switch err {
	case ErrNotFound:
		ctx.JSON(http.StatusNotFound, Response{err.Error()})
	default:
		ctx.JSON(http.StatusInternalServerError, Response{ErrInternal.Error()})
	}
}