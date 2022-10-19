package bilibili

import (
	// import built-in packages
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// import local packages
	"dragonfly/app/utils/request"
)

type RoomInit struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		RoomID      int  `json:"room_id"`
		ShortID     int  `json:"short_id"`
		UID         int  `json:"uid"`
		NeedP2P     int  `json:"need_p2p"`
		IsHidden    bool `json:"is_hidden"`
		IsLocked    bool `json:"is_locked"`
		IsPortrait  bool `json:"is_portrait"`
		LiveStatus  int  `json:"live_status"`
		HiddenTill  int  `json:"hidden_till"`
		LockTill    int  `json:"lock_till"`
		Encrypted   bool `json:"encrypted"`
		PwdVerified bool `json:"pwd_verified"`
		LiveTime    int  `json:"live_time"`
		RoomShield  int  `json:"room_shield"`
		IsSp        int  `json:"is_sp"`
		SpecialType int  `json:"special_type"`
	} `json:"data"`
}

type Info struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Info struct {
			UID            int    `json:"uid"`
			Uname          string `json:"uname"`
			Face           string `json:"face"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Gender int `json:"gender"`
		} `json:"info"`
		Exp struct {
			MasterLevel struct {
				Level   int   `json:"level"`
				Color   int   `json:"color"`
				Current []int `json:"current"`
				Next    []int `json:"next"`
			} `json:"master_level"`
		} `json:"exp"`
		FollowerNum  int    `json:"follower_num"`
		RoomID       int    `json:"room_id"`
		MedalName    string `json:"medal_name"`
		GloryCount   int    `json:"glory_count"`
		Pendant      string `json:"pendant"`
		LinkGroupNum int    `json:"link_group_num"`
		RoomNews     struct {
			Content   string `json:"content"`
			Ctime     string `json:"ctime"`
			CtimeText string `json:"ctime_text"`
		} `json:"room_news"`
	} `json:"data"`
}

func LiveStatus(roomId string) (bool, string, int) {
	roomInit := RoomInit{}
	url := fmt.Sprintf("https://api.live.bilibili.com/room/v1/Room/room_init?id=%s", roomId)
	resp := request.Get(url, http.Header{})
	err := json.Unmarshal([]byte(resp), &roomInit)
	// control flow
	if err != nil {
		fmt.Println(err)
	}
	// control flow
	if roomInit.Code == 60004 {
		// return
		return false, "", 0
	} else {
		// return
		return true, strconv.Itoa(roomInit.Data.RoomID), roomInit.Data.UID
	}
}

func Upper(uid int) string {
	info := Info{}
	url := fmt.Sprintf("http://api.live.bilibili.com/live_user/v1/Master/info?uid=%s", strconv.Itoa(uid))
	resp := request.Get(url, http.Header{})
	err := json.Unmarshal([]byte(resp), &info)
	// control flow
	if err != nil {
		fmt.Println(err)
	}
	return info.Data.Info.Uname
}
