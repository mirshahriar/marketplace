package types

import (
	"github.com/gookit/validate"
)

func (ProductBody) ConfigValidation(v *validate.Validation) {
	v.StringRules(validate.MS{
		"Name":        "required|maxLen:100",
		"Description": "required|maxLen:200",
		"Price":       "gte:0",
	})

	v.AddMessages(validate.MS{
		"required": "You must provide the {field}",
		"gte":      "You must provide a valid {field} greater than or equal to %v",
		"maxLen":   "You must provide a valid {field} with maximum length of %v",
	})
}
