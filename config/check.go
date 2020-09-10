package config

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/gokafka.poc/pkg/fmts"
	"github.com/jeffotoni/gokafka.poc/service/check"
	"gopkg.in/yaml.v2"
)

func Check() {
	fmts.Println(gcolor.YellowCor("....check in 1s"))
	time.Sleep(time.Second)
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
