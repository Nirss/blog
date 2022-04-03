package post

import (
	"time"

	"github.com/google/uuid"
)

//Структура записей блога
type Post struct {
	id          uuid.UUID
	title       Title
	text        Text
	tags        Tags
	createdDate time.Time
}

func NewPost(title string, text string, tags []string, createdDate time.Time) (Post, error) {
	titleVo, err := newTitle(title)
	if err != nil {
		return Post{}, err
	}

	textVo, err := newText(text)
	if err != nil {
		return Post{}, err
	}

	tagsVo, err := newTags(tags)
	if err != nil {
		return Post{}, err
	}

	return Post{
		id:          uuid.New(),
		title:       titleVo,
		text:        textVo,
		tags:        tagsVo,
		createdDate: createdDate,
	}, nil
}

func NewPostWithID(id uuid.UUID, title string, text string, tags []string, createdDate time.Time) (Post, error) {
	post, err := NewPost(title, text, tags, createdDate)
	if err != nil {
		return Post{}, err
	}

	post.id = id

	return post, nil
}

func (p Post) GetId() uuid.UUID {
	return p.id
}

func (p Post) GetTitle() string {
	return p.title.GetTitle()
}

func (p Post) GetText() string {
	return p.text.GetText()
}

func (p Post) GetTags() Tags {
	return p.tags
}

func (p Post) GetCreatedDate() time.Time {
	return p.createdDate
}
