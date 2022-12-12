package database

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"io"
	"log"
)

var es *elasticsearch.Client

func Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.0.29:9200",
		},
	}
	es, _ = elasticsearch.NewClient(cfg)
	log.Println(es.Info())
}

func SearchPaper(query bytes.Buffer) map[string]interface{} {
	var res map[string]interface{}
	resp, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("papers"),
		es.Search.WithBody(&query),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return res
}

func SearchAuthor(query bytes.Buffer) map[string]interface{} {
	var res map[string]interface{}
	resp, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("authors"),
		es.Search.WithBody(&query),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return res
}

func SearchAuthorWithoutContext(query bytes.Buffer) map[string]interface{} {
	var res map[string]interface{}
	resp, err := es.Search(
		es.Search.WithIndex("authors"),
		es.Search.WithBody(&query),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return res
}

func SearchAutoComplete(query bytes.Buffer) (map[string]interface{}, error) {
	var res map[string]interface{}
	log.Println(query.String())
	resp, err := es.Search(
		//es.Search.WithContext(context.Background()),
		es.Search.WithIndex("auto-complete"),
		es.Search.WithBody(&query),
		es.Search.WithPretty(),
	)
	log.Printf(resp.String())
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(resp.Status())
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing the response body: %s\n", err))
	}
	return res, nil
}

func MgetPaper(query bytes.Buffer) map[string]interface{} {
	var res map[string]interface{}
	resp, err := es.Mget(
		io.Reader(&query),
		es.Mget.WithContext(context.Background()),
		es.Mget.WithIndex("papers"),
		es.Mget.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return res
}

func MgetScholer(query bytes.Buffer) map[string]interface{} {
	var res map[string]interface{}
	resp, err := es.Mget(
		io.Reader(&query),
		es.Mget.WithContext(context.Background()),
		es.Mget.WithIndex("authors"),
		es.Mget.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return res
}

func UpdatePaper(query bytes.Buffer, id string) map[string]interface{} {
	var res map[string]interface{}
	resp, err := es.Update(
		"papers",
		id,
		io.Reader(&query),
		es.Update.WithContext(context.Background()),
		es.Update.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return res
}

func UpdateAuthor(query bytes.Buffer, id string) map[string]interface{} {
	var res map[string]interface{}
	resp, err := es.Update(
		"authors",
		id,
		io.Reader(&query),
		es.Update.WithContext(context.Background()),
		es.Update.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error parsing the response body: %s\n", err)
	}
	return res
}
