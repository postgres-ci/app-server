package limit

import (
	"github.com/postgres-ci/http200ok"
	"github.com/stretchr/testify/assert"

	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestOffset(t *testing.T) {

	{
		assets := map[int32]int32{
			-1: 0,
			0:  0,
			1:  0,
			2:  10,
			3:  20,
		}

		for currentPage, offset := range assets {

			context := &http200ok.Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: fmt.Sprintf("p=%d", currentPage),
					},
				},
			}

			assert.Equal(t, offset, Offset(context, 10))
		}
	}
	{
		assets := map[int32]int32{
			-1: 0,
			0:  0,
			1:  0,
			2:  13,
			3:  26,
		}

		for currentPage, offset := range assets {

			context := &http200ok.Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: fmt.Sprintf("p=%d", currentPage),
					},
				},
			}

			assert.Equal(t, offset, Offset(context, 13))
		}
	}
	{
		assets := map[int32]int32{
			-1: 0,
			0:  0,
			1:  0,
			2:  15,
			3:  30,
		}

		for currentPage, offset := range assets {

			context := &http200ok.Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: fmt.Sprintf("p=%d", currentPage),
					},
				},
			}

			assert.Equal(t, offset, Offset(context, 15))
		}
	}
}
