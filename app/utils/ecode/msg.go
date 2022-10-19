package ecode

var msg = map[int]string{
	Success:              "ok",
	InvalidParams:        "invalid params",
	ErrorCreateFavourite: "failed to create favourite",
}

func GetMsg(code int) string {
	return msg[code]
}
