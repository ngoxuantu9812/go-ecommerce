package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_vip"
)

// Consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
	})
}

// Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.PostForm("message"), c.PostForm("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s
	jsonBody, _ := json.Marshal(body)
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err":     "",
		"message": msg.Value,
	})
}

func RegisterConsumerATC(id int) {
	kafkaGroupId := "consumer-group-" + strconv.Itoa(id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) listen phien ATC:: \n", id)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) read error: %s\n", id, err)
		}
		fmt.Printf("Consumer(%d), hong topic: %v, partition:%v,offset:%v, time:%d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()
	r.POST("action/stock", actionStock)

	// dang ky 2 user de mua ATC trong (1) (2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)

	r.Run(":8999")
}
