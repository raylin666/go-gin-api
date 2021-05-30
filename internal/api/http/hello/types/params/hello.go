package params

type HelloReq struct {
	KeyWord string `json:"keyword" form:"keyword" validate:"required" label:"关键词"`
}

type HelloResp struct {
	Message string `json:"message"`
}
