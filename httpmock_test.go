package http_mock_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khanhtc1202/http-mock"
)

type doStubObject struct {
	client *http.Client
}

func (d *doStubObject) RequestHttp(url string) (string, error) {
	r, err := d.client.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	return string(body), nil
}

func ExampleMockClientByHandleFunction() {
	requiredUrl := "http://some-path"
	expectedRes := `whatever`

	mockClient := http_mock.MockClientByHandleFunction(func(req *http.Request) *http.Response {
		// Do stub with request object
		if req.URL.String() == requiredUrl {
			fmt.Println(req.URL.String())
		}

		return &http.Response{
			// Mocked status code
			StatusCode: 200,
			// Mocked response for testing
			Body: ioutil.NopCloser(bytes.NewBufferString(expectedRes)),
			// Mocked header
			// Header: make(http.Header),
		}
	})

	doStubObj := doStubObject{client: mockClient}
	res, _ := doStubObj.RequestHttp(requiredUrl)
	fmt.Println(res)
	// Output: http://some-path
	// whatever
}

func ExampleMockClientByExpectedResponse() {
	expectedResBody := `whatever`
	expectedRes := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(expectedResBody)),
	}

	mockClient := http_mock.MockClientByExpectedResponse(&expectedRes)

	doStubObj := doStubObject{client: mockClient}
	res, _ := doStubObj.RequestHttp("http://....")
	fmt.Println(res)
	// Output: whatever
}
