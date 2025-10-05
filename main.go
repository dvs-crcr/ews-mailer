package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

// Default values
const (
	DefaultRequestTimeout = 10 * time.Second
	DefaultUserAgent      = "MacOutlook/16.16.13.190811"
)

func main() {
	to := flag.String("to", "", "Comma-separated recipients (required)")
	subject := flag.String("subject", "", "Email subject (required)")
	body := flag.String("body", "", "Plaintext message body (required)")
	flag.Parse()

	if *to == "" {
		log.Fatal("Missing required flag: -to")
	}

	if *subject == "" {
		log.Fatal("Missing required flag: -subject")
	}

	if *body == "" {
		log.Fatal("Missing required flag: -body")
	}

	if os.Getenv("EWS_URL") == "" ||
		os.Getenv("EWS_DOMAIN") == "" ||
		os.Getenv("EWS_USERNAME") == "" ||
		os.Getenv("EWS_PASSWORD") == "" {
		log.Fatal("Missing required EWS_* environment variables")
	}

	cfg := Config{
		URL:      os.Getenv("EWS_URL"),
		Domain:   os.Getenv("EWS_DOMAIN"),
		Username: os.Getenv("EWS_USERNAME"),
		Password: os.Getenv("EWS_PASSWORD"),
		From:     os.Getenv("EWS_FROM"),
		To:       strings.Split(*to, ","),
		Subject:  *subject,
		Body:     *body,
		Timeout:  DefaultRequestTimeout,
	}

	message, err := buildMessage(&cfg)
	if err != nil {
		log.Fatalf("Build message failed: %+v", err)
	}

	resp, err := sendMessage(&cfg, message)
	if err != nil {
		log.Fatalf("Send message failed: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Send message failed: %d", resp.StatusCode)
	}
}
