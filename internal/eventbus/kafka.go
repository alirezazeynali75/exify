package eventbus

import (
	"context"
	"log/slog"

	"github.com/IBM/sarama"
)

type MessageToPublish struct {
	Topic string 
	Value string
}

type Kafka struct {
	logger *slog.Logger
	brokers []string
	configs *sarama.Config
}

func NewKafka(
	logger *slog.Logger,
	brokers []string,
	configs *sarama.Config,
) *Kafka {
	return &Kafka{
		logger: logger,
		brokers: brokers,
		configs: configs,
	}
}

func (k *Kafka) Produce(ctx context.Context, msgs []MessageToPublish) error {
	logger := k.logger.With(slog.Uint64("len", uint64(len(msgs))))

	logger.Debug("going to produce some message")

	producer, err := sarama.NewSyncProducer(k.brokers, k.configs)

	if err != nil {
		return err
	}
	defer producer.Close()

	msgsToProduce := make([]*sarama.ProducerMessage, len(msgs))

	for i, msg := range msgs {
		msgsToProduce[i] = &sarama.ProducerMessage{
			Topic: msg.Topic,
			Value: sarama.StringEncoder(msg.Value),
		}
	}

	err = producer.SendMessages(msgsToProduce)
	if err != nil {
		logger.Error("there is an error to produce messages")
		return err
	}
	return nil
}