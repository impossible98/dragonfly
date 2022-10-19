package favourite

import (
	// import local packages
	"dragonfly/app/database"
	"dragonfly/app/model"
)

func CreateFavourite(platform string, roomId string, upper string) (*model.Favourite, error) {
	favourite := model.Favourite{
		Platform: platform,
		RoomId:   roomId,
		Upper:    upper,
	}
	// control flow
	if err := database.SetupDatabase().Create(&favourite).Error; err != nil {
		return nil, err
	}
	return &favourite, nil
}
