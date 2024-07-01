package routes

import (
	"api/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Path:         "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequired: false,
	},
	{
		Path:         "/user/{id}",
		Method:       http.MethodGet,
		Function:     controllers.SearchUser,
		AuthRequired: false,
	},
	{
		Path:         "/user",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		Path:         "/user/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		Path:         "/user/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	},
}
