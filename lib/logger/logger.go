package logger

import (
	"fmt"
)

func LogError(msg string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", msg, err)
}

func Log(msg string) error {
	if  msg == "" {
		return fmt.Errorf("%s", "Message cannot be empty")
	}

	fmt.Printf("%s", msg)
	return nil
}