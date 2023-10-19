package errors

import (
	"regexp"
	"strings"
)

type Entity struct {
	Name string
	ID   string
}

var dbKeyDuplicateErrorPattern = `Error 1062 \(23000\): Duplicate entry '(?P<entry>[\w\d-]+)' for key '(?P<key>[\w.]+)'`

// ParseDuplicateEntry parses the error and returns true if the error is a duplicate entry error.
// It also returns the bindings that caused the error.
func ParseDuplicateEntry(err error) (bool, []string) {
	regex := regexp.MustCompile(dbKeyDuplicateErrorPattern)
	match := regex.FindStringSubmatch(err.Error())

	if len(match) < 3 {
		return false, nil
	}

	var entities []string

	entity := strings.Split(match[2], ".")
	entities = append(entities, entity[0])

	parts := strings.Split(match[1], "-")

	for _, v := range parts {
		if v != "" {
			entities = append(entities, v)
		}
	}

	return true, entities
}

func ParseMissingParentError(err error) bool {
	return strings.HasPrefix(err.Error(), "Error 1452 (23000)")
}
