package request

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ParseJSONRequest[T comparable](c *gin.Context, requestModel T) (T, error) {
	vd := validator.New()

	var req T
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("error: %v", err)
		return requestModel, err
	}
	if err := vd.Struct(req); err != nil {
		return requestModel, err
	}
	return req, nil
}
