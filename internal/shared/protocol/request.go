package protocol

import "encoding/json"

type Auth struct {
	ChallengeID string `json:"challenge_id"`
	Nonce       int    `json:"nonce"`
}

type Request struct {
	ID     int64           `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Auth   *Auth           `json:"auth,omitempty"`
}

func NewRequest(id int64, method string, params any) (*Request, error) {
	paramsData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return &Request{
		ID:     id,
		Method: method,
		Params: paramsData,
	}, nil
}
