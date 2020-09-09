package config

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/jeffotoni/enginer/retrymanage/pkg/fmts"
	"github.com/jeffotoni/enginer/retrymanage/repo/check"
	"github.com/jeffotoni/gcolor"
	"gopkg.in/yaml.v2"
)

type C struct {
	Rethinkdb struct {
		Host      []string `yaml:"host,flow"`
		Db        string   `yaml:"db"`
		User      string   `yaml:"user"`
		Password  string   `yaml:"password"`
		Topic     string   `yaml:"topic"`
		DeadTopic string   `yaml:"dead_topic"`
	}

	Metrics struct {
		Port          string `yaml:"port"`
		Read          int    `yaml:"read"`
		DelaySeconds  int    `yaml:"delaySeconds"`
		PeriodSeconds int    `yaml:"periodSeconds"`
	}

	Kafka struct {
		HostConsumer []string `yaml:"host_consumer,flow"`
		HostProducer []string `yaml:"host_producer,flow"`
		//Host        []string `yaml:"host,flow"`
		Retentive   string `yaml:"retentive"`
		Group       string `yaml:"group"`
		OffsetReset string `yaml:"offset_reset"`
	}

	Sys struct {
		DelaySeconds int    `yaml:"delaySeconds"`
		Read         int    `yaml:"read"`
		Write        int    `yaml:"write"`
		TtlLocker    int    `yaml:"ttlLocker"`
		Timezone     string `yaml:"timezone"`
		DelayError   int    `yaml:"delayError"`
		Readlocker   int    `yaml:"readLocker"`
		Debug        bool   `yaml:"debug"`
		NumCPU       int    `yaml:"numcpu"`
	}

	Http struct {
		Port string `yaml:"port"`
	}
}

func init() {
	fmts.Println(gcolor.YellowCor("....check in 1s"))
	time.Sleep(time.Second)
	CheckStatusDb()
	CheckStatusTable()
	CheckStatuskafka()
	CheckRegsOffLine()
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

func CheckStatusDb() {
	c := Config()
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Checando seu DB RethinkDB..       "))
	err := check.ConnRethinkDB(c.Rethinkdb.Host, c.Rethinkdb.Db, c.Rethinkdb.User, c.Rethinkdb.Password)
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		os.Exit(0)
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
	time.Sleep(time.Second)
}

func CheckStatusTable() {
	c := Config()
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Checando sua tabela RethinkDB..   "))
	err := check.ConnRethinkTable(c.Rethinkdb.Host, c.Rethinkdb.Db, c.Rethinkdb.Topic, c.Rethinkdb.User, c.Rethinkdb.Password)
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		os.Exit(0)
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
}

func CheckStatuskafka() {
	c := Config()
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Checando seu Kafka Consumer..          "))
	err := check.CheckConsumerKafka(c.Kafka.HostConsumer, c.Kafka.Retentive, c.Kafka.Group)
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		//os.Exit(0)
		return
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Checando seu Kafka Producer..          "))
	err = check.CheckConsumerKafka(c.Kafka.HostProducer, c.Kafka.Retentive, c.Kafka.Group)
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		//os.Exit(0)
		return
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
	time.Sleep(time.Second)
}

func CheckRegsOffLine() {
	c := Config()
	fmts.Println(gcolor.YellowCor("........................................."))
	fmts.Print(gcolor.GreenCor("Checando registros offline RethinkDB..   "))
	err := check.UpdateOffLine(c.Rethinkdb.Host, c.Rethinkdb.Db, c.Rethinkdb.Topic, c.Rethinkdb.User, c.Rethinkdb.Password)
	if err != nil {
		fmts.Println(gcolor.RedCor(" [error] => "), err.Error())
		os.Exit(0)
	}
	fmts.Println(gcolor.YellowCor(" [ok]"))
}
