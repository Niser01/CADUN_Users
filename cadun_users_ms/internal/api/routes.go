package api

import "github.com/labstack/echo/v4"

//routes for users
//the section that is ":id" is where the 1 id the user is placed to make the search
//the routes are divided by users and saved element, to give order to the requests

func (a *API) RegisterRoutes(e *echo.Echo) {
	users := e.Group("/user")
	users.POST("/create_user", a.Create_User)
	users.GET("/userid_by_email", a.Get_userid_Byemail)
	users.GET("/password_by_email", a.Get_password_Byemail)
	users.GET("/user_by_id", a.Read_userByid)
	users.PUT("/update_userbyid", a.Update_userByid)
	users.DELETE("/delete_userbyid", a.Delete_userByid)

	request := e.Group("/request")
	request.POST("/create_request", a.Create_request)
	request.POST("/create_tipo_de_request", a.Create_requesttype)
	request.GET("/request_status_byid", a.Get_requeststatus_Byid)
	request.GET("/request_status_byUser", a.Get_requeststatus_ByUser)
	request.PUT("/request_status_byrequestid", a.Update_requeststatus_Byid)
	request.DELETE("/delete_requests_byuserid", a.Delete_requests_ByUserid)

	cotizacion := e.Group("/cotizacion")
	cotizacion.POST("/create_cotizacion", a.Create_cotizacion)
	cotizacion.DELETE("/delete_cotizacion_byuserid", a.Delete_cotizacion_ByUserid)
	cotizacion.GET("/cotizacion_byrequest", a.Get_cotizacion_ByRequest)

	password := e.Group("/password")
	password.POST("/RevisarPassword", a.RevisarPassword)

}
