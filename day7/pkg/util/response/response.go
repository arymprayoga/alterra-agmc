package response

type Failed struct {
	Success bool   `json:"success" default:"false"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Success struct {
	Success bool   `json:"success" default:"true"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    interface{}
}
