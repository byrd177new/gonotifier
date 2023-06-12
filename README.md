# Notification Server Client for Golang

## 1. Installation

```shell
go get github.com/onemgvv/gonotifier
```

## 2. QuickStart

```go
	opts := &gonotifier.Options{
		IssuerName:             "orders",
		Address:                "https://notifications.yoursite.com/queue",
		Topic:                  "notifications",
		AllowAutoTopicCreation: true,
		SendAttemptCount:       5,
	}

  notifier := gonotifier.New(opts)

  message := gonotifier.Message{
    UserID: "uuid",
    Status: gonotifier.Success,
    Header: "some data here",
    Body: "another data here"
  }

  err := notifier.Notify(context.Background(), message)
  if err != nil {
    // do something
  }
```
