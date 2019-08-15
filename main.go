package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

var dataFile = "small.txt"
var trainData []string

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go Writer(ws, trainData[0])
	Reader(ws, trainData)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	readLines()
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

// read the file line by line
func readLines() error {
	file, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trainData = append(trainData, scanner.Text())
	}
	return scanner.Err()
}
