package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"time"
)

func DoWithAttemps(fn func() error, maxAttemps int, delay time.Duration) error {
	var err error

	for maxAttemps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttemps--
			continue
		}

		return nil
	}

	return err
}

func HashValue(value []byte, salt []byte) string {
	rawBytes := append(value, salt...)
	hasher := sha512.New()
	hasher.Write(rawBytes)
	hashedBytes := hasher.Sum(nil)
	hashedHex := hex.EncodeToString(hashedBytes)
	return hashedHex
}
