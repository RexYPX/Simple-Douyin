package pack

import (
	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/kitex_gen/comment"
)

// User pack user info
func User(u *comment.User) *api.User {
	if u == nil {
		return nil
	}

	return &api.User{
		ID:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}

func Comment(m *comment.Comment) *api.Comment {
	if m == nil {
		return nil
	}

	return &api.Comment{
		ID:         m.Id,
		User:       User(m.User),
		Content:    m.Content,
		CreateDate: m.CreateDate,
	}
}

// Comments pack list of comment info
func Comments(ms []*comment.Comment) []*api.Comment {
	comments := make([]*api.Comment, 0)
	for _, m := range ms {
		if n := Comment(m); n != nil {
			comments = append(comments, n)
		}
	}
	return comments
}
