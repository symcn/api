package api

import (
	"time"

	"k8s.io/apimachinery/pkg/types"
)

// NeedRequeue need requeue
type NeedRequeue bool

// Requeue requeue last
// Done mark this step don't need requeue
const (
	Requeue NeedRequeue = true
	Done    NeedRequeue = false
)

// Reconciler interface, define Reconcile handle
type Reconciler interface {
	// Reconcile request name and namespace
	// returns requeue, after, error
	// 1. if error is not empty, will readd ratelimit queue
	// 2. if after > 0, will add queue after `after` time
	// 3. if requeue is true, readd ratelimit queue
	Reconcile(req types.NamespacedName) (requeue NeedRequeue, after time.Duration, err error)
}
