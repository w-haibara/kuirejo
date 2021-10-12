package statemachine

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spyzhov/ajson"
)

type State interface {
	SetName(name string)
	SetID(id string)
	StateType() string
	String() string
	Transition(ctx context.Context, r *ajson.Node) (next string, w *ajson.Node, err error)
	SetLogger(l *logrus.Entry)
	Logger() *logrus.Entry
}
