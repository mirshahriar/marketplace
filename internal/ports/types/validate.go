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

	// These are global messages that will be used by all validators
	validate.AddGlobalMessages(validate.MS{
		"required":        "You must provide the {field}",
		"positive_number": "You must provide a valid {field}",
		"gt":              "You must provide a valid {field}",
		"valid_time":      "You must provide a valid {field}",
	})

	validate.AddValidator("positive_number", func(value interface{}) (res bool) {
		if value == nil || validate.IsNilObj(value) {
			return true
		}

		switch vt := value.(type) {
		case *int64:
			return *vt > 0
		case int64:
			return vt > 0
		}
		return false
	})

	validate.AddValidator("valid_time", func(value interface{}) (res bool) {
		var militaryTime int64
		switch vt := value.(type) {
		case *int64:
			militaryTime = *vt
		case int64:
			militaryTime = vt
		}

		if militaryTime < 0 || militaryTime > 2359 {
			return false
		}
		hour := militaryTime / 100
		minute := militaryTime % 100

		if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
			return false
		}

		return true
	})
}

func validateStruct(data interface{}, scene string) validate.Errors {
	if v := validate.Struct(data); v != nil {
		if !v.Validate(scene) && !v.Errors.Empty() {
			return v.Errors
		}
	}
	return nil
}

// Validate validates the data against the scene
func Validate(data interface{}, scene string) errors.Error {
	var list []interface{}
	switch data.(type) {
	default:
		list = append(list, data)
	}

	for _, v := range list {
		errData := validateStruct(v, scene)
		if errData != nil {
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
