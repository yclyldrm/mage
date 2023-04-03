package models

import "errors"

var (
	ErrINVALID_USERNAME_OR_PASSWORD = errors.New("invalid username or password")
	ErrLOGIN_ERROR                  = errors.New("login failed")
	ErrINVALID_TOKEN_ERROR          = errors.New("invalid token")

	TIMEFORMAT = "2006-01-02 15:04:05"
)
