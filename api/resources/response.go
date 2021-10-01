package resources

type ErrorResource struct {
	Message string `json:"message"`
}

func NewErrorResource(message string) *ErrorResource {
	return &ErrorResource{
		Message: message,
	}
}
