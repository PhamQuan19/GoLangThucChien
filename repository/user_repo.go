package repository

import (
	"context"
	"my-app/model"
)

type UserRepo interface {
	SaveUser(contex context.Context,user model.User)(model.User, error)
}