package pack

import (
	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/kitex_gen/relation"
)

func relationUser2ApiUser(relationUser *relation.User) (apiUser *api.User) {
	apiUser = new(api.User)
	apiUser.ID = relationUser.Id
	apiUser.Name = relationUser.Name
	apiUser.FollowCount = relationUser.FollowCount
	apiUser.FollowerCount = relationUser.FollowerCount
	apiUser.IsFollow = relationUser.IsFollow

	return apiUser
}

func RelationAction2ApiAction(relationResp *relation.RelationActionResponse, apiResp *api.RelationActionResponse) {
	apiResp.StatusCode = relationResp.StatusCode
	apiResp.StatusMsg = relationResp.StatusMsg
}

func RelationFollowList2ApiFollowList(relationResp *relation.RelationFollowListResponse, apiResp *api.RelationFollowListResponse) {
	apiResp.StatusCode = relationResp.StatusCode
	apiResp.StatusMsg = relationResp.StatusMsg
	for _, u := range relationResp.UserList {
		apiResp.UserList = append(apiResp.UserList, relationUser2ApiUser(u))
	}
}

func RelationFollowerList2ApiFollowerList(relationResp *relation.RelationFollowerListResponse, apiResp *api.RelationFollowerListResponse) {
	apiResp.StatusCode = relationResp.StatusCode
	apiResp.StatusMsg = relationResp.StatusMsg
	for _, u := range relationResp.UserList {
		apiResp.UserList = append(apiResp.UserList, relationUser2ApiUser(u))
	}
}

func RelationFriendList2ApiFriendList(relationResp *relation.RelationFriendListResponse, apiResp *api.RelationFriendListResponse) {
	apiResp.StatusCode = relationResp.StatusCode
	apiResp.StatusMsg = relationResp.StatusMsg
	for _, u := range relationResp.UserList {
		apiResp.UserList = append(apiResp.UserList, relationUser2ApiUser(u))
	}
}
