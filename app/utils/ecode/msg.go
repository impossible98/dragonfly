package ecode

var msg = map[int]string{
	Success:                "ok",
	InvalidParams:          "invalid params",
	ErrorCreateFavourite:   "failed to create favourite",
	ErrorReadFavourite:     "failed to read favourite",
	ErrorUpdateFavourite:   "failed to update favourite",
	ErrorDeleteFavourite:   "failed to delete favourite",
	ErrorReadFavouriteList: "failed to read favourite list",
}

func GetMsg(code int) string {
	return msg[code]
}
