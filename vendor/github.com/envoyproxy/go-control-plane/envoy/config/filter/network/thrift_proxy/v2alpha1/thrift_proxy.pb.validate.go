// Code generated by protoc-gen-validate
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on ThriftProxy with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ThriftProxy) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := TransportType_name[int32(m.GetTransport())]; !ok {
		return ThriftProxyValidationError{
			Field:  "Transport",
			Reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := ProtocolType_name[int32(m.GetProtocol())]; !ok {
		return ThriftProxyValidationError{
			Field:  "Protocol",
			Reason: "value must be one of the defined enum values",
		}
	}

	if len(m.GetStatPrefix()) < 1 {
		return ThriftProxyValidationError{
			Field:  "StatPrefix",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetRouteConfig()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return ThriftProxyValidationError{
				Field:  "RouteConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// ThriftProxyValidationError is the validation error returned by
// ThriftProxy.Validate if the designated constraints aren't met.
type ThriftProxyValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ThriftProxyValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftProxy.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ThriftProxyValidationError{}

// Validate checks the field values on ThriftProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ThriftProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := TransportType_name[int32(m.GetTransport())]; !ok {
		return ThriftProtocolOptionsValidationError{
			Field:  "Transport",
			Reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := ProtocolType_name[int32(m.GetProtocol())]; !ok {
		return ThriftProtocolOptionsValidationError{
			Field:  "Protocol",
			Reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ThriftProtocolOptionsValidationError is the validation error returned by
// ThriftProtocolOptions.Validate if the designated constraints aren't met.
type ThriftProtocolOptionsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ThriftProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftProtocolOptions.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ThriftProtocolOptionsValidationError{}