package favourite

import (
	// import local packages
	"dragonfly/app/database"
	"dragonfly/app/model"
)

func ReadFavouriteList() ([]*model.Favourite, error) {
	favouriteList := []*model.Favourite{}
	err := database.SetupDatabase().Find(&favouriteList).Error
	// control flow
	if err != nil {
		// return
		return nil, err
	} else {
		return favouriteList, nil
	}
}
