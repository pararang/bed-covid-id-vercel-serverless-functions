package main

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://0795802312194014b5e771fbb4f30902@o205268.ingest.sentry.io/5868996",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
