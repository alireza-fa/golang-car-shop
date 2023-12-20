package dto

type Sort struct {
	ColId string `json:"colId"`
	Sort  string `json:"sort"`
}

type Filter struct {
	// contains notContains equals notEqual startWith lessThan lessThanOrEqual greaterThan greaterThanOrEqual inRange endWith
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`
	// text number
	FilterType string `json:"filterType"`
}

type DynamicFilter struct {
	Sort   *[]Sort           `json:"sort"`
	Filter map[string]Filter `json:"filter"`
}

type PagedList[T any] struct {
	PageNumber      int   `json:"pageNumber"`
	TotalRows       int64 `json:"totalRows"`
	TotalPages      int   `json:"totalPages"`
	HasPreviousPage bool  `json:"hasPreviousPage"`
	HasNextPage     bool  `json:"hasNextPage"`
	Items           *[]T
}

type PaginationInout struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
}

type PaginationInputWithFilter struct {
	PaginationInout
	DynamicFilter
}

func (p *PaginationInputWithFilter) GetOffset() int {
	// 2 , 10 => 11-20
	return (p.GetPageNumber() - 1) * p.GetPageSize()
}

func (p *PaginationInputWithFilter) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *PaginationInputWithFilter) GetPageNumber() int {
	if p.PageNumber == 0 {
		p.PageSize = 1
	}
	return p.PageNumber
}
