package api

import (
	"gin_docs_server/api/image_api"
	"gin_docs_server/api/user_api"
)

type Api struct {
	UserApi  user_api.UserApi
	ImageApi image_api.ImageApi
}

var App = new(Api)
