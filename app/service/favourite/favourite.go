package favourite

import (
	// import built-in packages
	"errors"
	// import local packages
	"dragonfly/app/database"
	"dragonfly/app/model"
	"dragonfly/lib/bilibili"
)

func CreateFavourite(platform string, roomId string) (*model.Favourite, error) {
	// control flow
	if platform == "bilibili" {
		liveStatus, newRoomId, uid := bilibili.LiveStatus(roomId)
		// control flow
		if liveStatus {
			newUpper := bilibili.Upper(uid)
			favourite := model.Favourite{
				Platform: platform,
				RoomId:   newRoomId,
				Upper:    newUpper,
			}
			err := database.SetupDatabase().First(&favourite).Error
			// control flow
			if err != nil {
				// control flow
				err := database.SetupDatabase().Create(&favourite).Error
				if err != nil {
					// return
					return nil, err
				} else {
					// return
					return &favourite, nil
				}
			} else {
				return nil, errors.New("data is existed")
			}
		} else {
			return nil, errors.New("live room is not existed")
		}
	} else {
		// return
		return nil, errors.New("plaform is not existed")
	}
}

func ReadFavourite(platform string, roomId string) (*model.Favourite, error) {
	favourite := model.Favourite{
		Platform: platform,
		RoomId:   roomId,
	}
	err := database.SetupDatabase().First(&favourite).Error
	// control flow
	if err != nil {
		return nil, err
	} else {
		return &favourite, nil
	}
}

func UpdateFavourite(platform string, roomId string) (*model.Favourite, error) {
	// control flow
	if platform == "bilibili" {
		liveStatus, newRoomId, uid := bilibili.LiveStatus(roomId)
		// control flow
		if liveStatus {
			newUpper := bilibili.Upper(uid)
			favourite := model.Favourite{
				Platform: platform,
				RoomId:   newRoomId,
				Upper:    newUpper,
			}
			err := database.SetupDatabase().First(&favourite).Error
			// control flow
			if err != nil {
				// return
				return nil, err
			} else {
				if err := database.SetupDatabase().Save(&favourite).Error; err != nil {
					// return
					return nil, err
				} else {
					// return
					return &favourite, nil
				}
			}
		} else {
			return nil, errors.New("live room is not existed")
		}
	} else {
		// return
		return nil, errors.New("plaform is not existed")
	}
}

func DeleteFavourite(platform string, roomId string) (*model.Favourite, error) {
	favourite := model.Favourite{
		Platform: platform,
		RoomId:   roomId,
	}
	// control flow
	if err := database.SetupDatabase().First(&favourite).Error; err != nil {
		// return
		return nil, err
	}
	// control flow
	if err := database.SetupDatabase().Delete(&favourite).Error; err != nil {
		// return
		return nil, err
	}
	// return
	return &favourite, nil
}
