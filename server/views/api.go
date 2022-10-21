package views

type Response struct {
	Message string      `json:"message" example:"GET_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Token   string      `json:"token,omitempty"`
}
