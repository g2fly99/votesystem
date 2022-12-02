package response

type ResponseT struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

var (
	Success       = ResponseT{200, "success", ""}
	ErrParamWrong = ResponseT{300, "Param Type Not Match", ""}
	ErrNotFound   = ResponseT{400, "Not Found", ""}
	ErrEcIsFinish = ResponseT{401, "Not Found", "ec is finished"}
	ErrSystem     = ResponseT{500, "System Busy", ""}
)

func ResponseHandler(code int, msg string, data interface{}) *ResponseT {
	var ret ResponseT
	ret.Code = code
	ret.Msg = msg
	ret.Data = data

	return &ret
}

func ErrParamHandler(data interface{}) *ResponseT {

	ret := ErrParamWrong
	ret.Data = data

	return &ret
}

func SuccessHandle(data interface{}) ResponseT {
	var ret ResponseT
	ret.Code = 200
	ret.Msg = "successful"
	ret.Data = data

	return ret
}
