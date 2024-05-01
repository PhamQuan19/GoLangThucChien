package repository

import (
	"context"
	"my-app/model"
	"my-app/model/req"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req.ReqSigIn)(model.User, error)
	SaveUser(context context.Context,user model.User)(model.User, error)
}