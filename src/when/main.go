package main

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	switch os.Getenv("LOGLEVEL") {
	case "DEBUG": log.SetLevel(log.DebugLevel)
	case "INFO": log.SetLevel(log.InfoLevel)
	case "WARNING": log.SetLevel(log.WarnLevel)
	case "WARN": log.SetLevel(log.WarnLevel)
	default: log.SetLevel(log.WarnLevel)
	}
	log.SetLevel(log.DebugLevel)

	log.Debug("Lambda triggered")
}

func Handler(ctx context.Context, event events.CloudWatchEvent) {
	log.WithFields(log.Fields{
		"event": event,
	}).Debug("Lambda triggered")
	//eventj, _ := json.Marshal(event)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}