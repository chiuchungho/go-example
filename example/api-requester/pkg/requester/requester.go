package requester

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const path = "https://jsonplaceholder.typicode.com/posts"

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Requester struct {
	client Doer
}

func NewRequester(doer Doer) Requester {
	return Requester{
		client: doer,
	}
}

func (r Requester) GetJsonPlaceHolder(ctx context.Context) ([]Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to")
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to")
	}
	defer resp.Body.Close()

	var result []Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to")
	}

	return result, nil
}

// Sample response from https://jsonplaceholder.typicode.com/posts
// [
//  {
//     "userId": 1,
//     "id": 1,
//     "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
//     "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
//   }
// ]
type Response struct {
	UserId int64  `json:"userId"`
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
