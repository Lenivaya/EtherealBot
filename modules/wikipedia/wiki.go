package wikipedia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}

type Result struct {
	Name, Description, URL string
}

func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	array := []interface{}{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}

	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Results = append(sr.Results, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	return nil
}

func WikipediaAPI(request string) (answer []string, err error) {
	url := "https://en.wikipedia.org/w/api.php?action=opensearch&search=" + request + "&limit=3&origin=*&format=json"

	pages := make([]string, 3)

	//Sending request
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		contents, _ := ioutil.ReadAll(resp.Body)

		//Unmarshal answer and set it to SearchResults struct
		results := &SearchResults{}
		if err = json.Unmarshal([]byte(contents), results); err != nil {
			pages[0] = "Something going wrong, try to change your question"
		}

		//Check if struct is not empty
		if !results.ready {
			pages[0] = "Something going wrong, try to change your question"
		}

		//Loop through Results struct and assigning data to s slice
		for i := range results.Results {
			pages[i] = results.Results[i].URL
		}
	}
	return pages, nil
}
