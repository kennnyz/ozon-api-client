// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.5
// Revision: b9e7d1ac24b2b7f6a5b451fa3d21706ffd8d79e2
// Build Date: 2023-01-30T01:49:43Z
// Built By: goreleaser

package common

import (
	"fmt"
	"strings"
)

const (
	// ContentTypeApplicationJson is a ContentType of type application/json.
	ContentTypeApplicationJson ContentType = "application/json"
)

var ErrInvalidContentType = fmt.Errorf("not a valid ContentType, try [%s]", strings.Join(_ContentTypeNames, ", "))

var _ContentTypeNames = []string{
	string(ContentTypeApplicationJson),
}

// ContentTypeNames returns a list of possible string values of ContentType.
func ContentTypeNames() []string {
	tmp := make([]string, len(_ContentTypeNames))
	copy(tmp, _ContentTypeNames)
	return tmp
}

// String implements the Stringer interface.
func (x ContentType) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x ContentType) IsValid() bool {
	_, err := ParseContentType(string(x))
	return err == nil
}

var _ContentTypeValue = map[string]ContentType{
	"application/json": ContentTypeApplicationJson,
}

// ParseContentType attempts to convert a string to a ContentType.
func ParseContentType(name string) (ContentType, error) {
	if x, ok := _ContentTypeValue[name]; ok {
		return x, nil
	}
	return ContentType(""), fmt.Errorf("%s is %w", name, ErrInvalidContentType)
}

// MarshalText implements the text marshaller method.
func (x ContentType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ContentType) UnmarshalText(text []byte) error {
	tmp, err := ParseContentType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
