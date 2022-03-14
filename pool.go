package pool

import "sync/atomic"

type (
	Constructor func() (itf interface{})

	Destructor func(itf interface{})

	Pool struct {
		Constructor Constructor
		Destructor  Destructor
		Current     uint32
		Storage     chan interface{}
		Limit       uint32
		close       uint32
	}
)

func NewPool(constructor Constructor, destructor Destructor, limit uint32) (p *Pool) {
	p = &Pool{
		Constructor: constructor,
		Destructor:  destructor,
		Storage:     make(chan interface{}, limit),
		Limit:       limit,
	}

	return
}

func (p *Pool) Get() (itf interface{}) {
	if p.close == 1 {
		return
	}
	if p.Current == p.Limit {
		select {
		case itf = <-p.Storage:
		}
		return
	}
	select {
	case itf = <-p.Storage:
	default:
		itf = p.Constructor()
		atomic.AddUint32(&p.Current, 1)
	}

	return
}
func (p *Pool) Put(itf interface{}) {
	p.Storage <- itf
}

func (p *Pool) Close() (err error) {
	if !atomic.CompareAndSwapUint32(&p.close, 0, 1) {
		return
	}
	for i := 0; i < int(p.Current); i++ {
		itf := <-p.Storage
		p.Destructor(itf)
	}
	return
}
