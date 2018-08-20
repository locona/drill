package drill

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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

type queryResponse struct {
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

func (this *Drill) Query(query string) (*queryResponse, error) {
	body := queryRequest{QueryType: "SQL", Query: query}
	endpoint := fmt.Sprintf(queryEndpoint, this.URL)

	b, err := json.Marshal(&body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := Do("POST", endpoint, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := &queryResponse{}
	err = json.Unmarshal(b, result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pp.Println(result.Columns)
	pp.Println(result.Rows)

	return result, nil
}
