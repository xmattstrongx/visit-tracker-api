package controllers

import (
	"RestApiProject/logger"
	"RestApiProject/models"
	"fmt"

	restful "github.com/emicklei/go-restful"
)

func RegisterVisitorApiRoutes() {
	ws := new(restful.WebService)
	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Filter(basicAuthenticate)

	ws.Route(ws.GET("/states").To(getStates).
		Doc("Lists all states").
		Writes([]models.State{}))

	ws.Route(ws.GET(fmt.Sprintf("/states/{%s}", StateAbbreviation)).To(getSpecificState).
		Doc("Lists details of a specific state").
		Param(ws.PathParameter(StateAbbreviation, "abbreviation of the state").DataType(DataTypeString)).
		Writes(models.State{}))

	ws.Route(ws.GET(fmt.Sprintf("/states/{%s}/cities", StateAbbreviation)).To(getCitiesForSpecificState).
		Doc("Lists all cities of a specific state").
		Param(ws.PathParameter(StateAbbreviation, "abbreviation of the state").DataType(DataTypeString)).
		Param(ws.QueryParameter(Limit, "optional, paired with offset, limits the number of results returned").DataType(DataTypeString)).
		Param(ws.QueryParameter(Offset, "optional, paired with limit, specifies the number of entries to be skipped before the desired page").DataType(DataTypeString)).
		Param(ws.QueryParameter(Page, "optional, paired with per_page, specifies the desired page to be returned").DataType(DataTypeString)).
		Param(ws.QueryParameter(PerPage, "optional, paired with page, specifies the number of entries on each page").DataType(DataTypeString)).
		Writes([]models.City{}))

	ws.Route(ws.GET(fmt.Sprintf("/states/{%s}/cities/{%s}", StateAbbreviation, CityName)).To(GetDetailsForSpecificCity).
		Doc("Lists all cities of a specific state").
		Param(ws.PathParameter(StateAbbreviation, "abbreviation of the state").DataType(DataTypeString)).
		Param(ws.PathParameter(CityName, "name of city").DataType(DataTypeString)).
		Writes(models.CityDetails{}))

	ws.Route(ws.GET("/users").To(getUsers).
		Doc("Lists all users").
		Writes([]models.State{}))

	ws.Route(ws.GET(fmt.Sprintf("/users/{%s}/visits", UserID)).To(getAllOfAUsersVisits).
		Doc("List all of a users visits").
		Param(ws.PathParameter(UserID, "The user's unique identifier").DataType(DataTypeString)).
		Param(ws.QueryParameter(Limit, "optional, paired with offset, limits the number of results returned").DataType(DataTypeString)).
		Param(ws.QueryParameter(Offset, "optional, paired with limit, specifies the number of entries to be skipped before the desired page").DataType(DataTypeString)).
		Param(ws.QueryParameter(Page, "optional, paired with per_page, specifies the desired page to be returned").DataType(DataTypeString)).
		Param(ws.QueryParameter(PerPage, "optional, paired with page, specifies the number of entries on each page").DataType(DataTypeString)).
		Writes([]models.Visit{}))

	ws.Route(ws.POST(fmt.Sprintf("/users/{%s}/visits", UserID)).To(postVisit).
		Doc("Post a new visit for a user").
		Param(ws.PathParameter(UserID, "The user's unique identifier").DataType(DataTypeString)).
		Writes(models.Visit{}))

	ws.Route(ws.DELETE(fmt.Sprintf("/users/{%s}/visits/{%s}", UserID, VisitID)).To(deleteVisit).
		Doc("Delete a visit for a user").
		Param(ws.PathParameter(UserID, "The user's unique identifier").DataType(DataTypeString)).
		Param(ws.PathParameter(VisitID, "The visit's unique identifier").DataType(DataTypeString)))

	ws.Route(ws.GET(fmt.Sprintf("/users/{%s}/visits/states", UserID)).To(getAllOfAUsersStateVisits).
		Doc("List all of a users visits to a specific state").
		Param(ws.PathParameter(UserID, "The user's unique identifier").DataType(DataTypeString)).
		Writes([]models.State{}))

	logger.Message("Registering routes")
	restful.Add(ws)
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	encoded := req.Request.Header.Get("Authorization")
	// usr/pwd = admin/admin
	// real code does some decoding
	// would be better to authorize with OAUTH2 and proper ISP
	if len(encoded) == 0 || "Basic YWRtaW46YWRtaW4=" != encoded {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, resp)
}
