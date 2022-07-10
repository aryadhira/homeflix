package helper

type HttpResponse struct {
	Message string
	Data    interface{}
}

func SetHttpResponse(msg string, data interface{}) *HttpResponse {
	response := new(HttpResponse)
	response.Message = msg
	response.Data = data
	return response
}
