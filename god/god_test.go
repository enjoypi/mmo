package god

import (
	"ext"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestGod(t *testing.T) {
	require.NoError(t,
		Start("amqp://guest:guest@localhost:5672/", 0, ext.RandomUint64()),
		"start god")

	var producer, consumer *Session
	var err error

	consumer, err = NewSession()
	require.NoError(t, err, "new consumer")

	exchange := "god.test"
	routingKeyType := ext.RandomUint16()
	routingKey := ext.RandomUint64()

	q, err := consumer.Subscribe(exchange, routingKeyType, routingKey)
	require.NoError(t, err, "pull msgs")
	ci := int64(0)
	go consumer.Handle(q,
		func(service string, method string, msg proto.Message) error {
			require.Equal(t, method, "Test")
			require.Equal(t, service, "Test")
			test := msg.(*Test)
			require.Equal(t, test.Count, ci)
			ci++
			return nil
		})

	producer, err = NewSession()
	require.NoError(t, err, "new producer")

	for i := int64(0); i < 1000; i++ {
		var test Test
		test.Count = i
		err = producer.Post(exchange,
			routingKeyType, routingKey,
			"Test", "Test", &test)
		require.NoError(t, err, "post")
	}

	time.Sleep(time.Millisecond * 100)
}

func BenchmarkStartNode(b *testing.B) {
}
