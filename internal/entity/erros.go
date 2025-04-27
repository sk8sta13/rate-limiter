package entity

import "errors"

var (
	ErrIPExceededMaxRequests   = errors.New("IP blocked for exceeding maximum request")
	ErrTokenExceededMaxRequest = errors.New("API TOKEN blocked for exceeding maximum request")
	ErrTokenUnauthorized       = errors.New("Access denied")
)
