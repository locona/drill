package drill

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/k0kubun/pp"
)

var (
	queryEndpoint = "%s/query.json"
)

type Drill struct {
	URL string
}

type queryRequest struct {
	QueryType string `json:"queryType"`
	Query     string `json:"query"`
}

type queryResposne struct {
	Columns []string                 `json:"columns"`
	Rows    []map[string]interface{} `json:"rows"`
}

type Config struct {
	URL  string
	Port int
}

func New(url string) *Drill {
	return &Drill{
		URL: url,
	}
}

func (this *Drill) Query(query string) (*queryResposne, error) {
	body := queryRequest{QueryType: "SQL", Query: query}
	endpoint := fmt.Sprintf(queryEndpoint, this.URL)

	b, err := json.Marshal(&body)
	if err != nil {
		os.Exit(91)
	}

	resp, err := Do("POST", endpoint, bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result := &queryResposne{}
	err = json.Unmarshal(b, result)
	if err != nil {
		// pp.Println(err)
		return nil, err
	}
	pp.Println(result.Columns)
	pp.Println(result.Rows)

	return result, nil
}

func Do(method, endpoint string, body io.Reader) (*http.Response, error) {
	parsedUrl, _ := url.Parse(endpoint)
	pp.Println(method, parsedUrl.String(), body)
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
		pp.Println(resp.StatusCode)
		return nil, errors.New("Query Error")
	}

}
