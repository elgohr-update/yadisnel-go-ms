package pkger

import (
	"context"

	"github.com/yadisnel/go-ms/v2/config/source"
)

type pkgerPathKey struct{}

// WithPath sets the path to pkger
func WithPath(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, pkgerPathKey{}, p)
	}
}
