package skafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	SegmentBytes = "1024000" // 1Mb
)

type CreateTopicKafka struct {
	Host              string
	PolicyCleanup     string `json:"policy_cleanup"`
	Name              string `json:"name"`
	NumPartitions     int    `json:"partition"`
	ReplicationFactor int    `json:"replication_factor"`
}

func (ct CreateTopicKafka) TopicCreate() (err error) {
	config := []kafka.ConfigEntry{
		{ConfigName: "cleanup.policy", ConfigValue: ct.PolicyCleanup}, //"compact" or delete
		{ConfigName: "segment.bytes", ConfigValue: SegmentBytes},
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

	conn, err := dialer.DialContext(ctx, "tcp", ct.Host)
	if err != nil {
		return
	}
	err = conn.CreateTopics(TopicConfig)
	if err != nil {
		return
	}
	return
}
