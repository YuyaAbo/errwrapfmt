package testdata

import (
	"errors"
	"fmt"
)

func _() {
	err := errors.New("error")

	_ = fmt.Errorf("error: %w", err) // OK
	_ = fmt.Errorf("error :%w", err) // want "wrong error wrap format"
	_ = fmt.Errorf("error %w", err)  // want "wrong error wrap format"

	_ = "error: %w" // OK
	_ = "error %w"  // want "wrong error wrap format"
}
