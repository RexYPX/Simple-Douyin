package main

import (
	"Simple-Douyin/cmd/publish/service"
	publish "Simple-Douyin/kitex_gen/publish"
	"context"
	"log"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	log.Println("[ypx debug] kitex receive PublishAction Request")
	resp = new(publish.PublishActionResponse)

	if req.UserId < 0 || len(req.Title) <= 0 || len(req.Data) <= 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "Publish Action request inValid"
		log.Println("[ypx debug] kitex PublishAction Request inValid")
		return resp, nil
	}

	log.Println("[ypx debug] kitex prepare to PublishAction rpc")
	err = service.NewPublishActionService(ctx).PublishAction(req)
	if err != nil {
		resp.StatusCode = -2
		resp.StatusMsg = "Publish Action failed"
		log.Println("[ypx debug] kitex PublishAction rpc err")
		return resp, nil
	}
	log.Println("[ypx debug] kitex PublishAction rpc success")
	resp.StatusCode = 0
	resp.StatusMsg = "Publish Action success"
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	log.Println("[ypx debug] kitex receive PublishList Request")

	resp = new(publish.PublishListResponse)

	if req.UserId < 0 || req.MUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "Publish List request inValid"
		log.Println("[ypx debug] kitex PublishList Request inValid")
		return resp, nil
	}

	log.Println("[ypx debug] kitex prepare to service.NewPublishListService(ctx).PublishList(req)")
	videos, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		log.Println("[ypx debug] kitex service.NewPublishListService(ctx).PublishList(req) err")
		resp.StatusCode = -2
		resp.StatusMsg = "Publish List failed"
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "Publish List success"
	resp.VideoList = videos
	log.Println("[ypx debug] kitex PublishList success")
	return resp, nil
}

// PublishIds2List implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishIds2List(ctx context.Context, req *publish.Ids2ListRequest) (resp *publish.Ids2ListResponse, err error) {
	resp, err = service.NewPublishIds2ListService(ctx).PublishIds2List(req)
	return resp, err
}
