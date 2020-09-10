package config

type C struct {
	Kafka struct {
		Host              string   `yaml:"host"`
		HostConsumer      []string `yaml:"host_consumer,flow"`
		HostProducer      []string `yaml:"host_producer,flow"`
		Retentive         string   `yaml:"retentive"`
		Group             string   `yaml:"group"`
		OffsetReset       string   `yaml:"offset_reset"`
		TopicGame         string   `yaml:"topic_game"`
		NumPartitions     int      `yaml:"num_partitions"`
		ReplicationFactor int      `yaml:"replication_factor"`
		PolicyCleanup     string   `yaml:"policy_cleanup"`
	}

	Sys struct {
		Debug  bool `yaml:"debug"`
		NumCPU int  `yaml:"numcpu"`
	}

	Http struct {
		Port string `yaml:"port"`
	}
}
