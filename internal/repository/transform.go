package repository

import (
	"encoding/json"

	"github.com/Nirss/blog/internal/domain/post"
)

func transformPosts(posts []PostDTO) ([]post.Post, error) {
	var result []post.Post

	for _, value := range posts {
		var tags []string

		err := json.Unmarshal([]byte(value.Tags), &tags)
		if err != nil {
			return nil, err
		}

		postDto, err := post.NewPostWithID(value.Id, value.Title, value.Text, tags, value.CreatedDate)
		if err != nil {
			return nil, err
		}

		result = append(result, postDto)
	}

	return result, nil
}
