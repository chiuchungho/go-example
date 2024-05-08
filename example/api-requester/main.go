package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/chiuchungho/go-example/example/api-requester/pkg/requester"
)

func main() {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Info("App started")

	requester := requester.NewRequester(
		&http.Client{
			Timeout: time.Second * 30,
		},
	)
	res, err := requester.GetJsonPlaceHolder(context.Background())
	if err != nil {
		panic(err)
	}
	logger.Info(fmt.Sprintln(res[0]))
	logger.Info("App ended")
}
