package actionerror

import "fmt"

// RevisionNotFoundError is returned when a requested application is not
// found.
type RevisionNotFoundError struct {
	Version int
}

func (e RevisionNotFoundError) Error() string {
	return fmt.Sprintf("Revision '%d' not found", e.Version)
}
