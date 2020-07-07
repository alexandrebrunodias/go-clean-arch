package response

type Response struct {
	Message string `json:"message"`
}

func NewResponse(message string) *Response {
	return &Response{message}
}

func ItemNotFound() *Response {
	return &Response{"Requested item not found"}
}

func SuccessfullyCreated() *Response {
	return &Response{"Successfully created"}
}

func InternalServerError() *Response {
	return &Response{"Internal server error"}
}