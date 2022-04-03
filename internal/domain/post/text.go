package post

import (
	"errors"
	"fmt"
)

// MaxTextLength Примерное максимальное количество символов в тексте поста
const MaxTextLength = 10000

var ErrTextMustNotBeEmpty = errors.New("post text must not be empty")
var ErrTextIsToLong = fmt.Errorf("text length cannot be greater than %d characters", MaxTextLength)

type Text struct {
	value string
}

func newText(text string) (Text, error) {
	if text == "" {
		return Text{}, ErrTextMustNotBeEmpty
	}

	if len(text) > MaxTextLength {
		return Text{}, ErrTextIsToLong
	}

	return Text{value: text}, nil
}

func (t Text) GetText() string {
	return t.value
}
