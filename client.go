package drill

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	GET  = "GET"
	POST = "POST"
)

func Do(method, endpoint string, body io.Reader) (*http.Response, error) {
	parsedUrl, _ := url.Parse(endpoint)
	req, err := http.NewRequest(method, parsedUrl.String(), body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	// pp.Println(resp, err)
	// defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return resp, nil
	default:
		return nil, errors.New("Query Error")
	}

}
