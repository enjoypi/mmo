package ext

const (
	_ = iota
	constCmdSet
	constCmdGet
	constCmdDelete
)

type cmdInfo struct {
	cmd int
	k   interface{}
	v   interface{}
	ret chan interface{}
}

type chanMap struct {
	m       map[interface{}]interface{}
	cmdChan chan cmdInfo // get chan
}

func NewChanMap() ParallelMap {
	c := chanMap{}
	c.m = make(map[interface{}]interface{})
	c.cmdChan = make(chan cmdInfo)
	go c.runCmd()
	return &c
}

func (c *chanMap) Set(k, v interface{}) {
	c.pushCmd(constCmdSet, k, v, nil)
}

func (c *chanMap) Get(k interface{}) interface{} {
	ret := make(chan interface{}, 1)
	c.pushCmd(constCmdGet, k, nil, ret)
	return <-ret
}

func (c *chanMap) Delete(k interface{}) {
	c.pushCmd(constCmdDelete, k, nil, nil)
}

func (c *chanMap) Len() int {
	return len(c.m)
}

func (c *chanMap) pushCmd(cmd int, k, v interface{}, ret chan interface{}) {
	c.cmdChan <- cmdInfo{cmd, k, v, ret}
}

func (c *chanMap) runCmd() {
	for {
		ci := <-c.cmdChan

		switch ci.cmd {
		case constCmdDelete:
			delete(c.m, ci.k)
		case constCmdGet:
			if ci.ret == nil {
				return
			}
			v, ok := c.m[ci.k]
			if ok {
				ci.ret <- v
			} else {
				ci.ret <- nil
			}
		case constCmdSet:
			c.m[ci.k] = ci.v
		}
	}
}
