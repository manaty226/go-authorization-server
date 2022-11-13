package response

import (
	"net/http"
)

type ErrorCode string

const (
	INVALID_REQUEST           ErrorCode = "invalid_request"
	UNAUTHORIZED_CLIENT       ErrorCode = "unauthorized_client"
	ACCESS_DENIED             ErrorCode = "access_denied"
	UNSUPPORTED_RESPONSE_TYPE ErrorCode = "unsupported_response_type"
	INVALID_SCOPE             ErrorCode = "invalid_scope"
	SERVER_ERROR              ErrorCode = "server_error"
	TEMPORARILY_UNAVAILABLE   ErrorCode = "temporarily_unavailable"
)

func (e ErrorCode) HttpStatus() int {
	switch e {
	case INVALID_REQUEST:
	case UNAUTHORIZED_CLIENT:
	case UNSUPPORTED_RESPONSE_TYPE:
	case INVALID_SCOPE:
		return http.StatusBadRequest
	case ACCESS_DENIED:
		return http.StatusUnauthorized
	case SERVER_ERROR:
		return http.StatusInternalServerError
	case TEMPORARILY_UNAVAILABLE:
		return http.StatusServiceUnavailable
	default:
	}
	return http.StatusInternalServerError
}
