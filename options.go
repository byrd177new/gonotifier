package gonotifier

type Options struct {
	IssuerName             string
	Address                string
	Addresses              []string
	Topic                  string
	AllowAutoTopicCreation bool
	SendAttemptCount       uint
}
