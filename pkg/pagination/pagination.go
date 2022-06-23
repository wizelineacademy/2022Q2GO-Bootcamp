package pagination

import (
	"net/http"
	"strconv"
)

var (

	// DefaultPageSize specifies the default page size
	DefaultPageSize = 100
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 1000
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// Limit specifies the query parameter name for page size
	Limit = "limit"
)

// Pagination represents a paginated list of data items.
type Pagination struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

/*
* @Description
* New creates a new Pagination instance.
* @Params
* page is 1-based and refers to the current page index/number.
* limit refers to the number of items on each page.
* total specifies the total number of data items.
* @Rules
* If total is less than 0, it means total is unknown.
 */
func New(page, limit, total int) *Pagination {
	if limit <= 0 {
		limit = DefaultPageSize
	}
	if limit > MaxPageSize {
		limit = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + limit - 1) / limit
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &Pagination{
		Page:       page,
		PerPage:    limit,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

/*
* @Description
* creates a Pagitantion object using the query parameters found in the given HTTP request.
* @Params
* count stands for the total number of items. Use -1 if this is unknown.
 */
func NewFromRequest(req *http.Request, count int) *Pagination {
	page := parseInt(req.URL.Query().Get(PageVar), 1)
	perPage := parseInt(req.URL.Query().Get(Limit), DefaultPageSize)
	return New(page, perPage, count)
}

/*
* @Description
* parses a string into an integer. If parsing is failed, defaultValue will be returned.
 */
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

/*
* Offset
* @Description
* returns the OFFSET value that can be used in a SQL statement.
 */
func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PerPage
}

/*
* Limit
* @Description
* returns the LIMIT value that can be used in a SQL statement.
 */
func (p *Pagination) Limit() int {
	return p.PerPage
}
