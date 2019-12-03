package entities

import (
	"github.com/pkg/errors"
)

type Title string

const (
	MinTitleLength = 1
	MaxTitleLength = 128
)

func NewTitle(in string) (title Title, err error) {
	if len(in) <= MinTitleLength {
		err = errors.Errorf("title must be longer than %d chars", MinTitleLength)
		return
	}
	if MaxTitleLength <= len(in) {
		err = errors.Errorf("title must be shorter than %d chars", MaxTitleLength)
		return
	}
	title = Title(in)
	return
}

func (t Title) String() string {
	return string(t)
}
