package schemas

type WelcomeResponseSchema struct {
	Version string `json:"version"`
	About   string `json:"about"`
}
