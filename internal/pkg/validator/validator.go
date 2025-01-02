package validator

import (
	"fmt"
	"net/url"
	"strings"
)

// ValidationError represents an error for a specific field with a message
type ValidationError struct {
	Field   string
	Message string
}

// Error returns the error message for the ValidationError
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// URLValidator validates URLs
type URLValidator struct{}

// NewURLValidator creates a new instance of URLValidator
func NewURLValidator() *URLValidator {
	return &URLValidator{}
}

// ValidateURL validates the given URL string
func (v *URLValidator) ValidateURL(urlStr string) error {
	if strings.TrimSpace(urlStr) == "" {
		return newValidationError("url", "URL cannot be empty")
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil || !parsedURL.IsAbs() {
		return newValidationError("url", "Invalid URL format or URL must be absolute")
	}

	if !isValidScheme(parsedURL.Scheme) {
		return newValidationError("url", "URL must use HTTP or HTTPS protocol")
	}

	return nil
}

// newValidationError creates a new ValidationError
func newValidationError(field, message string) error {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// isValidScheme checks if the URL scheme is valid (HTTP or HTTPS)
func isValidScheme(scheme string) bool {
	scheme = strings.ToLower(scheme)
	return scheme == "http" || scheme == "https"
}
