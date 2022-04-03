package post

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

const MaxTagLength = 50

var ErrTagMustNotBeEmpty = errors.New("post tag must not be empty")
var ErrTagIsToLong = fmt.Errorf("tag length cannot be greater than %d characters", MaxTagLength)

type Tag struct {
	value string
}

func newTag(tag string) (Tag, error) {
	if tag == "" {
		return Tag{}, ErrTagMustNotBeEmpty
	}

	if len(tag) > MaxTagLength {
		return Tag{}, ErrTagIsToLong
	}

	return Tag{value: tag}, nil
}

func (t Tag) String() string {
	return t.value
}

type Tags struct {
	tags []Tag
}

func newTags(tags []string) (Tags, error) {
	var result []Tag

	for _, value := range tags {
		tag, err := newTag(value)
		if err != nil {
			return Tags{}, err
		}

		result = append(result, tag)
	}

	return Tags{tags: result}, nil
}

func (t Tags) GetTags() []Tag {
	return t.tags
}

func (t Tags) Value() (driver.Value, error) {
	result := make([]string, 0, len(t.tags))

	for _, tag := range t.tags {
		result = append(result, tag.String())
	}

	return result, nil
}
