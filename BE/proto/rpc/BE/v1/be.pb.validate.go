// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: BE/v1/be.proto

package iam

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

// Validate checks the field values on StringProcessRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StringProcessRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringProcessRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StringProcessRequestMultiError, or nil if none found.
func (m *StringProcessRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StringProcessRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := StringProcessRequestValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StringProcessRequestMultiError(errors)
	}

	return nil
}

// StringProcessRequestMultiError is an error wrapping multiple validation
// errors returned by StringProcessRequest.ValidateAll() if the designated
// constraints aren't met.
type StringProcessRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringProcessRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringProcessRequestMultiError) AllErrors() []error { return m }

// StringProcessRequestValidationError is the validation error returned by
// StringProcessRequest.Validate if the designated constraints aren't met.
type StringProcessRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringProcessRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringProcessRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringProcessRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringProcessRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringProcessRequestValidationError) ErrorName() string {
	return "StringProcessRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StringProcessRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringProcessRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringProcessRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringProcessRequestValidationError{}

// Validate checks the field values on StringProcessResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StringProcessResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringProcessResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StringProcessResponseMultiError, or nil if none found.
func (m *StringProcessResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StringProcessResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StringProcessResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StringProcessResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StringProcessResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StringProcessResponseMultiError(errors)
	}

	return nil
}

// StringProcessResponseMultiError is an error wrapping multiple validation
// errors returned by StringProcessResponse.ValidateAll() if the designated
// constraints aren't met.
type StringProcessResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringProcessResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringProcessResponseMultiError) AllErrors() []error { return m }

// StringProcessResponseValidationError is the validation error returned by
// StringProcessResponse.Validate if the designated constraints aren't met.
type StringProcessResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringProcessResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringProcessResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringProcessResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringProcessResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringProcessResponseValidationError) ErrorName() string {
	return "StringProcessResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StringProcessResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringProcessResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringProcessResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringProcessResponseValidationError{}

// Validate checks the field values on StringData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StringData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StringDataMultiError, or
// nil if none found.
func (m *StringData) ValidateAll() error {
	return m.validate(true)
}

func (m *StringData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return StringDataMultiError(errors)
	}

	return nil
}

// StringDataMultiError is an error wrapping multiple validation errors
// returned by StringData.ValidateAll() if the designated constraints aren't met.
type StringDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringDataMultiError) AllErrors() []error { return m }

// StringDataValidationError is the validation error returned by
// StringData.Validate if the designated constraints aren't met.
type StringDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringDataValidationError) ErrorName() string { return "StringDataValidationError" }

// Error satisfies the builtin error interface
func (e StringDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringDataValidationError{}