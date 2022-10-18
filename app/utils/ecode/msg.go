package ecode

var msg = map[int]string{
	Success: "ok",
}

func GetMsg(code int) string {
	return msg[code]
}
