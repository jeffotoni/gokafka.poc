package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"

    //"os"
    "flag"
    "strings"
    "time"

    "github.com/jeffotoni/gcolor"
    "github.com/jeffotoni/gconcat"
    kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
    brokers := strings.Split(kafkaURL, ",")
    return kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        GroupID:  groupID,
        Topic:    topic,
        MinBytes: 10e3, // 10KB
        MaxBytes: 10e6, // 10MB
    })
}

func main() {
    flagHost := flag.String("host", "localhost:9092", "string")
    flagTopic := flag.String("topic", "gogameblocks", "string")
    flagGroup := flag.String("group", "logger-group1", "string")
    flag.Parse()

    // get kafka reader using environment variables.
    kafkaURL := *flagHost
    topic := *flagTopic
    groupID := *flagGroup

    println(".......................................................")
    println(gcolor.YellowCor(gconcat.Build("Url:   ", kafkaURL)))
    println(gcolor.YellowCor(gconcat.Build("Topic: ", topic)))
    println(gcolor.YellowCor(gconcat.Build("Group: ", groupID)))

    reader := getKafkaReader(kafkaURL, topic, groupID)
    defer reader.Close()

    forever := make(chan struct{})
    ctx, cancel := context.WithCancel(context.Background())

    go func(ctx context.Context, cancelFunc context.CancelFunc) {
        sigchan := make(chan os.Signal)
        signal.Notify(sigchan, os.Interrupt)
        signal.Notify(sigchan, syscall.SIGTERM)
        signal.Notify(sigchan, syscall.SIGHUP)
        <-sigchan
        cancelFunc()
        println(ctx.Err().Error())
        forever <- struct{}{}
    }(ctx, cancel)

    go func(ctx context.Context) {
        for {
            m, err := reader.ReadMessage(context.Background())
            if err != nil {
                log.Fatalln(err)
            }
            fmt.Printf("message at date:%s topic:%v partition:%v offset:%v  %s = %s\n", time.Now().Format("2006-01-02 15:04:05"),
                m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
            <-time.After(50 * time.Millisecond)
        }
    }(ctx)

    <-forever
    println("finish...")
}
