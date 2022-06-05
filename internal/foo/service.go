package foo

import (
	"context"

	"gitlab.com/modanisatech/marketplace/shared/errors"
	"gitlab.com/modanisatech/marketplace/shared/log"
)

var ErrNegativeRepeatCount = errors.BadRequest("repeat count can not be negative")

func fooRepeater(ctx context.Context, repeatCount int) (string, error) {
	logger := log.FromContext(ctx)

	if repeatCount < 0 {
		return "", ErrNegativeRepeatCount
	}

	result := ""
	for i := 0; i < repeatCount; i++ {
		result += "foo"
	}

	logger.Info("foo successfully repeated")

	return result, nil
}
