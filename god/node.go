package god

import (
	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/nsf/termbox-go"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
)

const (
	adminService = "god.Admin"
)

type node struct {
	*amqp.Connection
	*Session

	kind uint16
	id   uint64
}

var _ AdminServer = (*node)(nil)

var self node

func Start(url string, nodeType uint16, nodeID uint64) error {
	c, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	self.Connection = c
	s, err := NewSession()
	if err != nil {
		s.Close()
		return err
	}

	q, err := s.Subscribe(adminService, nodeType, nodeID)
	if err != nil {
		s.Close()
		return err
	}

	self.Session = s
	self.kind = nodeType
	self.id = nodeID

	var req AuthReq
	req.ID = nodeID
	postAdmin("Auth", &req)

	self.register(&_Admin_serviceDesc, &self)
	go self.Handle(q, nil)
	return nil
}

func RunConsole() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			termbox.Flush()
		}
	}
}

func Close() {
	self.Close()
}

func postAdmin(method string, msg proto.Message) error {
	return self.Post(adminService,
		self.kind, self.id,
		adminService, method, msg)
}

func (n *node) Auth(c context.Context, req *AuthReq) (*AuthAck, error) {
	log.Infof("%#v", req)
	return &AuthAck{Code: ErrorCode_OK}, nil
}
