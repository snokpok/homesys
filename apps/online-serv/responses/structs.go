package responses

type ResponseMeta struct {
	Msg string `json:"msg" xml:"msg" form:"msg"`
}

type SuccessfulResponse struct {
	Meta ResponseMeta `json:"meta" xml:"meta" form:"meta"`
	Data interface{}  `json:"data" xml:"data" form:"data"`
}
