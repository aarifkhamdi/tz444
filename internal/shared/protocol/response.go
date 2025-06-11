package protocol

import "encoding/json"

type Response struct {
	ID     int64           `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponse(id int64, result any) (*Response, error) {
	resultData, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return &Response{
		ID:     id,
		Result: resultData,
	}, nil
}

func NewErrorResponse(id int64, code int, message string) *Response {
	return &Response{
		ID: id,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}
}
