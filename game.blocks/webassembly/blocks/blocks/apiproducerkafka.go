package blocks

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	XKey     = "x&*0987665.33.43.x.2.2.1.o9*"
	EndPoint = "http://localhost:8181/producer"
)

type producerkafka struct {
	Key      string
	Name     string
	Score    int
	DataHora string
}

func ApiProducerKafka(key, name string, score int) {
	datahora := time.Now().Format("2006-01-02 15:04:05")
	var pk = &producerkafka{Key: key, Name: name, Score: score, DataHora: datahora}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(pk)
	if err != nil {
		log.Println("Error json.NewEncoder:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", EndPoint, buf) // bytes.NewBuffer(jsonB)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Key-User", XKey)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error http.DefaultClient:", err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error ioutil.ReadAll:", err)
		return
	}

	defer response.Body.Close()
	println("body:", string(body))
	println("code: ", response.StatusCode)
}
