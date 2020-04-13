package pool

// Pool -
type Pool interface {
	Get() (interface{}, error)
	Put(interface{})
}
