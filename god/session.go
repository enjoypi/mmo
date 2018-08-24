package god

import (
	"fmt"
	"reflect"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

// service consists of the information of the server serving this service and
// the methods in this service.
type service struct {
	server interface{} // the server for service methods
	md     map[string]*grpc.MethodDesc
}

type Session struct {
	*amqp.Channel
	m map[string]*service // service name -> service info
}

func NewSession() (*Session, error) {
	ch, err := self.Connection.Channel()
	if err != nil {
		return nil, err
	}

	var s Session
	s.Channel = ch
	return &s, nil
}

func combine(routingKeyType uint16, routingKey uint64) string {
	return fmt.Sprintf("%d.%d", routingKeyType, routingKey)
}

func (s *Session) Declare(exchange string) error {
	return s.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
}

func (s *Session) Post(
	exchange string,
	routingKeyType uint16, routingKey uint64,
	service string,
	method string,
	msg proto.Message) error {

	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return s.Publish(
		exchange, // exchange
		combine(routingKeyType, routingKey), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/octet-stream",
			Type:         proto.MessageName(msg),
			MessageId:    method,
			Body:         body,
			AppId:        service,
		})
}

func (s *Session) Subscribe(exchange string,
	routingKeyType uint16, routingKey uint64) (string, error) {
	err := s.Declare(exchange)
	if err != nil {
		return "", err
	}

	q, err := s.declareQueue()
	if err != nil {
		return "", err
	}

	return q.Name,
		s.bind(exchange,
			q.Name,
			combine(routingKeyType, routingKey))
}

func (s *Session) declareQueue() (amqp.Queue, error) {
	return s.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func (s *Session) bind(exchange string, queue string, routingKey string) error {
	return s.QueueBind(
		queue,      // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil)
}

func (s *Session) Pull(queue string) (<-chan amqp.Delivery, error) {
	return s.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

func (s *Session) register(sd *grpc.ServiceDesc, ss interface{}) {
	if s.m == nil {
		s.m = make(map[string]*service)
	}

	if _, ok := s.m[sd.ServiceName]; ok {
	}
	srv := &service{
		server: ss,
		md:     make(map[string]*grpc.MethodDesc),
	}
	for i := range sd.Methods {
		d := &sd.Methods[i]
		srv.md[d.MethodName] = d
	}
	s.m[sd.ServiceName] = srv
}

type Dispatch func(service string, method string, msg proto.Message) error

func (s *Session) Handle(queue string, dispatch Dispatch) error {
	msgs, err := s.Pull(queue)
	if err != nil {
		return err
	}

	for d := range msgs {
		if dispatch != nil {
			msgType := proto.MessageType(d.Type).Elem()
			msg := reflect.New(msgType).Interface().(proto.Message)
			if err := proto.Unmarshal(d.Body, msg); err != nil {
				return err
			}

			if err := dispatch(d.AppId, d.MessageId, msg); err != nil {
				return err
			}
		} else {
			srv := s.m[d.AppId]
			if srv == nil {
				break
			}

			md := srv.md[d.MessageId]
			if md == nil {
				break
			}

			_, err := md.Handler(srv.server, nil,
				func(msg interface{}) error {
					return proto.Unmarshal(d.Body, msg.(proto.Message))
				})
			if err != nil {
				return err
			}

		}
		d.Ack(false)
	}
	return nil
}
