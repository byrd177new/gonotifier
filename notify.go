package gonotifier

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Notifier interface {
	Notify(context.Context, Message) error
}

func (n notifier) Notify(ctx context.Context, message Message) error {
	byteMessage, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	notificationKey := fmt.Sprintf("%s:%v", message.UserID, time.Now())
	if n.issuer != "" {
		notificationKey = fmt.Sprintf("%s:%s", n.issuer, notificationKey)
	}

	kafkaMessage := kafka.Message{
		Key:   []byte(notificationKey),
		Value: byteMessage,
	}

	for i := 1; i <= int(n.sendAttemptCount); i++ {
		err := n.writer.WriteMessages(ctx, kafkaMessage)
		if err == nil {
			break
		}

		if err != nil && i == int(n.sendAttemptCount) {
			return fmt.Errorf("%d/%d attempts done. Last error reason: %w", i, n.sendAttemptCount, err)
		}
	}

	return nil
}
