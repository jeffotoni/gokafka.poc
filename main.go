package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"

	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
	"github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/gokafka.poc/config"
	hping "github.com/jeffotoni/gokafka.poc/controller/handler/ping"
	mw "github.com/jeffotoni/gokafka.poc/controller/middleware"
	"github.com/jeffotoni/gokafka.poc/pkg/fmts"
	kafka "github.com/segmentio/kafka-go"
)

// type producerkafka struct {
// 	Key      string
// 	Name     string
// 	Score    int
// 	DataHora string
// }

type createTopic struct {
	Name              string `json:"name"`
	NumPartitions     int    `json:"partition"`
	ReplicationFactor int    `json:"replication_factor"`
}

var (
	conf          = config.Config()
	nameTopicGame = "gogameblocks"
	//kafkaWriter   *kafka.Writer
	//kafkaWriter = getKafkaWriter(c.Kafka.HostProducer, nameTopicGame)
)

func init() {
	if conf.Sys.NumCPU > 0 {
		runtime.GOMAXPROCS(conf.Sys.NumCPU)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	config.Check()

	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Check seu topic game...          "))
	err := createTopic{Name: nameTopicGame, NumPartitions: 1, ReplicationFactor: 1}.TopicCreate()
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		return
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
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

	kafkaWriter := getKafkaWriter(conf.Kafka.HostProducer, nameTopicGame)
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
			//c.Status(400)
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

// kafka-go
func mwCreateTopic(c *fiber.Ctx) {
	var ct createTopic
	err := c.BodyParser(&ct)
	if err != nil {
		log.Println("error BodyParser:", err.Error())
		c.Status(400)
		return
	}

	err = ct.TopicCreate()
	if err != nil {
		log.Println("error createTopic:", err.Error())
		c.Status(400)
		return
	}

	c.Status(200)
}

func (ct createTopic) TopicCreate() (err error) {
	config := []kafka.ConfigEntry{
		{ConfigName: "cleanup.policy", ConfigValue: "compact"},
		{ConfigName: "segment.bytes", ConfigValue: "10240"},
	}

	TopicConfig := kafka.TopicConfig{
		Topic:             ct.Name,
		NumPartitions:     ct.NumPartitions,
		ReplicationFactor: ct.ReplicationFactor,
		ConfigEntries:     config,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	conn, err := dialer.DialContext(ctx, "tcp", "localhost:9092")
	if err != nil {
		return
	}
	err = conn.CreateTopics(TopicConfig)
	if err != nil {
		return
	}
	return
}

// sarama
func ListAllTopics(c *fiber.Ctx) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	//kafka end point
	brokers := []string{"localhost:9092"}

	//get broker
	cluster, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Println("Error NewConsumer", err)
		c.Status(400)
		return
	}

	defer func() {
		if err := cluster.Close(); err != nil {
			log.Println("Error NewConsumer", err)
			c.Status(400)
			return
		}
	}()

	//get all topic from cluster
	topics, _ := cluster.Topics()
	// for index := range topics {
	// 	fmt.Println(topics[index])
	// }
	c.Status(200).JSON(topics)
}

// sarama
func DeleteTopic(c *fiber.Ctx) {
	topic := c.Params("topic")
	if len(topic) <= 0 {
		c.Status(400).JSON(`{"msg":"Error remove topic!"}`)
		return
	}
	config := sarama.NewConfig()
	config.Version = sarama.V1_1_0_0
	clusterAdmin, err := sarama.NewClusterAdmin([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Println("error NewClusterAdmin:", err.Error())
		c.Status(400)
		return
	}
	clusterAdmin.DeleteTopic(topic)
	c.Status(200)
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
	app.Post("/ping", hping.Ping)
	app.Post("/topic", mwCreateTopic)
	app.Get("/topic", ListAllTopics)
	app.Delete("/topic/:topic", DeleteTopic)
	app.Post("/producer", Producer)
	app.Listen(conf.Http.Port)
}
