package errors

import "context"

// Localizer allows localized errors strings to be applied in a standardized manner
type Localizer interface {
	// Localize the error message
	Localize(ctx context.Context, err error) error
}

// LocalizerFunc implements Localizer
type LocalizerFunc func(ctx context.Context, err error) error

// Localize the error message
func (fn LocalizerFunc) Localize(ctx context.Context, err error) error {
	return fn(ctx, err)
}
