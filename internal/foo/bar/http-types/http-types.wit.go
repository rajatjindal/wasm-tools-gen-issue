// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package httptypes represents the imported interface "foo:bar/http-types".
package httptypes

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// HTTPStatus represents the u16 "foo:bar/http-types#http-status".
//
//	type http-status = u16
type HTTPStatus uint16

// Body represents the list "foo:bar/http-types#body".
//
//	type body = list<u8>
type Body cm.List[uint8]

// Headers represents the list "foo:bar/http-types#headers".
//
//	type headers = list<tuple<string, string>>
type Headers cm.List[[2]string]

// Params represents the list "foo:bar/http-types#params".
//
//	type params = list<tuple<string, string>>
type Params cm.List[[2]string]

// URI represents the string "foo:bar/http-types#uri".
//
//	type uri = string
type URI string

// Method represents the enum "foo:bar/http-types#method".
//
//	enum method {
//		get,
//		post,
//		put,
//		delete,
//		patch,
//		head,
//		options
//	}
type Method uint8

const (
	MethodGet Method = iota
	MethodPost
	MethodPut
	MethodDelete
	MethodPatch
	MethodHead
	MethodOptions
)

// Request represents the record "foo:bar/http-types#request".
//
//	record request {
//		method: method,
//		uri: uri,
//		headers: headers,
//		params: params,
//		body: option<body>,
//	}
type Request struct {
	Method  Method
	URI     URI
	Headers Headers
	Params  Params
	Body    cm.Option[Body]
}

// Response represents the record "foo:bar/http-types#response".
//
//	record response {
//		status: http-status,
//		headers: option<headers>,
//		body: option<body>,
//	}
type Response struct {
	Status  HTTPStatus
	Headers cm.Option[Headers]
	Body    cm.Option[Body]
}

// HTTPError represents the enum "foo:bar/http-types#http-error".
//
//	enum http-error {
//		success,
//		destination-not-allowed,
//		invalid-url,
//		request-error,
//		runtime-error,
//		too-many-requests
//	}
type HTTPError uint8

const (
	HTTPErrorSuccess HTTPError = iota
	HTTPErrorDestinationNotAllowed
	HTTPErrorInvalidURL
	HTTPErrorRequestError
	HTTPErrorRuntimeError
	HTTPErrorTooManyRequests
)
