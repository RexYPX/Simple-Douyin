// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package constants

const (
	UserTableName     = "user"
	CommentTableName  = "comment"
	RelationTableName = "relation"
	MessageTableName  = "message"
	FavoriteTableName = "favorite"
	VideoTableName    = "video"

	SecretKey   = "secret key"
	IdentityKey = "id"
	Token       = "JWT_TOKEN"

	Comments = "comments"

	ApiServiceName      = "api"
	FavoriteServiceName = "favorite"
	CommentServiceName  = "comment"
	UserServiceName     = "user"
	MessageServiceName  = "message"
	RelationServiceName = "relation"
	FeedServiceName     = "feed"
	PublishServiceName  = "publish"

	// UserServiceAddr     = ":8866"
	// CommentServiceAddr  = ":8888"
	// FavoriteServiceAddr = ":8989"
	// MessageServicePort  = ":8801"
	// RelationServiceAddr = ":9999"
	// FeedServiceAddr     = ":9901"
	// PublishServiceAddr  = ":5557"
	// HertzServiceAddr    = ":8080"
	FileServerAddr = ":8000"
	HertzServiceIP = "124.223.111.3"

	UserServiceAddr     = ":5555"
	CommentServiceAddr  = ":5558"
	FavoriteServiceAddr = ":5989"
	MessageServicePort  = ":5801"
	RelationServiceAddr = ":5999"
	FeedServiceAddr     = ":5901"
	PublishServiceAddr  = ":5557"
	HertzServiceAddr    = ":8081"

	TCP = "tcp"

	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"

	EtcdAddress    = "127.0.0.1:2379"
	ExportEndpoint = ":4317"

	DefaultLimit = 10

	MaxFeed        = 30
	MaxPublishSize = 1000000000
)
