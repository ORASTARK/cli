package actionerror

import "fmt"

// RevisionNotFoundError is returned when a requested application is not
// found.
type RevisionNotFoundError struct {
	Version string
}

func (e RevisionNotFoundError) Error() string {
	return fmt.Sprintf("Revision '%s' not found", e.Version)
}

type RevisionAmbiguousError struct {
	Version string
}

func (e RevisionAmbiguousError) Error() string {
	return fmt.Sprintf("More than one revision '%s' found", e.Version)
}
