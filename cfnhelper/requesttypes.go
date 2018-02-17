package cfnhelper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RequestType represents which type of action CloudFormation would like to perform against the specified resource.
type RequestType string

const (
	// RequestTypeCreate signifies that a request is expectd to create a resource.
	RequestTypeCreate RequestType = "Create"
	// RequestTypeUpdate signifies that a request is to update a resource.
	RequestTypeUpdate RequestType = "Update"
	// RequestTypeDelete signifies that a request is to delete a resource.
	RequestTypeDelete RequestType = "Delete"
)

// Request represents a request to either create, update, or delete a custom CloudFormation resource.
//  - PhysicalResourceID is only supplied for Update and Delete requests.
//  - OldResourceProperties is only supplied if the RequestType is Update.
type Request struct {
	RequestType           *RequestType `json:"RequestType"`
	RequestID             *string      `json:"RequestId"`
	ResponseURL           *string      `json:"ResponseURL"`
	ResourceType          *string      `json:"ResourceType"`
	LogicalResourceID     *string      `json:"LogicalResourceId"`
	StackID               *string      `json:"StackId"`
	PhysicalResourceID    *string      `json:"PhysicalResourceId"`
	ResourceProperties    interface{}  `json:"ResourceProperties"`
	OldResourceProperties interface{}  `json:"OldResourceProperties"`
}

// SendResponse posts a response back to the request's ResponseURL.
func (request *Request) SendResponse(response *Response) (*http.Response, error) {
	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	r, err := http.Post(*request.ResponseURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return r, nil
}
