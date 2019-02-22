package error

var MsgFlags = map[int]string{
	SUCCESS:            "ok",
	ERROR:              "fail",
	ERROR_AUTH_INVALID: "invalid token",
	ERROR_TOKEN_NONE:   "token is none",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

func GetMsgWithError(code int, err error) string {
	msg, ok := MsgFlags[code]
	if ok {
		if err == nil {
			return msg
		} else {
			return msg + ",Error: " + err.Error()
		}
	}
	return MsgFlags[ERROR]
}
