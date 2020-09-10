package config

import (
	"io/ioutil"
	"os"

	"github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/gokafka.poc/pkg/fmts"
	skafka "github.com/jeffotoni/gokafka.poc/pkg/kafka"
	"github.com/jeffotoni/gokafka.poc/service/check"
	"gopkg.in/yaml.v2"
)

func CreateTopicGame() {
	conf := Config()
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Check seu topic game...          "))
	err := skafka.CreateTopicKafka{
		Host:              conf.Kafka.Host,
		PolicyCleanup:     conf.Kafka.PolicyCleanup,
		Name:              conf.Kafka.TopicGame,
		NumPartitions:     conf.Kafka.NumPartitions,
		ReplicationFactor: conf.Kafka.ReplicationFactor}.TopicCreate()
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		return
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
}

func Check() {
	CreateTopicGame()
	CheckStatuskafka()
	fmts.Println(gcolor.YellowCor("........................................."))
}

func Config() (c *C) {
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmts.Println(gcolor.RedCor("..........................................."))
		fmts.Println(gcolor.RedCor("O arquivo config.yaml precisa está no raiz."))
		fmts.Println(gcolor.RedCor("..........................................."))
		os.Exit(0)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		fmts.Println(gcolor.RedCor("..........................................."))
		fmts.Println(gcolor.RedCor("Erro ao fazer parse em config.yaml"))
		fmts.Println(gcolor.RedCor("..........................................."))
		os.Exit(0)
	}
	return c
}

func CheckStatuskafka() {
	c := Config()
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Check seu Kafka Consumer..          "))
	err := check.CheckConsumerKafka(c.Kafka.HostConsumer, c.Kafka.Retentive, c.Kafka.Group)
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		return
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Check seu Kafka Producer..          "))
	err = check.CheckConsumerKafka(c.Kafka.HostProducer, c.Kafka.Retentive, c.Kafka.Group)
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		return
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
}
