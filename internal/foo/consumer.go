package foo

import (
	"context"
	"encoding/json"

	"gitlab.com/modanisatech/marketplace/shared/kafka"
	"gitlab.com/modanisatech/marketplace/shared/log"
	"go.uber.org/zap"
)

type Consumer struct {
	Logger   *zap.Logger
	Producer kafka.Producer
}

type RepeatEvent struct {
	Count int `json:"count"`
}

// RepeatEventHandler receives kafka events and processes them
// in case of returning error, event will be retried later until max retry count is reached
func (c *Consumer) RepeatEventHandler(event kafka.Event) error {
	ctx := log.Inject(context.Background(), c.Logger)

	repeatEvent := RepeatEvent{}
	err := json.Unmarshal(event.Value, &repeatEvent)
	if err != nil {
		return err
	}

	result, err := fooRepeater(ctx, repeatEvent.Count)
	if err != nil {
		return err
	}

	return c.Producer.Produce(ctx, "results", map[string]string{"result": result})
}
