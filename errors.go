package eqtypes

import "fmt"

type ScanError struct {
	got      string
	expected string
}

func (e ScanError) Error() string {
	return fmt.Sprintf("eqtypes: error scanning value, got: %s, expected: %s", e.got, e.expected)
}
