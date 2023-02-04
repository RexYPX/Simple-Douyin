package pack

import (
	"Simple-Douyin/cmd/comment/dal/db"
	"Simple-Douyin/kitex_gen/comment"
)

// 填写除 User 之外的所有字段，填充 User 在 service 中完成
func Comment(m *db.Comment) *comment.Comment {
	if m == nil {
		return nil
	}

	return &comment.Comment{
		Id:         int64(m.ID),
		Content:    m.Content,
		CreateDate: m.CreatedAt.Format("01-02"), // mm-dd
	}
}

// Comments pack list of comment info
func Comments(ms []*db.Comment) []*comment.Comment {
	comments := make([]*comment.Comment, 0)
	for _, m := range ms {
		if n := Comment(m); n != nil {
			comments = append(comments, n)
		}
	}
	return comments
}

func UserIds(ms []*db.Comment) []int64 {
	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserId] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}
