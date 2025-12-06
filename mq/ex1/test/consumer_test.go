package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
)

func TestConsumer(t *testing.T) {
	cfg := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup([]string{"192.168.1.3:9094"}, "test_group", cfg)
	require.NoError(t, err)

	err = consumer.Consume(context.Background(), []string{"test"}, &TestConsumerHandler{})
	t.Log(err)
}

type TestConsumerHandler struct {
}

func (t *TestConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	slog.Info("Setup")
	// topic => offset
	partitions := session.Claims()["test"]
	for _, partition := range partitions {
		session.ResetOffset("test", partition, sarama.OffsetOldest, "consumer metadata")
	}
	return nil
}

func (t *TestConsumerHandler) Cleanup(p0 sarama.ConsumerGroupSession) error {
	slog.Info("Cleanup")
	return nil
}

// 连接生命周期会话
func (t *TestConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Println(string(msg.Value))
		session.MarkMessage(msg, "consumer metadata")
	}

	//
	return nil
}

func Decode[T any](bs []byte) (T, error) {
	var v T
	if err := json.Unmarshal(bs, &v); err != nil {
		return v, err
	}

	return v, nil
}
