package decode

import (
	"errors"
	"strings"
)

var errContentType = errors.New("unexpected content type")

func isBinaryContentType(contentType string) bool {
	return strings.HasPrefix(contentType, "image/")
}

func isImage(contentType string) bool {
	return strings.HasPrefix(contentType, "image/")
}

func guardedInBody(v interface{}, b []byte, contentType string) error {
	if data, ok := v.(*[]byte); ok {
		if !isBinaryContentType(contentType) {
			return errContentType
		}
		*data = b
		return nil
	}
	return nil
}

func guardedInCondition(v interface{}, b []byte, contentType string) {
	if data, ok := v.(*[]byte); ok && isBinaryContentType(contentType) {
		*data = b
	}
}

func unguarded(v interface{}, b []byte) {
	if data, ok := v.(*[]byte); ok { // want `\*\[\]byte passthrough skips the content-type check: call isBinaryContentType before returning raw bytes`
		*data = b
	}
}

func checkedWithAnotherFunction(v interface{}, b []byte, contentType string) {
	if data, ok := v.(*[]byte); ok && isImage(contentType) { // want `passthrough skips the content-type check`
		*data = b
	}
}

// The check must be part of the branch that the assertion guards.
func checkedBeforeTheBranch(v interface{}, b []byte, contentType string) error {
	if !isBinaryContentType(contentType) {
		return errContentType
	}
	if data, ok := v.(*[]byte); ok { // want `passthrough skips the content-type check`
		*data = b
	}
	return nil
}

func checkedByLocalFunction(v interface{}, b []byte, contentType string) {
	isBinaryContentType := func(string) bool { return true }
	if data, ok := v.(*[]byte); ok && isBinaryContentType(contentType) { // want `passthrough skips the content-type check`
		*data = b
	}
}

func typeSwitchGuarded(v interface{}, b []byte, contentType string) error {
	switch data := v.(type) {
	case *string:
		*data = string(b)
	case *[]byte:
		if !isBinaryContentType(contentType) {
			return errContentType
		}
		*data = b
	}
	return nil
}

func typeSwitchUnguarded(v interface{}, b []byte) {
	switch data := v.(type) {
	case *string:
		*data = string(b)
	case *[]byte: // want `passthrough skips the content-type check`
		*data = b
	}
}

func plainAssertionGuarded(v interface{}, b []byte, contentType string) error {
	data := v.(*[]byte)
	if !isBinaryContentType(contentType) {
		return errContentType
	}
	*data = b
	return nil
}

func plainAssertionUnguarded(v interface{}, b []byte) {
	data := v.(*[]byte) // want `passthrough skips the content-type check`
	*data = b
}

// A function literal is checked on its own, apart from the function around it.
func inFunctionLiteral(v interface{}, b []byte, contentType string) func() {
	binary := isBinaryContentType(contentType)
	return func() {
		if binary {
			*v.(*[]byte) = b // want `passthrough skips the content-type check`
		}
	}
}

func otherTypes(v interface{}) {
	_ = v.([]byte)
	_ = v.(*[]string)
	_ = v.(*string)
	switch v.(type) {
	case []byte, *string:
	}
}
