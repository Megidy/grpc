package types

import (
	"context"

	events "github.com/Megidy/eventsMicroservice/services/common/genproto/events/protobuf"
)

type EventStore interface {
	CreateEvent(context.Context, *events.Event) error
	GetEvent(context.Context, int) (*events.Event, error)
	UpdateEvent(context.Context, *events.Event) error
}
