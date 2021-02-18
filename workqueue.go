package api

import "context"

// WorkQueue define workqueue
type WorkQueue interface {
	Add(item interface{})
	Start(ctx context.Context) error
}
