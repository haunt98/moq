package withresets

import "context"

// Reset is a reset.
type Reset struct {
	ID      string
	Name    string
	Company string
	Website string
}

// ResetStore stores resets.
type ResetStore interface {
	Get(ctx context.Context, id string) (*Reset, error)
	Create(ctx context.Context, person *Reset, confirm bool) error
	ClearCache(id string)
}

type ResetStoreGeneric[T, S any] interface {
	Get(ctx context.Context, id T) (string, error)
	Create(ctx context.Context, id T, value S) error
}
