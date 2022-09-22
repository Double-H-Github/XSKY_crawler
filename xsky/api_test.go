package xsky

import (
	"log"
	"testing"
)

func TestGetToken(t *testing.T) {
	token, err := GetToken()
	if err != nil {
		log.Fatalf("GetToken err: %v", err)
		return
	}
	log.Printf("token %s", token)
}