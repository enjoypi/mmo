package ext

type Protocol interface {
}

type protocol struct {
}

func NewProtocol() Protocol {
	return &protocol{}
}
