package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func getEmailHash(email string) uint32 {
	hasher := sha256.New()
	hasher.Write([]byte(email))
	hashBytes := hasher.Sum(nil)

	// Tomar los primeros 4 byt>es del hash y convertirlos a un entero sin signo de 32 bits
	hashAsInt := binary.BigEndian.Uint32(hashBytes[:4])
	return hashAsInt
}

func main() {
	email := "example@example.com"
	emailHash := getEmailHash(email)
	fmt.Println("Email:", email)
	fmt.Println("Hash:", emailHash)
}
