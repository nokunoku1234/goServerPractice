package validator

import "goServerPractice/internal/transport"

type ValidationError struct {
	Details []transport.FieldIssue
}

func (e *ValidationError) Error() string {
	return "validation failed"
}
