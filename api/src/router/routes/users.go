package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                    "/usuarios",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/usuarios",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequired: false,
	},
}
