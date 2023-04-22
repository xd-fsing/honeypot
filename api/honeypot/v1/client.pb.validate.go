// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: honeypot/v1/client.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SpyLoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SpyLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SpyLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SpyLoginRequestMultiError, or nil if none found.
func (m *SpyLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SpyLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 10 {
		err := SpyLoginRequestValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 10 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 8 || l > 16 {
		err := SpyLoginRequestValidationError{
			field:  "Password",
			reason: "value length must be between 8 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SpyLoginRequestMultiError(errors)
	}

	return nil
}

// SpyLoginRequestMultiError is an error wrapping multiple validation errors
// returned by SpyLoginRequest.ValidateAll() if the designated constraints
// aren't met.
type SpyLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SpyLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SpyLoginRequestMultiError) AllErrors() []error { return m }

// SpyLoginRequestValidationError is the validation error returned by
// SpyLoginRequest.Validate if the designated constraints aren't met.
type SpyLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpyLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpyLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpyLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpyLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpyLoginRequestValidationError) ErrorName() string { return "SpyLoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e SpyLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpyLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpyLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpyLoginRequestValidationError{}

// Validate checks the field values on SpyLoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SpyLoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SpyLoginReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SpyLoginReplyMultiError, or
// nil if none found.
func (m *SpyLoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SpyLoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if len(errors) > 0 {
		return SpyLoginReplyMultiError(errors)
	}

	return nil
}

// SpyLoginReplyMultiError is an error wrapping multiple validation errors
// returned by SpyLoginReply.ValidateAll() if the designated constraints
// aren't met.
type SpyLoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SpyLoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SpyLoginReplyMultiError) AllErrors() []error { return m }

// SpyLoginReplyValidationError is the validation error returned by
// SpyLoginReply.Validate if the designated constraints aren't met.
type SpyLoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpyLoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpyLoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpyLoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpyLoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpyLoginReplyValidationError) ErrorName() string { return "SpyLoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e SpyLoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpyLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpyLoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpyLoginReplyValidationError{}
