package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	// Generate 32-byte seed
	seed := make([]byte, 32)
	if _, err := rand.Read(seed); err != nil {
		panic(err)
	}

	privateKey := ed25519.NewKeyFromSeed(seed) // 64 bytes
	publicKey := privateKey[32:]

	full := append(seed, publicKey...)

	privB64 := base64.URLEncoding.EncodeToString(full)
	pubB64 := base64.URLEncoding.EncodeToString(publicKey)

	fmt.Println("PRIVATE=" + privB64)
	fmt.Println("PUBLIC=" + pubB64)
}
