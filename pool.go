package pool

type (
	Constructor func() (itf interface{})

	Destructor func(itf interface{})

	Pool struct {
		Constructor Constructor
		Destructor  Destructor
		Current     int
		Storage     chan interface{}
	}
)

func NewPool(constructor Constructor, destructor Destructor, limit int) (p *Pool) {
	p = &Pool{
		Constructor: constructor,
		Destructor:  destructor,
		Storage:     make(chan interface{}, limit),
	}

	return
}

func (p *Pool) Get() (itf interface{}) {
	select {
	case itf = <-p.Storage:
	default:
		itf = p.Constructor()
	}

	return
}
func (p *Pool) Put(itf interface{}) {
	p.Storage <- itf
}
