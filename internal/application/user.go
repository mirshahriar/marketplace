package app

import (
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/helper/utils"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// GetUserByToken returns a user by its token. This is used by authentication middleware
func (a Adapter) GetUserByToken(token string) (types.LoggedInUser, bool, errors.Error) {
	// We keep the token in the database encrypted, so we need to encrypt the provided token
	encryptToken, err := utils.Encrypt(a.config.TokenEncryptionKey, token)
	if err != nil {
		return types.LoggedInUser{}, false, errors.InternalError(err)
	}

	// Getting the user by the encrypted token
	user, exists, cErr := a.db.GetUserByToken(encryptToken)
	if cErr != nil {
		return types.LoggedInUser{}, exists, cErr
	}

	if !exists {
		return types.LoggedInUser{}, exists, nil
	}

	// If the user exists, we return the user ID
	return types.LoggedInUser{ID: user.ID}, true, nil
}
