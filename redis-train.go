package main

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func init() {
	setupRedis()
}

type classifications struct {
	Sentiment string `json:"sentiment"`
	Category  string `json:"category"`
}

//RedisTrain trains a document and adds it to a redis serve
func RedisTrain(dataSet string, jsonData string) {
	c := redisPool.Get()
	defer c.Close()

	totalDocuments := hget(c, "TotalDocuments", "count")
	totalWords := hget(c, "TotalWords", "count")

	y := parse(jsonData)

	w := countWords(dataSet)
	for _, count := range w {
		totalWords += count
	}
	totalDocuments++

	train(y.Category, dataSet, c, w)
	train(y.Sentiment, dataSet, c, w)

	c.Send("HMSET", "TotalDocuments", "count", totalDocuments)
	c.Send("HMSET", "TotalWords", "count", totalWords)

	c.Flush()
}

func train(category string, dataSet string, c redis.Conn, w map[string]float64) error {

	wordCategoryTotal := hget(c, "word:"+category, "count")
	documentCategoryTotal := hget(c, "doc:"+category, "count")

	for word, count := range w {
		reply := hget(c, category+":"+word, "count")
		wordCount := count + reply
		c.Send("HMSET", category+":"+word, "count", wordCount)

		wordCategoryTotal += count
	}

	documentCategoryTotal++

	c.Send("HMSET", "word:"+category, "count", wordCategoryTotal)
	c.Send("HMSET", "doc:"+category, "count", documentCategoryTotal)

	return nil
}

func hget(c redis.Conn, key string, value string) float64 {
	reply, err := redis.Float64(c.Do("HGET", key, value))
	if err != nil {
		//fmt.Println(err)
	}
	return reply
}

func parse(data string) classifications {
	var classification classifications

	err := json.Unmarshal([]byte(data), &classification)
	if err != nil {
		fmt.Println(err)
	}
	return classification
}
