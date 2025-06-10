package handler

import (
	"net/http"
	"user-service/internal/adapter/handler/request"
	"user-service/internal/adapter/handler/response"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type UserHandlerInterface interface {
	SignIn(c echo.Context) error
}

type userHandler struct {
	userService service.UserServiceInterface
}

var err error

func NewUserHandler(e *echo.Echo, userService service.UserServiceInterface) UserHandlerInterface {
	userHandler := &userHandler{
		userService: userService,
	}

	e.Use(middleware.Recover())
	e.POST("/signin", userHandler.SignIn)

	return userHandler
}

func (handler *userHandler) SignIn(c echo.Context) error {
	var (
		request    = request.SignInRequests{}
		resp       = response.DefaultResponse{}
		respSignIn = response.SignInResponse{}
		ctx        = c.Request().Context()
	)

	if err = c.Bind(&request); err != nil {
		log.Errorf("[UserHandler-1] SignIn: %v", err)
		resp.Message = err.Error()
		resp.Data = nil
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	if err = c.Validate(&request); err != nil {
		log.Errorf("[UserHandler-2] SignIn: %v", err)
		resp.Message = err.Error()
		resp.Data = nil
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	reqEntity := entity.UserEntity{
		Email:    request.Email,
		Password: request.Password,
	}

	user, token, err := handler.userService.SignIn(ctx, reqEntity)
	if err != nil {
		if err.Error() == "404" {
			log.Errorf("[UserHandler-3] SignIn: %s", "User not found")
			resp.Message = "User not found"
			resp.Data = nil
			return c.JSON(http.StatusNotFound, resp)
		}
		log.Errorf("[UserHandler-3] SignIn: %v", err)
		resp.Message = err.Error()
		resp.Data = nil
		return c.JSON(http.StatusInternalServerError, resp)
	}

	respSignIn.ID = user.ID
	respSignIn.AccessToken = token
	respSignIn.Name = user.Name
	respSignIn.Email = user.Email
	respSignIn.Phone = user.Phone
	respSignIn.Lat = user.Lat
	respSignIn.Lng = user.Lng
	respSignIn.Role = user.RoleName

	resp.Message = "Success"
	resp.Data = respSignIn

	return c.JSON(http.StatusOK, resp)

}
