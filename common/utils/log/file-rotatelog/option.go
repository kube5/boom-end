package rotatelogs

type Interface interface {
	Name() string
	Value() interface{}
}

type Opt struct {
	name  string
	value interface{}
}

func NewOpt(name string, value interface{}) *Opt {
	return &Opt{
		name:  name,
		value: value,
	}
}

func (o *Opt) Name() string {
	return o.name
}
func (o *Opt) Value() interface{} {
	return o.value
}
