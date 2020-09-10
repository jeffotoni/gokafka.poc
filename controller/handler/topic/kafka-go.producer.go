package htopic

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaWriter = getKafkaWriter(Conf.Kafka.HostProducer, Conf.Kafka.TopicGame) // kafkaWriter   *kafka.Writer
)

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

func Producer(c *fiber.Ctx) {
	if c.Get("X-Key-User") != "x&*0987665.33.43.x.2.2.1.o9*" {
		log.Println("Error kafka sem body, msg obritagoria")
		c.Status(400)
	}

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
			log.Println("error kafka WriteMessages:", err)
			return
		}
	}(body)

	c.Status(200)
}
