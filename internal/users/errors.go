package users

import "errors"

var (
    ErrEmailNotFound   = errors.New("email not found")
    ErrWrongPassword   = errors.New("wrong password")
    ErrTokenIssue      = errors.New("issue creating token")
)