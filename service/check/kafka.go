package check

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

func CheckConsumerKafka(KafkaUrl []string, topic, grupo_id string) error {
	partition := 0
	for _, host := range KafkaUrl {
		//println("host:", host)
		_, err := kafka.DialLeader(context.Background(), "tcp", host, topic, partition)
		if err != nil {
			return err
		}
	}
	return nil
}
