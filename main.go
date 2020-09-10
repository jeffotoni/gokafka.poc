package main

import (
	"runtime"

	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
	"github.com/jeffotoni/gokafka.poc/config"
	hping "github.com/jeffotoni/gokafka.poc/controller/handler/ping"
	htopic "github.com/jeffotoni/gokafka.poc/controller/handler/topic"
	mw "github.com/jeffotoni/gokafka.poc/controller/middleware"
)

var (
	conf = config.Config()
)

func init() {
	if conf.Sys.NumCPU > 0 {
		runtime.GOMAXPROCS(conf.Sys.NumCPU)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// exec chec in
	// system
	config.Check()
}

var (
	//Tamanho para todas as requisições
	sizeBodyDefault  = 3 * 1024 * 1024 //maximo para requests normais
	sizeBodyProducer = 1 * 1024 * 1024 // maximo geral
)

func main() {

	// app http
	app := fiber.New()

	// max body global
	app.Settings.BodyLimit = sizeBodyDefault

	//Rate Limite
	app.Use(limiter.New(limiter.Config{
		Timeout:    1,
		Max:        100000,
		Filter:     nil,
		StatusCode: 401,
		Message:    `{"msg":"Much Request #bloqued"}`,
	}))
	//==========================================

	// middlewares
	mw.Cors(app)
	mw.Logger(app)
	mw.Compress(app)

	// GET ping -> pong
	app.Get("/ping", hping.Ping)

	// POST ping -> pong
	app.Post("/ping", hping.Ping)

	// handler create topic kafka-go
	app.Post("/topic", htopic.Create)

	// handler list all topics sarama
	app.Get("/topic", htopic.ListAllTopics)

	// handler delete topic sarama
	app.Delete("/topic/:topic", htopic.Delete)

	// middleware size producer
	app.Use(mw.MaxBody(sizeBodyProducer))

	// handler producer kafka-go
	app.Post("/producer", htopic.Producer)

	app.Listen(conf.Http.Port)
}
