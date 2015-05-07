package controllers

import (
	"RestApiProject/services"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func getUsers(req *restful.Request, resp *restful.Response) {
	users, err := services.GetAllUsers()
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
		return
	}

	resp.WriteHeaderAndJson(http.StatusOK, users, restful.MIME_JSON)
}
