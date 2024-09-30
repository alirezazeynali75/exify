package responses

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}
