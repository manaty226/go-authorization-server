package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code ErrorCode, state, redirect_uri string) {
	if redirect_uri == "" {
		c.JSON(
			code.HttpStatus(),
			gin.H{
				"error": code,
				"state": state,
			},
		)
	} else {
		c.Redirect(
			code.HttpStatus(),
			fmt.Sprintf("%s?error=%s&state=%s", redirect_uri, code, state),
		)
	}
}
