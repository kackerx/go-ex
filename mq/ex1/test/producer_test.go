package test

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"192.168.1.3:9094"}, cfg)
	require.NoError(t, err)

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Headers: []sarama.RecordHeader{{
			Key:   []byte("trace_id"),
			Value: []byte("xxxx-xxxx-xxxx-xxxx"),
		}},
		Metadata: "this is metadata",
		Topic:    "test",
		Value:    sarama.StringEncoder("hllo"),
	})
	assert.NoError(t, err)
}

func TestAsyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true // 关心成功消息
	cfg.Producer.Return.Errors = true    // 关心错误消息
	producer, err := sarama.NewAsyncProducer([]string{"192.168.1.3:9094"}, cfg)
	require.NoError(t, err)

	msgCh := producer.Input()
	msgCh <- &sarama.ProducerMessage{
		Headers: []sarama.RecordHeader{{
			Key:   []byte("trace_id"),
			Value: []byte("xxxx-xxxx-xxxx-xxxx"),
		}},
		Metadata: "this is metadata",
		Topic:    "test",
		Value:    sarama.StringEncoder("async hello"),
	}

	errCh := producer.Errors()
	successCh := producer.Successes()
	select {
	case err := <-errCh:
		t.Logf("error: %v", err)
	case success := <-successCh:
		t.Logf("success: %v", success)
	}
}
