package docgen

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Name string `json:"name"`
	Fields `json:"fields"`
	CreatedTime string `json:"createTime"`
	UpdatedTime string `json:"updateTime"`
}

type Fields struct {
	Text `json:"text"`
}

type Text struct {
	StringValue string `json:"stringValue"`
}

func GetDataByID(id string) (interface{}, error) {
	var url = "https://firestore.googleapis.com/v1/projects/doc-station/databases/(default)/documents/docs/"
	resp, err := http.Get(url + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data Item
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return data.Fields.Text.StringValue, nil
}
