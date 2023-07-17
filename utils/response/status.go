package response

const (
	SuccessCode = iota + 2000
	SuccessDataCode
)

const (
	AuthMissingCode = iota + 4000
	TokenCheckFailedCode
	UserNotExistsCode
	TokenTimeOutCode
	PermissionNotAllowCode
)

const (
	InternalErrorCode = 5000
)

var (
	SuccessMsg            = NewStatus(SuccessCode, "操作成功")
	TokenCheckFailedMsg   = NewStatus(TokenCheckFailedCode, "认证失败")
	TokenTimeOutMsg       = NewStatus(TokenTimeOutCode, "登录失效")
	InternalErrorMsg      = NewStatus(InternalErrorCode, "程序内部错误")
	LoginMissingMsg       = NewStatus(AuthMissingCode, "用户未登录")
	SuccessDataMsg        = NewStatus(SuccessDataCode, "获取数据成功")
	UserNotExistsMsg      = NewStatus(UserNotExistsCode, "用户不存在")
	PermissionNotAllowMsg = NewStatus(PermissionNotAllowCode, "权限不足")
)

type Status struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Detail []string `json:"detail"`
}

func (s *Status) AddDetail(d ...string) Status {
	n := *s
	n.Detail = append(s.Detail, d...)
	return n
}

func NewStatus(code int, msg string) Status {
	return Status{
		code,
		msg,
		[]string{},
	}
}
