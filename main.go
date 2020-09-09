package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
	hping "github.com/jeffotoni/gokafka.poc/controller/handler/hping"
	mw "github.com/jeffotoni/gokafka.poc/controller/middleware"
	kafka "github.com/segmentio/kafka-go"
)

func producerHandler(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Println("Body:", err)
				return
			}
			msg := kafka.Message{
				Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
				Value: body,
			}
			err = kafkaWriter.WriteMessages(req.Context(), msg)
			if err != nil {
				wrt.WriteHeader(200)
				wrt.Write([]byte(err.Error()))
				log.Println("kafka:", err)
			}
		} else {
			wrt.WriteHeader(400)
			wrt.Write([]byte(`{"msg":"Method not allowed"}`))
		}
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

var (
	//Tamanho para todas as requisições
	sizeBodyDefault = 3 * 1024 * 1024  //maximo para requests normais
	sizeBodyFiber   = 10 * 1024 * 1024 // maximo geral
)

func main() {

	kafkaWriter := getKafkaWriter("localhost:9092", "test")

	defer kafkaWriter.Close()

	app := fiber.New()

	app.Settings.BodyLimit = sizeBodyFiber

	app.Use(mw.MaxBody(sizeBodyDefault))

	//Rate Limite
	app.Use(limiter.New(limiter.Config{
		Timeout:    1,
		Max:        1000,
		Filter:     nil,
		StatusCode: 401,
		Message:    `{"msg":"Much Request #bloqued"}`,
	}))
	//==========================================

	mw.Cors(app)

	mw.Logger(app)

	mw.Compress(app)

	app.Get("/ping", hping.Ping)
	app.Post("/producer", hping.Ping)

	app.Listen(8181)

	// Add handle func for producer.
	//http.HandleFunc("/producer", producerHandler(kafkaWriter))
	// Run the web server.
	//fmt.Println("start producer-api 8181 ... !!")
	//log.Fatal(http.ListenAndServe(":8181", nil))
}
