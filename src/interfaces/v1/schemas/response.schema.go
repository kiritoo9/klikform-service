package schemas

type ResponseSchema struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}
