package config

import (
	"time"

	"github.com/IBM/sarama"
)


type KafkaConfig struct {
	Brokers  []string `env:"KAFKA_BROKERS"   envSeparator:"," envDefault:"localhost:9092"`
}


func (conf KafkaConfig) ToSaramaConfig() (*sarama.Config, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll         // Wait for all in-sync replicas to acknowledge
	config.Producer.Retry.Max = 5                            // Retry up to 5 times to produce the message
	config.Producer.Return.Successes = true                 // Return successfully produced messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Flush.Messages = 10  
	return config, nil
}
