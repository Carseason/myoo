package plugins

import (
	"myoo/models"
	"time"
)

type Posts struct {
	Id        int64
	Title     string
	Thumbnail string
	Date      time.Time
	Views     int64
	Category  *models.Category
	Author    PostsAuthor
}
type PostsAuthor struct {
	Id     int64
	Name   string
	Avatar string
}

func CheckThumbnail(value string) string {
	if value == "" {
		return "/src/img/default.jpg"
	}
	return value
}

func CheckPosts(value models.Posts) Posts {
	return Posts{
		Id:        value.Id,
		Title:     value.Title,
		Thumbnail: CheckThumbnail(value.Thumbnail),
		Date:      value.Date,
		Views:     value.Views,
		Category:  value.Category,
		Author: PostsAuthor{
			Id:     value.Author.Id,
			Name:   value.Author.Name,
			Avatar: CheckAvatar(value.Author.Avatar),
		},
	}
}

func CheckPostsAll(value []models.Posts) []Posts {
	var template []Posts
	for v := range value {
		template = append(template, Posts{
			Id:        value[v].Id,
			Title:     value[v].Title,
			Thumbnail: CheckThumbnail(value[v].Thumbnail),
			Date:      value[v].Date,
			Views:     value[v].Views,
			Category:  value[v].Category,
			Author: PostsAuthor{
				Id:     value[v].Author.Id,
				Name:   value[v].Author.Name,
				Avatar: CheckAvatar(value[v].Author.Avatar),
			},
		})
	}
	return template
}
