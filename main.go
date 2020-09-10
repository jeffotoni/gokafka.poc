package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/google/uuid"

	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
	"github.com/jeffotoni/gokafka.poc/config"
	hping "github.com/jeffotoni/gokafka.poc/controller/handler/ping"
	htopic "github.com/jeffotoni/gokafka.poc/controller/handler/topic"
	mw "github.com/jeffotoni/gokafka.poc/controller/middleware"
	kafka "github.com/segmentio/kafka-go"
)

var (
	conf          = config.Config() // faz a checagem
	nameTopicGame = conf.Kafka.TopicGame
	//kafkaWriter   *kafka.Writer
	kafkaWriter = getKafkaWriter(conf.Kafka.HostProducer, nameTopicGame)
)

func init() {
	if conf.Sys.NumCPU > 0 {
		runtime.GOMAXPROCS(conf.Sys.NumCPU)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	config.Check()
}

// kafka-go
func mwProducer(kafkaWriter *kafka.Writer) fiber.Handler {
	return func(c *fiber.Ctx) {
	}
}

func Producer(c *fiber.Ctx) {
	println("header:", c.Get("X-Key-User"))
	body := c.Body()
	if len(body) <= 0 {
		log.Println("Error kafka sem body, msg obritagoria")
		c.Status(400)
		return
	}

	//kafkaWriter := getKafkaWriter(conf.Kafka.HostProducer, nameTopicGame)
	go func(body string) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		Key := uuid.New().String()
		msg := kafka.Message{
			//Key:   []byte(fmt.Sprintf("address-%s", c.IP())),
			Key:   []byte(Key),
			Value: []byte(body),
		}

		err := kafkaWriter.WriteMessages(ctx, msg)
		if err != nil {
			log.Println("kafka:", err)
			return
		}
	}(body)

	c.Status(200)
}

// kafka-go
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
	//defer kafkaWriter.Close()
	// app http
	app := fiber.New()
	app.Settings.BodyLimit = sizeBodyFiber
	app.Use(mw.MaxBody(sizeBodyDefault))

	//Rate Limite
	app.Use(limiter.New(limiter.Config{
		Timeout:    1,
		Max:        100000,
		Filter:     nil,
		StatusCode: 401,
		Message:    `{"msg":"Much Request #bloqued"}`,
	}))
	//==========================================

	// middlewares
	mw.Cors(app)
	mw.Logger(app)
	mw.Compress(app)

	// GET ping -> pong
	app.Get("/ping", hping.Ping)

	// POST ping -> pong
	app.Post("/ping", hping.Ping)

	// handler create topic kafka-go
	app.Post("/topic", htopic.Create)

	// handler list all topics sarama
	app.Get("/topic", htopic.ListAllTopics)

	// handler delete topic sarama
	app.Delete("/topic/:topic", htopic.Delete)

	// handler producer kafka-go
	app.Post("/producer", Producer)

	app.Listen(conf.Http.Port)
}
