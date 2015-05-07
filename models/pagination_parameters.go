package models

import "strconv"

type PaginationParameters struct {
	limit   *uint64
	offset  *uint64
	page    *uint64
	perPage *uint64
}

func NewPaginationParameters(limit, offset, page, perPage string) PaginationParameters {
	newPaginationParameters := PaginationParameters{}

	limitUInt64, err := strconv.ParseUint(limit, 10, 64)
	if err == nil && limitUInt64 >= 0 {
		newPaginationParameters.limit = &limitUInt64
	}

	offsetUInt64, err := strconv.ParseUint(offset, 10, 64)
	if err == nil && offsetUInt64 >= 0 {
		newPaginationParameters.offset = &offsetUInt64
	}

	pageUInt64, err := strconv.ParseUint(page, 10, 64)
	if err == nil && pageUInt64 >= 0 {
		newPaginationParameters.page = &pageUInt64
	}

	perPageUInt64, err := strconv.ParseUint(perPage, 10, 64)
	if err == nil && perPageUInt64 >= 0 {
		newPaginationParameters.perPage = &perPageUInt64
	}
	return newPaginationParameters
}

func (p *PaginationParameters) GetLimit() *uint64 {
	return p.limit
}

func (p *PaginationParameters) GetOffset() *uint64 {
	return p.offset
}

func (p *PaginationParameters) GetPage() *uint64 {
	return p.page
}

func (p *PaginationParameters) GetPerPage() *uint64 {
	return p.perPage
}

func (p *PaginationParameters) HasLimit() bool {
	return p.limit != nil
}

func (p *PaginationParameters) HasOffset() bool {
	return p.offset != nil
}

func (p *PaginationParameters) HasPage() bool {
	return p.page != nil && *p.page >= 1
}

func (p *PaginationParameters) HasPerPage() bool {
	return p.perPage != nil && *p.perPage >= 1
}
