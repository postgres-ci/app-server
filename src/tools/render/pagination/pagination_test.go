package pagination

import (
	"github.com/postgres-ci/http200ok"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/url"
	"testing"
)

func TestPaginationSimple(t *testing.T) {

	for i := 0; i <= 10; i++ {

		assert.Len(t, New(&http200ok.Context{}, int32(i), 10).Pages, 0)
	}
}

func TestPagination(t *testing.T) {

	{
		context := &http200ok.Context{
			Request: &http.Request{
				URL: &url.URL{
					RawQuery: "p=2",
				},
			},
		}

		pagination := New(context, 15, 10)

		if assert.Len(t, pagination.Pages, 2) {

			if assert.Nil(t, pagination.Previous) && assert.Nil(t, pagination.Next) {

				assert.True(t, pagination.Pages[1].IsCurrent)
			}
		}
	}

	{
		context := &http200ok.Context{
			Request: &http.Request{
				URL: &url.URL{
					RawQuery: "p=9",
				},
			},
		}

		pagination := New(context, 100, 10)

		if assert.Len(t, pagination.Pages, 10) {

			if assert.Nil(t, pagination.Previous) && assert.Nil(t, pagination.Next) {

				assert.True(t, pagination.Pages[8].IsCurrent)
			}
		}
	}

	{
		context := &http200ok.Context{
			Request: &http.Request{
				URL: &url.URL{
					RawQuery: "p=15",
				},
			},
		}

		pagination := New(context, 256, 10)

		if assert.Len(t, pagination.Pages, 11) {

			if assert.NotNil(t, pagination.Previous) && assert.NotNil(t, pagination.Next) {

				assert.Equal(t, int32(14), pagination.Previous.Num)
				assert.Equal(t, int32(16), pagination.Next.Num)
			}

			assert.Equal(t, int32(10), pagination.Pages[0].Num)
			assert.Equal(t, int32(20), pagination.Pages[10].Num)

		}
	}

	{
		context := &http200ok.Context{
			Request: &http.Request{
				URL: &url.URL{
					RawQuery: "p=25",
				},
			},
		}

		pagination := New(context, 256, 10)

		if assert.Len(t, pagination.Pages, 11) {

			if assert.NotNil(t, pagination.Previous) && assert.Nil(t, pagination.Next) {

				assert.Equal(t, int32(16), pagination.Pages[0].Num)
				assert.Equal(t, int32(26), pagination.Pages[10].Num)
			}
		}
	}
}

func TestNumPages(t *testing.T) {

	{
		assets := map[int32]int32{
			0:  0,
			1:  1,
			10: 1,
			11: 2,
			15: 2,
			20: 2,
			21: 3,
			29: 3,
			30: 3,
			31: 4,
		}

		for total, num := range assets {

			assert.Equal(t, num, numPages(total, 10))
		}
	}

	{
		assets := map[int32]int32{
			0:  0,
			1:  1,
			10: 1,
			11: 1,
			20: 2,
			21: 2,
			29: 2,
			30: 2,
			31: 3,
		}

		for total, num := range assets {

			assert.Equal(t, num, numPages(total, 15))
		}
	}
}
