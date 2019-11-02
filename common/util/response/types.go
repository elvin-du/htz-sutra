package response

type ReturnCode int

const (
	SuccessCode             ReturnCode = 200
	NotFoundCode            ReturnCode = 404
	InternalServerErrorCode ReturnCode = 500
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Err  interface{} `json:"error"`
}

func New(code ReturnCode, data interface{}, error interface{}) Response {
	return Response{int(code), data, error}
}

func Ok(data interface{}) Response {
	return New(SuccessCode, data, nil)
}

func NotFound(err interface{}) Response {
	return New(NotFoundCode, nil, err)

}

func InternalServerError(err interface{}) Response {
	return New(InternalServerErrorCode, nil, err)
}

func Error(code ReturnCode, error interface{}) Response {
	return New(code, nil, error)
}
