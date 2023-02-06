package main

import (
	favorite "Simple-Douyin/cmd/favorite/kitex_gen/favorite"
	"Simple-Douyin/cmd/favorite/service"
	"context"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.FavoriteActionResponse)

	if len(req.Token) == 0 || req.ActionType <= 0 || req.ActionType >= 3 {
		resp.StatusCode = 1
		resp.StatusMsg = "点赞不合法,token、_id不能空,action_type只能是1或者2"
		return resp, nil
	}

	err = service.NewActionFavoriteService(ctx).ActionFavorite(req)

	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "点赞失败"
		return resp, nil
	}
	resp.StatusCode = 0
	resp.StatusMsg = "点赞成功"
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.FavoriteListResponse)
	if req.UserId < 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "UserId非法"
		return resp, nil
	}

	video_list := make([]*favorite.Video, 0)
	video_list, err = service.NewFavoriteListService(ctx).FavoriteList(req)

	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "拉取点赞视频失败"
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "拉取点赞视频成功"
	resp.VideoList = video_list
	return resp, nil
}
