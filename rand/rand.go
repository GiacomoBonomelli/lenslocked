package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	//creo una slice di byte
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes:%w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("bytes: didn't read enough random bytes")
	}
	return b, nil
}

// Quando dovrò creare dei token per le sessions, chiamo solamente questa funzione
func String(n int) (string, error) {
	//fa un encoding della slice restituita dalla funz. Bytes
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string:%w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
