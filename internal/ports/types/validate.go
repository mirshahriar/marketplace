package types

import (
	"github.com/gookit/validate"
	"github.com/mirshahriar/marketplace/helper/errors"
)

func init() {
	validate.Config(func(opt *validate.GlobalOption) {
		// StopOnError will stop validation when an error occurred
		opt.StopOnError = true
		// SkipOnEmpty will skip validation when the value is empty
		opt.SkipOnEmpty = true
		opt.UpdateSource = false
	})
}

// Validate validates the data against the scene
func Validate(data interface{}, scene string) errors.Error {
	if v := validate.Struct(data); v != nil {
		if !v.Validate(scene) && !v.Errors.Empty() {
			errData := v.Errors
			for _, val := range errData {
				for _, message := range val {
					if message != "-" {
						return errors.ValidationError(message)
					}
				}
			}
			return errors.ValidationError(errData.One())
		}
	}
	return nil
}
