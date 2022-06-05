package user

import "github.com/pkg/errors"

var UserAlreadyExistError error = errors.New("User already exists")
var UserNotFound error = errors.New("User not found!")
var UserAlreadyActivated error = errors.New("User already activated!")
