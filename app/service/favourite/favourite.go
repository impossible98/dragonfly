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
		// return
		return nil, err
	}
	// return
	return &favourite, nil
}

func ReadFavourite(platform string, roomId string) (*model.Favourite, error) {
	favourite := model.Favourite{
		Platform: platform,
		RoomId:   roomId,
	}
	// control flow
	if err := database.SetupDatabase().First(&favourite).Limit(1).Error; err != nil {
		// return
		return nil, err
	}
	// return
	return &favourite, nil
}
