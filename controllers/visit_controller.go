package controllers

import (
	"RestApiProject/models"
	"RestApiProject/services"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func getAllOfAUsersVisits(req *restful.Request, resp *restful.Response) {
	userID := req.PathParameter(UserID)
	paginationParameters := parseQueryOptions(req)
	visits, err := services.GetAllOfAUsersVisits(userID, paginationParameters)

	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
	} else {
		resp.WriteHeaderAndJson(http.StatusOK, visits, restful.MIME_JSON)
	}
}

func getAllOfAUsersStateVisits(req *restful.Request, resp *restful.Response) {
	userID := req.PathParameter(UserID)
	visits, err := services.GetAllOfAUsersStateVisits(userID)

	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
	} else {
		resp.WriteHeaderAndJson(http.StatusOK, visits, restful.MIME_JSON)
	}
}

func postVisit(req *restful.Request, resp *restful.Response) {
	var visit models.Visit
	if err := req.ReadEntity(&visit); err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}

	userID := req.PathParameter(UserID)
	newVisit, err := services.PostVisit(userID, visit)

	if err != nil {
		switch err.(type) {
		case services.CityNotExistInStateError:
			resp.WriteError(http.StatusBadRequest, err)
		case services.VisitValidationError:
			resp.WriteError(http.StatusBadRequest, err)
		default:
			resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
		}
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK, newVisit, restful.MIME_JSON)
}

func deleteVisit(req *restful.Request, resp *restful.Response) {
	userID := req.PathParameter(UserID)
	visitID := req.PathParameter(VisitID)

	err := services.DeleteVisit(userID, visitID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
	} else {
		resp.WriteHeader(http.StatusNoContent)
	}
}
