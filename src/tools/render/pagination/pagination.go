package pagination

import (
	"github.com/postgres-ci/http200ok"

	"fmt"
	"math"
	"net/url"
	"strconv"
)

const window int32 = 10

func New(c *http200ok.Context, total, perPage int32) *pagination {

	num := numPages(total, perPage)

	if num <= 1 {

		return &pagination{}
	}

	var currentPage int32 = 1

	if value, err := strconv.ParseInt(c.Request.URL.Query().Get("p"), 10, 32); err == nil {

		currentPage = int32(value)
	}

	pagination := &pagination{
		Pages: make([]page, 0, perPage),
	}

	var (
		pages []int32
		query = c.Request.URL.Query()
	)

	switch true {

	case num <= window:
		pages = pagesList(1, num)
	case currentPage > (num - window):
		pages = pagesList(num-window, num)
	default:
		pages = pagesList(currentPage-(window/2), currentPage+(window/2))

		query.Set("p", fmt.Sprint(currentPage-1))
		pagination.Previous = &page{
			Num: currentPage - 1,
			URL: (&url.URL{
				Path:     c.Request.URL.Path,
				RawQuery: query.Encode(),
			}).String(),
		}

		query.Set("p", fmt.Sprint(currentPage+1))
		pagination.Next = &page{
			Num: currentPage + 1,
			URL: (&url.URL{
				Path:     c.Request.URL.Path,
				RawQuery: query.Encode(),
			}).String(),
		}
	}

	for _, i := range pages {

		query.Set("p", fmt.Sprint(i))

		pagination.Pages = append(pagination.Pages, page{
			Num: i,
			URL: (&url.URL{
				Path:     c.Request.URL.Path,
				RawQuery: query.Encode(),
			}).String(),
			IsCurrent: i == currentPage,
		})
	}

	return pagination
}

type pagination struct {
	Pages          []page
	Previous, Next *page
}

type page struct {
	IsCurrent bool
	URL       string
	Num       int32
}

func numPages(total, perPage int32) int32 {

	return int32(math.Ceil(float64(total) / float64(perPage)))
}

func pagesList(start, end int32) []int32 {

	var (
		i     int32
		pages = make([]int32, 0, window)
	)

	for i = start; i <= end; i++ {

		pages = append(pages, i)
	}

	return pages
}
