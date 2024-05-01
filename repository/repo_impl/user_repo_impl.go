package repoimpl

import (
	"context"
	// "log"
	"my-app/db"
	"my-app/error_message"
	"my-app/model"
	"my-app/repository"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo{
	return &UserRepoImpl{
		sql:sql,
	}
}

func( u UserRepoImpl) SaveUser(context context.Context,user model.User)(model.User, error){
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt=time.Now()
	user.UpdateAt=time.Now()

	_, err := u.sql.Db.NamedExecContext(context,statement,user)

	if err !=nil{
		// return user, err
		log.Error(err.Error())
		if err, ok:=err.(*pq.Error);ok{
			if err.Code.Name() == "unique_violation" {
				return user,error_message.UserConflict
			}
		}
		return user, error_message.SignUpFail
	}

	return user, nil
}