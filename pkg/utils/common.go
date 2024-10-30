package utils

import "time"

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
