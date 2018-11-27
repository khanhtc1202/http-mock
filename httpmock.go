package http_mock

import "net/http"

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// Mock http.Client by HandleFunction
// returns *http.Client with Transport replaced to avoid making real calls
func MockHandleClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// Mock http.Client by ExpectedResponse
// returns *http.Client with declared response
func MockResponseClient(res *http.Response) *http.Client {
	return MockHandleClient(func(req *http.Request) *http.Response {
		return res
	})
}
