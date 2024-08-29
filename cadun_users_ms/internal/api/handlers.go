package api

import (
	"net/http"

	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/api/dtos"
	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/views"
	"github.com/labstack/echo/v4"
)

//Handlers for the API
//There is a handdler for each route that was created, each handler calls the respective view that is defined
//If there is an error with the request each handler throws an exception
//The handlers that recevies a parametes through the URL, that parameter is recived by the echo context and used in the query
//All the other handlers use the dtos parameters since the parameters are given by a JSON format.
//param es el parametro tomado desde la URL /users/123

func (a *API) Create_User(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Create_User{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Create_user(ctx, parametros.Names, parametros.LastNames, parametros.Alias, parametros.Password, parametros.EMail, parametros.PhoneNumber, parametros.Country)
	if err != nil {
		if err == views.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (a *API) Get_userid_Byemail(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_userid_Byemail{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := a.view.Get_userid_Byemail(ctx, parametros.EMail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, id)
}

func (a *API) Read_userByid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Read_userByid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := a.view.Read_userByid(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (a *API) Update_userByid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Update_userByid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Update_userByid(ctx, parametros.Names, parametros.LastNames, parametros.Alias, parametros.Password, parametros.EMail, parametros.PhoneNumber, parametros.Country, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (a *API) Delete_userByid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Delete_userByid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.view.Delete_userByid(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (a *API) Get_requeststatus_Byid(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_requeststatus_Byid{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := a.view.Get_requeststatus_Byid(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (a *API) Get_requeststatus_ByUser(c echo.Context) error {
	ctx := c.Request().Context()
	parametros := dtos.Get_requeststatus_ByUser{}
	err := c.Bind(&parametros)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := a.view.Get_requeststatus_ByUser(ctx, parametros.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}
