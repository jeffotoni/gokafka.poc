package htopic

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber"
)

// sarama
func Delete(c *fiber.Ctx) {
	topic := c.Params("topic")
	if len(topic) <= 0 {
		c.Status(400).JSON(`{"msg":"Error remove topic!"}`)
		return
	}

	config := sarama.NewConfig()
	config.Version = sarama.V1_1_0_0
	clusterAdmin, err := sarama.NewClusterAdmin(Conf.Kafka.HostConsumer, config)
	if err != nil {
		log.Println("error NewClusterAdmin:", err.Error())
		c.Status(400)
		return
	}
	clusterAdmin.DeleteTopic(topic)
	c.Status(200)
}
