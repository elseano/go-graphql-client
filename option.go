package graphql

import (
	"io"
	"net/http"
)

// OptionType represents the logic of graphql query construction
type OptionType string

const (
	// optionTypeOperationName is private because it's option is built-in and unique
	optionTypeOperationName      OptionType = "operation_name"
	optionTypeCustomErrorDecoder OptionType = "custom_error_decoder"
	OptionTypeOperationDirective OptionType = "operation_directive"
)

// Option abstracts an extra render interface for the query string
// They are optional parts. By default GraphQL queries can request data without them
type Option interface {
	// Type returns the supported type of the renderer
	// available types: operation_name and operation_directive
	Type() OptionType
	// String returns the query component string
	String() string
}

// operationNameOption represents the operation name render component
type operationNameOption struct {
	name string
}

func (ono operationNameOption) Type() OptionType {
	return optionTypeOperationName
}

func (ono operationNameOption) String() string {
	return ono.name
}

// OperationName creates the operation name option
func OperationName(name string) Option {
	return operationNameOption{name}
}

type ErrorDecoder func(resp *http.Response, reader io.ReadSeeker) (Errors, error)

type customErrorDecoder struct {
	decoder ErrorDecoder
}

func (ced customErrorDecoder) Type() OptionType {
	return optionTypeCustomErrorDecoder
}

func (ced customErrorDecoder) String() string {
	return "error_decoder"
}

func CustomErrorDecoder(decoder ErrorDecoder) Option {
	return customErrorDecoder{decoder}
}
