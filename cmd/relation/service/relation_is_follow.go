package service

import (
	"Simple-Douyin/cmd/relation/dal/db"
	"Simple-Douyin/kitex_gen/relation"
	"context"
)

type RelationIsFollowService struct {
	ctx context.Context
}

// NewRelationIsFollowService new RelationIsFollowService
func NewRelationIsFollowService(ctx context.Context) *RelationIsFollowService {
	return &RelationIsFollowService{ctx: ctx}
}

// RelationFollowCount return number of users who followed each other
func (s *RelationIsFollowService) RelationIsFollow(req *relation.RelationIsFollowRequest) (bool, error) {
	var resp bool

	userId := req.UserId
	toUserId := req.ToUserId

	isFollow, err := db.QueryIsFollow(s.ctx, userId, toUserId)
	if err != nil {
		return resp, err
	}

	resp = isFollow

	return resp, nil
}
