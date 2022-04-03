package httpapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mock_service "github.com/Nirss/blog/internal/api/http/mocks"
	"github.com/Nirss/blog/internal/domain/post"
	"github.com/Nirss/blog/internal/libhttp"
)

//go:generate mockgen -destination ./mocks/blog_service.go github.com/Nirss/blog/internal/service BlogService

func TestMakeAddPostHandler(t *testing.T) {
	validPost, err := post.NewPost(
		"Краткая история Dell UNIX",
		"Личные воспоминания одного из разработчиков Dell UNIX об истории создания этой системы",
		[]string{"nix", "История IT", "Старое железо"},
		time.Now())
	require.NoError(t, err)

	type blogSvcMock struct {
		arg post.Post
		err error
	}

	tests := []struct {
		name             string
		req              string
		mock             *blogSvcMock
		wantStatus       int
		wantResponseBody string
	}{
		{
			name: "success",
			req: `{
				"title": "Краткая история Dell UNIX",
				"text": "Личные воспоминания одного из разработчиков Dell UNIX об истории создания этой системы",
				"tags": ["nix", "История IT", "Старое железо"],
				"created_date": "2014-02-01T09:28:56.321-10:00"
			}`,
			mock: &blogSvcMock{
				arg: validPost,
				err: nil,
			},
			wantStatus:       http.StatusCreated,
			wantResponseBody: "",
		},
		{
			name: "invalid input data",
			req: `{
				"title": "",
				"text": "Личные воспоминания одного из разработчиков Dell UNIX об истории создания этой системы",
				"tags": ["nix", "История IT", "Старое железо"],
				"created_date": "2014-02-01T09:28:56.321-10:00"
			}`,
			mock:             nil,
			wantStatus:       http.StatusBadRequest,
			wantResponseBody: "{\"Error\":\"post title must not be empty\"}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mock := mock_service.NewMockBlogService(ctrl)
			handler := MakeAddPostHandler(mock)
			ts := httptest.NewServer(handler)
			defer ts.Close()

			if tt.mock != nil {
				mock.EXPECT().CreateRecord(gomock.Any(), gomock.Any()).Return(tt.mock.err)
			}

			res, err := http.Post(ts.URL, libhttp.ContentTypeJson, strings.NewReader(tt.req))
			require.NoError(t, err)

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err)

			assert.Equal(t, tt.wantStatus, res.StatusCode)
			assert.Equal(t, tt.wantResponseBody, string(body))
		})
	}
}
