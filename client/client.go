package client

import "fmt"

func wrapError(customMsg string, originalErr error) error {
	return fmt.Errorf("%s: %v", customMsg, originalErr)
}
