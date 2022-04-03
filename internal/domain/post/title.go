package post

import (
	"errors"
	"fmt"
)

const MaxTitleLength = 250

var ErrTitleMustNotBeEmpty = errors.New("post title must not be empty")
var ErrTitleIsToLong = fmt.Errorf("title length cannot be greater than %d characters", MaxTitleLength)

type Title struct {
	value string
}

func newTitle(title string) (Title, error) {
	if title == "" {
		return Title{}, ErrTitleMustNotBeEmpty
	}

	if len(title) > MaxTitleLength {
		return Title{}, ErrTitleIsToLong
	}

	return Title{value: title}, nil
}

func (t Title) GetTitle() string {
	return t.value
}
