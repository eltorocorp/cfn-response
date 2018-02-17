package cfnhelper

// ResponseStatus represents whether the custom resource provider is signalling either success of failure.
type ResponseStatus string

const (
	// ResponseStatusSuccess signifies that a request was handled successfully.
	ResponseStatusSuccess ResponseStatus = "SUCCESS"
	// ResponseStatusFailed signifies that a request failed.
	ResponseStatusFailed ResponseStatus = "FAILED"
)

// Response represents a response to a processed request.
type Response struct {
	Status             *ResponseStatus `json:"ResponseStatus"`
	Reason             *string         `json:"Reason"`
	PhysicalResourceID *string         `json:"PhysicalResourceId"`
	StackID            *string         `json:"StackId"`
	LogicalResourceID  *string         `json:"LogicalResourceId"`
	NoEcho             *bool           `json:"NoEcho"`
	Data               interface{}     `json:"Data"`
}
