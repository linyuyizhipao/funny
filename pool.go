package funny

import (
	"errors"
	"fmt"
	"time"
)

//需要被管理的资源对象实现 newObject，并通过它返回你需要的对象
type Object interface {
	newObject()Object
}

type Pool chan Object

//池里面复用几个资源
func NewPool(total int,o Object) *Pool {
	p := make(Pool, total)
	for i := 0; i < total; i++ {
		p <- o.newObject()
	}

	return &p
}

func (p Pool)Get(timeout time.Duration)(o Object,err error)  {
	if timeout < 0{
		err = fmt.Errorf("timeout error%v",timeout)
		return
	}

	select {
	case o= <- p:
		return o,nil
	case <-time.After(timeout):
		err = errors.New("get pool timeout")
		return

	}
}

func (p Pool)Close(o Object)  {
	p <- o
}
