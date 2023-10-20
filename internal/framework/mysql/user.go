package db

import (
	gError "errors"

	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
	"gorm.io/gorm"
)

func (a Adapter) GetUserByToken(token string) (types.User, bool, errors.Error) {
	var user types.User

	db := a.db.Model(&user)
	db = db.Joins("JOIN tokens ON tokens.user_id = users.id")
	db = db.Where("tokens.token = ?", token)

	if err := db.First(&user).Error; err != nil {
		if gError.Is(err, gorm.ErrRecordNotFound) {
			return types.User{}, false, nil
		}

		return types.User{}, false, errors.InternalDBError(err)
	}

	return user, true, nil
}
