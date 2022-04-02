package kfk

import (
	"encoding/json"
	"strings"

	"github.com/Shopify/sarama"

	"github.com/yyliziqiu/waf/logs"
)

var producer sarama.SyncProducer

func InitializeProducer(c Config) {
	var err error

	config := sarama.NewConfig()
	config.ClientID = c.ClientId
	config.Version = sarama.V0_10_2_0
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer(strings.Split(c.Addr, ","), config)
	if err != nil {
		logs.Fatal(err)
	}
	// defer producer.Close()
}

func SendMessage(pm *sarama.ProducerMessage) (int32, int64, error) {
	return producer.SendMessage(pm)
}

func Send(topic string, msg interface{}) error {
	msgValue, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, _, err = SendMessage(&sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(msgValue)})
	if err != nil {
		return err
	}
	return nil
}
