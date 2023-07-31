package e

const (
	Ok                    int = iota
	Failed                    = -1
	BadRequest                = 400
	Unauthorized              = 401
	Forbidden                 = 403
	NotFound                  = 404
	StatusTooManyRequests     = 429
	InternalServerError       = 500

	//用户,公共错误
	RequestError  = 100000
	UserNameError = 100001
	JWTError      = 100002
	UploadError   = 100003
)
