package schemas

type WorkspaceBodySchema struct {
	Title        string `json:"title" validate:"required"`
	Descriptions string `json:"descriptions" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Remark       string `json:"remark"`
}

type WorkspaceResponseSchema struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Status       string `json:"status"`
	Remark       string `json:"remark"`
}
