package http_mock

import "net/http"

// An implement of Transport interface (from net/http Client)
//
// Transport specifies the mechanism by which individual
// HTTP requests are made.
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
//
// Override it with own implementation to control http request
// with the interface:
//
// input: http.Request
// output: http.Response
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// Mock http.Client by HandleFunction
//
// returns *http.Client with Transport replaced with handle mechanism that
// passed as function parameter.
func MockHandleClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// Mock http.Client by ExpectedResponse
//
// returns *http.Client with Transport replaced and always returns
// defined http.Response as a response from http call.
func MockResponseClient(res *http.Response) *http.Client {
	return MockHandleClient(func(req *http.Request) *http.Response {
		return res
	})
}
