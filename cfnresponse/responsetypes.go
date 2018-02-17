package cfnresponse

// ResponseStatus represents whether the custom resource provider is signalling either success of failure.
type ResponseStatus string

const (
	// ResponseStatusSuccess signifies that a request was handled successfully.
	ResponseStatusSuccess ResponseStatus = "SUCCESS"
	// ResponseStatusFailed signifies that a request failed.
	ResponseStatusFailed ResponseStatus = "FAILED"
)

// CFNResponse represents a response to a processed request.
type CFNResponse struct {
	Status             *ResponseStatus
	Reason             *string
	PhysicalResourceID *string
	StackID            *string
	LogicalResourceID  *string
	NoEcho             *bool
	Data               interface{}
}
