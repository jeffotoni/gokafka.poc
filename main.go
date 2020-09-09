package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
	hping "github.com/jeffotoni/gokafka.poc/controller/handler/ping"
	mw "github.com/jeffotoni/gokafka.poc/controller/middleware"
	kafka "github.com/segmentio/kafka-go"
)

func mwProducer(kafkaWriter *kafka.Writer) fiber.Handler {
	return func(c *fiber.Ctx) {
		if len(c.Body()) <= 0 {
			log.Println("Error kafka sem body, msg obritagoria")
			c.Next(fiber.ErrBadRequest)
			return
		}
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", c.IP())),
			Value: []byte(c.Body()),
		}

		err := kafkaWriter.WriteMessages(c.Context(), msg)
		if err != nil {
			log.Println("kafka:", err)
			c.Next(fiber.ErrBadRequest)
			return
		}
		c.Next()
	}
}

func getKafkaWriter(kafkaURL []string, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      kafkaURL,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: -1, // acks = 0, acks = 1, acks = -1
		//CommitInterval: time.Second * 5, // flushes commits to Kafka every second
	})
}

var (
	//Tamanho para todas as requisições
	sizeBodyDefault = 3 * 1024 * 1024  //maximo para requests normais
	sizeBodyFiber   = 10 * 1024 * 1024 // maximo geral
)

func main() {

	// config := []kafka.ConfigEntry{
	// 	{ConfigName: "cleanup.policy", ConfigValue: "compact"},
	// 	{ConfigName: "segment.bytes", ConfigValue: "10240"},
	// }

	// topic := kafka.TopicConfig{
	// 	Topic:         kafkaCheckpointTopic,
	// 	NumPartitions: 1,
	// 	ConfigEntries: config,
	// }

	// error := conn.CreateTopics(topic)
	// if error != nil {
	// 	glog.Error(error)
	// }

	kafkaWriter := getKafkaWriter([]string{"localhost:9092", "localhost:9092"}, "test")

	defer kafkaWriter.Close()

	app := fiber.New()

	app.Settings.BodyLimit = sizeBodyFiber

	app.Use(mw.MaxBody(sizeBodyDefault))

	//Rate Limite
	app.Use(limiter.New(limiter.Config{
		Timeout:    1,
		Max:        10000,
		Filter:     nil,
		StatusCode: 401,
		Message:    `{"msg":"Much Request #bloqued"}`,
	}))
	//==========================================

	mw.Cors(app)
	mw.Logger(app)
	mw.Compress(app)

	app.Get("/ping", hping.Ping)
	app.Post("/producer", mwProducer(kafkaWriter))

	app.Listen(8181)

	// Add handle func for producer.
	//http.HandleFunc("/producer", producerHandler(kafkaWriter))
	// Run the web server.
	//fmt.Println("start producer-api 8181 ... !!")
	//log.Fatal(http.ListenAndServe(":8181", nil))
}
