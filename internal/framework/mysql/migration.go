package db

import (
	"github.com/blang/semver"
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/helper/utils"
	"github.com/mirshahriar/marketplace/internal/framework/mysql/migration"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// Migrate migrates the database to the latest version
func (a Adapter) Migrate() errors.Error {
	err := a.db.AutoMigrate(&types.SystemConfig{})
	if err != nil {
		return errors.InternalDBError(err)
	}

	var system types.SystemConfig
	rowAffected := a.db.Model(&types.SystemConfig{}).
		Where(&types.SystemConfig{Name: "DatabaseVersion"}).
		Find(&system).RowsAffected

	if rowAffected == 0 {
		system.Value = "0.0.0"
	}

	currentVersion, err := semver.Parse(system.Value)
	if err != nil {
		return errors.DBMigrationError(err)
	}

	return migration.Migrate(a.db, currentVersion)
}

func (a Adapter) CreateFakeUser() (types.Token, error) {
	password := utils.RandString(15)

	encryptedPass, err := utils.Encrypt(a.appConfig.PasswordEncryptionKey, password)
	if err != nil {
		return types.Token{}, err
	}

	user := types.User{
		Username: utils.RandString(5),
		Password: encryptedPass,
		Email:    utils.RandString(10) + "@gmail.com",
	}

	err = a.db.Create(&user).Error
	if err != nil {
		return types.Token{}, err
	}

	randToken := utils.RandString(16)
	encryptedToken, err := utils.Encrypt(a.appConfig.TokenEncryptionKey, randToken)
	if err != nil {
		return types.Token{}, err
	}

	token := types.Token{
		UserID: user.ID,
		Token:  encryptedToken,
	}

	err = a.db.Create(&token).Error
	if err != nil {
		return types.Token{}, err
	}

	token.Token = randToken

	return token, nil
}
