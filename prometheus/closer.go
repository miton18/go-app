package prometheus

type Closer interface {
	Close() error
}
