package httpapi

import (
	"github.com/Nirss/blog/internal/domain/post"
)

func transformToGetPostsResponse(posts []post.Post) []GetPostResponse {
	var result []GetPostResponse

	for _, value := range posts {
		tags := transformTags(value.GetTags())

		result = append(result, GetPostResponse{
			Id:          value.GetId(),
			Title:       value.GetTitle(),
			Text:        value.GetText(),
			Tags:        tags,
			CreatedDate: value.GetCreatedDate(),
		})
	}

	return result
}

func transformTags(tags post.Tags) []string {
	var result []string

	for _, value := range tags.GetTags() {
		result = append(result, value.String())
	}

	return result
}
