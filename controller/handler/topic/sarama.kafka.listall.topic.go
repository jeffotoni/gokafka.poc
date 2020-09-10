package htopic

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber"
)

// sarama
func ListAllTopics(c *fiber.Ctx) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	//kafka end point
	brokers := Conf.Kafka.HostConsumer

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
	c.Status(200).JSON(topics)
}
