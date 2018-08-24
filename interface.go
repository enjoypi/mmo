package ext

type ParallelMap interface {
	Set(k, v interface{})
	Get(k interface{}) interface{}
	Delete(k interface{})
	Len() int
}
