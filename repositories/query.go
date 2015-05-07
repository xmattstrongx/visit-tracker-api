package repositories

import (
	"RestApiProject/models"

	"gopkg.in/mgutz/dat.v1"
)

func ApplyPaging(query *dat.SelectBuilder, paginationParameters models.PaginationParameters) {
	if paginationParameters.HasLimit() {
		query.Limit(*paginationParameters.GetLimit())
	}
	if paginationParameters.HasOffset() {
		query.Offset(*paginationParameters.GetOffset())
	}

	if paginationParameters.HasPage() && paginationParameters.HasPerPage() {
		query.Paginate(*paginationParameters.GetPage(), *paginationParameters.GetPerPage())
	}
}
