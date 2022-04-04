package httpapi

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Nirss/blog/internal/domain/post"
	"github.com/Nirss/blog/internal/libhttp"
	"github.com/Nirss/blog/internal/service"
	"github.com/gorilla/mux"
)

type PostRecordRequest struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Tags        []string  `json:"tags"`
	CreatedDate time.Time `json:"created_date"`
}

func MakeHandlers(blogSvc service.BlogService) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/posts", MakeGetPostsHandler(blogSvc)).Methods(http.MethodGet)
	r.HandleFunc("/posts", MakeAddPostHandler(blogSvc)).Methods(http.MethodPost)

	return r
}

func MakeAddPostHandler(blogSvc service.BlogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostRecordRequest

		if err := libhttp.JsonDecode(r, &req); err != nil {
			libhttp.SendError(w, http.StatusBadRequest, err)
			return
		}

		record, err := post.NewPost(req.Title, req.Text, req.Tags, req.CreatedDate)
		if err != nil {
			libhttp.SendError(w, http.StatusBadRequest, err)
			return
		}

		err = blogSvc.CreatePost(r.Context(), record)
		if err != nil {
			libhttp.SendError(w, http.StatusBadRequest, err)
			return
		}

		libhttp.SendResponse(w, http.StatusCreated)
	}
}

type GetPostResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Tags        []string  `json:"tags"`
	CreatedDate time.Time `json:"created_date"`
}

func MakeGetPostsHandler(blogSvc service.BlogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := blogSvc.GetAllPosts(r.Context())
		if err != nil {
			libhttp.SendError(w, http.StatusInternalServerError, err)
			return
		}

		result := transformToGetPostsResponse(records)

		libhttp.JsonEncode(w, http.StatusOK, result)
	}
}
