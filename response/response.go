package response

type ResponseT struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

var (
	Success                 = ResponseT{200, "success", ""}
	ErrParamWrong           = ResponseT{300, "Param Type Not Match", ""}
	ErrParamEmailWrong      = ResponseT{301, "Param Email is  invalid", ""}
	ErrParamIdentityNoWrong = ResponseT{302, "Param Identity Number is invalid", ""}
	ErrNotFound             = ResponseT{400, "Not Found", ""}
	ErrIsExist              = ResponseT{401, "Ident or email used", ""}
	ErrEcIsFinish           = ResponseT{402, "Not Found", "ec is finished"}
	ErrHaveVoted            = ResponseT{403, "only can vote 1 times", ""}
	ErrVoteInvalid          = ResponseT{404, "vote the invalid candidate", ""}
	ErrSystem               = ResponseT{500, "System Busy", ""}
	ErrCancdidateNotEnough  = ResponseT{600, "candidate not enough", ""}
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
