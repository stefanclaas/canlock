package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func generateCancelKey(secretKey, messageID string) string {
	key := []byte(secretKey)
	message := []byte(messageID)

	h := hmac.New(sha256.New, key)
	h.Write(message)

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func generateCancelLock(cancelKey string) string {
	key := []byte(cancelKey)

	h := sha256.New()
	h.Write(key)

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	var secretKey, messageID string
	var generateLock, generateKey bool

	flag.StringVar(&secretKey, "secret", "", "Secret key for the HMAC calculation")
	flag.StringVar(&messageID, "message-id", "", "Message-ID of the article")
	flag.BoolVar(&generateLock, "generate-lock", false, "Generate Cancel-Lock")
	flag.BoolVar(&generateKey, "generate-key", false, "Generate Cancel-Key")

	flag.Parse()

	if secretKey == "" || messageID == "" {
		fmt.Println("Error: Secret key and Message-ID must be specified.")
		os.Exit(1)
	}

	if generateKey {
		key := generateCancelKey(secretKey, messageID)
		fmt.Printf("Cancel-Key: sha256:%s\n", key)
	}

	if generateLock {
		key := generateCancelKey(secretKey, messageID)
		lock := generateCancelLock(key)
		fmt.Printf("Cancel-Lock: sha256:%s\n", lock)
	}
}

