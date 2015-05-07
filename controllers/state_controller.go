package controllers

import (
	"RestApiProject/services"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func getStates(req *restful.Request, resp *restful.Response) {
	states, err := services.GetAllStates()
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
		return
	}

	resp.WriteHeaderAndJson(http.StatusOK, states, restful.MIME_JSON)
}

func getSpecificState(req *restful.Request, resp *restful.Response) {
	stateAbbreviation := req.PathParameter(StateAbbreviation)
	states, err := services.GetSpecificState(stateAbbreviation)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
		return
	}

	resp.WriteHeaderAndJson(http.StatusOK, states, restful.MIME_JSON)
}
