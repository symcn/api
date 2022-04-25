package api

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/types"
)

// NeedRequeue need requeue
type NeedRequeue bool

// Event event type
type Event int

// Requeue requeue last
// Done mark this step don't need requeue
const (
	Requeue NeedRequeue = true
	Done    NeedRequeue = false
)

const (
	AddEvent Event = iota
	UpdateEvent
	DeleteEvent
)

// Reconciler interface, define Reconcile handle
type Reconciler interface {
	// Reconcile request name and namespace
	// returns requeue, after, error
	// 1. if error is not empty, will readd ratelimit queue
	// 2. if after > 0, will add queue after `after` time
	// 3. if requeue is true, readd ratelimit queue
	Reconcile(ctx context.Context, req types.NamespacedName) (requeue NeedRequeue, after time.Duration, err error)
}

// WrapNamespacedName wrap standard namespacedname with queue name
type WrapNamespacedName struct {
	types.NamespacedName

	// queue name
	QName string
}

// WrapReconciler interface, define Reconcile handle
type WrapReconciler interface {
	// Reconcile request name and namespace
	// returns requeue, after, error
	// 1. if error is not empty, will readd ratelimit queue
	// 2. if after > 0, will add queue after `after` time
	// 3. if requeue is true, readd ratelimit queue
	Reconcile(ctx context.Context, req WrapNamespacedName) (requeue NeedRequeue, after time.Duration, err error)
}

type EventRequest struct {
	EventType Event

	OldResource interface{}
	NewResource interface{}
}

// EnhanceResourceEventReconciler interface, warpper an ResourceEventHandler
// qname is queue name, obj raw resource which add watched.
// returns requeue, after, error
// 1. if error is not empty, will readd ratelimit queue
// 2. if after > 0, will add queue after `after` time
// 3. if requeue is true, readd ratelimit queue
type EventReonciler interface {
	// OnAdd is called when an object is added.
	OnAdd(ctx context.Context, qname string, obj interface{}) (requeue NeedRequeue, after time.Duration, err error)

	// OnUpdate is called when an object is modified. Note that oldObj is the
	// last known state of the object-- it is possible that several changes
	// were combined together, so you can't use this to see every single
	// change. OnUpdate is also called when a re-list happens, and it will
	// get called even if nothing changed. This is useful for periodically
	// evaluating or syncing something.
	OnUpdate(ctx context.Context, qname string, oldObj, newObj interface{}) (requeue NeedRequeue, after time.Duration, err error)

	// OnDelete will get the final state of the item if it is known, otherwise
	// it will get an object of type DeletedFinalStateUnknown. This can
	// happen if the watch is closed and misses the delete event and we don't
	// notice the deletion until the subsequent re-list.
	OnDelete(ctx context.Context, qname string, obj interface{}) (requeue NeedRequeue, after time.Duration, err error)
}
