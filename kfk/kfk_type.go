package kfk

type Config struct {
	Addr     string
	ClientId string
	Producer ProducerConfig
	Consumer ConsumerConfig
}

type ProducerConfig struct {
}

type ConsumerConfig struct {
}
