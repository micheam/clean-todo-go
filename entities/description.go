package entities

import (
	"github.com/pkg/errors"
)

type Description string

const (
	MinDescriptionLength = 0
	MaxDescriptionLength = 1000
)

func NewDescription(value string) (desc Description, err error) {
	if len(value) <= MinDescriptionLength {
		err = errors.Errorf("descption must be longer than %d chars", MinDescriptionLength)
		return
	}
	if MaxDescriptionLength <= len(value) {
		err = errors.Errorf("descption must be shorter than %d chars", MaxDescriptionLength)
		return
	}
	desc = Description(value)
	return
}

func (d *Description) String() string {
	return string(*d)
}
