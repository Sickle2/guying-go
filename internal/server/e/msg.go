package e

var msgMap = map[int]string{
	Ok:                    "ok",
	Failed:                "error",
	BadRequest:            "请求参数错误",
	Unauthorized:          "认证失败",
	Forbidden:             "权限不足",
	StatusTooManyRequests: "请求过多，请稍后重试",
	NotFound:              "路径不存在",
	InternalServerError:   "服务器内部错误",

	RequestError:  "请求错误",
	UserNameError: "用户名或者密码错误",
	JWTError:      "生成token出错",
	UploadError:   "上传文件失败",
}

func GetMessage(code int) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}

	return msgMap[Failed]
}
