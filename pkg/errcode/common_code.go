package errcode


var (
	Success                           = NewError(0, "成功")
	ServerError                       = NewError(1000000, "服务器内部错误")
	InvalidParams                     = NewError(1000001, "入参错误")
	NotFound                          = NewError(1000002, "找不到")
	UnauthorizedAuthNotExist          = NewError(1000003, "鉴权失败，找不到对应的appkey")
	UnauthorizedTokenError            = NewError(1000004, "鉴权失败，toke错误")
	UnauthorizedTokenTimeOut          = NewError(1000005, "鉴权失败，toke超时")
	UnauthorizedTokenGenerate         = NewError(1000006, "鉴权失败，toke生成失败")
	TooManyRequest                    = NewError(1000007, "请求过多")
)
