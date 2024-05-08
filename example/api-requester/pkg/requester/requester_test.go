package requester

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var expectedResponse = Response{
	UserId: 1,
	Id:     1,
	Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
	Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
}

func TestGetJsonPlaceHolder(t *testing.T) {
	requester := NewRequester(
		&http.Client{
			Timeout: time.Second * 30,
		},
	)

	res, err := requester.GetJsonPlaceHolder(context.Background())
	require.NoError(t, err)

	assert.Greater(t, len(res), 0)
	assert.Equal(t, expectedResponse, res[0])
}
