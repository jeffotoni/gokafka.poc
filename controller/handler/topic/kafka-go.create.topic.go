package htopic

import (
	"log"

	"github.com/gofiber/fiber"
	skafka "github.com/jeffotoni/gokafka.poc/pkg/kafka"
)

// kafka-go
func Create(c *fiber.Ctx) {
	var ct skafka.CreateTopicKafka
	err := c.BodyParser(&ct)
	if err != nil {
		log.Println("error BodyParser:", err.Error())
		c.Status(400)
		return
	}

	ct.Host = Conf.Kafka.Host
	ct.PolicyCleanup = Conf.Kafka.PolicyCleanup
	err = ct.TopicCreate()
	if err != nil {
		log.Println("error createTopic:", err.Error())
		c.Status(400)
		return
	}
	c.Status(200)
}
