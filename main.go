package gonotifier

import (
	"github.com/segmentio/kafka-go"
)

type notifier struct {
	issuer           string
	writer           *kafka.Writer
	sendAttemptCount uint
}

func New(opts *Options) Notifier {
	var addresses []string
	sendAttemptCount := uint(3)

	if opts.SendAttemptCount != 0 {
		sendAttemptCount = opts.SendAttemptCount
	}

	if len(opts.Addresses) == 0 {
		addresses = append(addresses, opts.Address)
	} else {
		addresses = opts.Addresses
		addresses = append(addresses, opts.Address)
	}

	writer := &kafka.Writer{
		Addr:                   kafka.TCP(addresses...),
		Topic:                  opts.Topic,
		AllowAutoTopicCreation: opts.AllowAutoTopicCreation,
	}

	return notifier{
		writer:           writer,
		sendAttemptCount: sendAttemptCount,
		issuer:           opts.IssuerName,
	}
}
