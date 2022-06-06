package news

import "github.com/pkg/errors"

var NewsTitleAlreadyExist error = errors.New("News already exist with same title new.")
var NewsNotFound error = errors.New("News not found.")
