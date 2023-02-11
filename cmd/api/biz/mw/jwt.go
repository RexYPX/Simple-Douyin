// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package mw

import (
	"Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/kitex_gen/user"
	"Simple-Douyin/pkg/constants"
	"Simple-Douyin/pkg/errno"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"Simple-Douyin/cmd/api/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(constants.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, form: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   constants.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userid, _ := claims[constants.IdentityKey].(json.Number).Int64()
			return &api.User{
				ID: userid,
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req api.UserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			u, err := rpc.LoginUser(context.Background(), &user.UserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			c.Set(constants.IdentityId, u.UserId)

			return u.UserId, err
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			uid, exist := c.Get(constants.IdentityId)
			if !exist {
				c.JSON(http.StatusOK, utils.H{
					"status_code": errno.Success.ErrCode,
					"status_msg":  "There is no token in the context",
					"user_id":     0,
					"token":       "InvalidToken",
				})
				return
			}
			c.JSON(http.StatusOK, utils.H{
				"status_code": errno.Success.ErrCode,
				"status_msg":  "Login success!",
				"user_id":     uid,
				"token":       token,
				// "expire":      expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": errno.AuthorizationFailedErr.ErrCode,
				"status_msg":  "",
				"user_id":     0,
				"token":       "UnauthorizedToken",
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
		ParseOptions: []jwtv4.ParserOption{jwtv4.WithJSONNumber()},
	})
}
