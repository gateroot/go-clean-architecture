package echo

import (
	"awesomeProject1/model"
	"awesomeProject1/user"
	"awesomeProject1/user/delivery"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)


type userDelivery struct {
	uc user.UserUsecase
}

func NewUserDelivery(usecase user.UserUsecase) delivery.UserDelivery {
	return userDelivery{usecase}
}

func (ud userDelivery) Start(addr string) error {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/users/:id", ud.get())
	e.POST("/users", ud.post())
	e.PUT("/users/:id", ud.edit())
	e.DELETE("/users/:id", ud.delete())

	return e.Start(addr)
}

func (ud userDelivery) get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "cannot parse user id.")
		}
		u, err := ud.uc.Get(userID)
		if err != nil {
			return c.String(http.StatusInternalServerError, "internal server error.")
		}
		message := fmt.Sprintf("get user successfully. id: %d, name: %s", u.Id, u.Name)
		return c.String(http.StatusOK, message)
	}
}

func (ud userDelivery) post() echo.HandlerFunc {
	return func(c echo.Context) error {
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		name, ok := m["name"].(string)
		if !ok {
			return c.String(http.StatusBadRequest, "name is required.")
		}
		user := model.User{Name: name}
		u, err := ud.uc.Add(&user)
		if err != nil {
			return c.String(http.StatusInternalServerError, "user add failed.")
		}
		message := fmt.Sprintf("add user successfully. id: %d, name: %s", u.Id, u.Name)
		return c.String(http.StatusOK, message)
	}
}

func (ud userDelivery) edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "cannot parse user id.")
		}
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		name, ok := m["name"].(string)
		if !ok {
			return c.String(http.StatusBadRequest, "name is required.")
		}
		editUser := model.User{userID, name}
		if err := ud.uc.Edit(&editUser); err != nil {
			return err
		}
		message := fmt.Sprintf("edit user successfully. id: %d, name: %s", editUser.Id, editUser.Name)
		return c.String(http.StatusOK, message)
	}
}

func (ud userDelivery) delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "cannot parse user id.")
		}
		if err := ud.uc.Delete(userID); err != nil {
			return err
		}
		message := fmt.Sprintf("delete user successfully. id: %d", userID)
		return c.String(http.StatusOK, message)
	}
}
