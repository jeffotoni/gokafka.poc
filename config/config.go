package config

type C struct {
	Kafka struct {
		HostConsumer []string `yaml:"host_consumer,flow"`
		HostProducer []string `yaml:"host_producer,flow"`
		Retentive    string   `yaml:"retentive"`
		Group        string   `yaml:"group"`
		OffsetReset  string   `yaml:"offset_reset"`
	}

	Sys struct {
		Debug  bool `yaml:"debug"`
		NumCPU int  `yaml:"numcpu"`
	}

	Http struct {
		Port string `yaml:"port"`
	}
}
