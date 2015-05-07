package controllers

import (
	"RestApiProject/models"
	"RestApiProject/services"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func getCitiesForSpecificState(req *restful.Request, resp *restful.Response) {
	stateAbbreviation := req.PathParameter(StateAbbreviation)
	paginationParameters := parseQueryOptions(req)

	cities, err := services.GetCitiesForSpecificState(stateAbbreviation, paginationParameters)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
	} else {
		resp.WriteHeaderAndJson(http.StatusOK, cities, restful.MIME_JSON)
	}
}

func GetDetailsForSpecificCity(req *restful.Request, resp *restful.Response) {
	stateAbbreviation := req.PathParameter(StateAbbreviation)
	cityName := req.PathParameter(CityName)
	cities, err := services.GetDetailsForSpecificCity(stateAbbreviation, cityName)

	if err != nil {
		resp.WriteError(http.StatusInternalServerError, newDefaultInternalServerError(err))
	} else {
		resp.WriteHeaderAndJson(http.StatusOK, cities, restful.MIME_JSON)
	}
}

func parseQueryOptions(req *restful.Request) models.PaginationParameters {
	limit := req.QueryParameter(Limit)
	offset := req.QueryParameter(Offset)
	page := req.QueryParameter(Page)
	perPage := req.QueryParameter(PerPage)

	return models.NewPaginationParameters(limit, offset, page, perPage)
}
