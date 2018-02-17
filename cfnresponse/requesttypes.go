package cfnresponse

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

// CFNRequest represents a request to either create, update, or delete a custom CloudFormation resource.
type CFNRequest struct {
	RequestType       *RequestType
	RequestID         *string
	ResponseURL       *string
	ResourceType      *string
	LogicalResourceID *string
	StackID           *string
	// PhysicalResourceID is only supplied for Update and Delete requests.
	PhysicalResourceID *string
	ResourceProperties interface{}
	// OldResourceProperties is only supplied if the RequestType is Update.
	OldResourceProperties interface{}
}
