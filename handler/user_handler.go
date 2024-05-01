package handler

import (
	// "hash"
	"my-app/model"
	req "my-app/model/req"
	"my-app/repository"
	"my-app/sercurity"
	"net/http"

	// "os/user"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}


func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSigUp{}
	//neu co loi
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Respone{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

	}

	validate := validator.New()
	// err := validate.Struct(req)
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Respone{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//ma hoa mat khau
	hash := sercurity.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Respone{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}
	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Respone{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Password=""
	return c.JSON(http.StatusOK, model.Respone{
		StatusCode: http.StatusOK,
		Message:    "Đăng kí thành công",
		Data:       user,
	})
}

//SignIn
func (u *UserHandler) HandleSignIn(c echo.Context) error {
	req:=req.ReqSigIn{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Respone{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

	}


	//Validate
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Respone{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user,err:= u.UserRepo.CheckLogin(c.Request().Context(),req)
	if err !=nil{
		return c.JSON(http.StatusUnauthorized, model.Respone{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//Check mat khẩu
	isTheSame:= sercurity.ComparePasswords(user.Password,[]byte(req.Password))
	if !isTheSame{
		return c.JSON(http.StatusUnauthorized, model.Respone{
			StatusCode: http.StatusUnauthorized,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK,model.Respone{
		StatusCode: http.StatusOK,
		Message:    "Đăng nhập thành công!",
		Data:       user,
	})
}