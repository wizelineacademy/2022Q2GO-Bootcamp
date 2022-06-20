package pagination

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

/*
* BuildLinkHeader
* @Description
* returns an HTTP header containing the links about the pagination.
 */
func (p *Pagination) BuildLinkHeader(baseURL string, defaultPerPage int) string {
	links := p.BuildLinks(baseURL, defaultPerPage)
	header := ""
	if links[0] != "" {
		header += fmt.Sprintf("<%v>; rel=\"first\", ", links[0])
		header += fmt.Sprintf("<%v>; rel=\"prev\"", links[1])
	}
	if links[2] != "" {
		if header != "" {
			header += ", "
		}
		header += fmt.Sprintf("<%v>; rel=\"next\"", links[2])
		if links[3] != "" {
			header += fmt.Sprintf(", <%v>; rel=\"last\"", links[3])
		}
	}
	return header
}

/*
* BuildLinks
* @Description
* returns the first, prev, next, and last links corresponding to the pagination.
* @ Notes
* A link could be an empty string if it is not needed.
* For example, if the pagination is at the first page, then both first and prev links
* will be empty.
 */
func (p *Pagination) BuildLinks(baseURL string, defaultPerPage int) [4]string {
	var links [4]string
	pageCount := p.PageCount
	page := p.Page
	if pageCount >= 0 && page > pageCount {
		page = pageCount
	}
	if strings.Contains(baseURL, "?") {
		baseURL += "&"
	} else {
		baseURL += "?"
	}
	if page > 1 {
		links[0] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, 1)
		links[1] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page-1)
	}
	if pageCount >= 0 && page < pageCount {
		links[2] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page+1)
		links[3] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, pageCount)
	} else if pageCount < 0 {
		links[2] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page+1)
	}
	if perPage := p.PerPage; perPage != defaultPerPage {
		for i := 0; i < 4; i++ {
			if links[i] != "" {
				links[i] += fmt.Sprintf("&%v=%v", Limit, perPage)
			}
		}
	}

	return links
}
