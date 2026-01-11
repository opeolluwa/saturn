package responses

type ApiResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(message string) ApiResponse {
	return ApiResponse{Message: message}
}
