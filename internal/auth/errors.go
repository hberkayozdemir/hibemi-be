package auth

import "errors"

var NotAuthorizedError = errors.New("User is not authorized to reach content")
