package prometheus

import "context"

type option struct {
	ctx context.Context
}

func WithContext(ctx context.Context) func(*option) {
	return func(o *option) {
		o.ctx = ctx
	}
}
