package statemachine

import (
	"context"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spyzhov/ajson"
	"github.com/w-haibara/kuirejo/log"
)

type CommonState struct {
	Name           string      `json:"-"`
	StateMachineID string      `json:"-"`
	Type           string      `json:"Type"`
	Next           string      `json:"Next"`
	End            bool        `json:"End"`
	Comment        string      `json:"Comment"`
	InputPath      string      `json:"InputPath"`
	OutputPath     string      `json:"OutputPath"`
	logger         *log.Logger `json:"-"`
}

func (s *CommonState) SetName(name string) {
	s.Name = name
}

func (s *CommonState) SetID(id string) {
	s.StateMachineID = id
}

func (s *CommonState) StateType() string {
	if s == nil {
		return ""
	}

	return s.Type
}

func (s *CommonState) String() string {
	if s == nil {
		return ""
	}

	return pp.Sprintln(s)
}

func (s *CommonState) FilterInput(ctx context.Context, input *ajson.Node) (*ajson.Node, error) {
	node, err := filterByInputPath(input, s.InputPath)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (s *CommonState) FilterOutput(ctx context.Context, output *ajson.Node) (*ajson.Node, error) {
	node, err := filterByOutputPath(output, s.OutputPath)
	if err != nil {
		return nil, err
	}

	return node, nil
}

type TransitionFunc func(ctx context.Context, r *ajson.Node) (string, *ajson.Node, error)

func (s *CommonState) Transition(ctx context.Context, r *ajson.Node, fn TransitionFunc) (string, *ajson.Node, error) {
	if s == nil {
		return "", nil, nil
	}

	select {
	case <-ctx.Done():
		return "", nil, ErrStoppedStateMachine
	default:
	}

	next := s.Next
	if fn != nil {
		n, w, err := fn(ctx, r)
		if err != nil {
			return "", r, err
		}
		if strings.TrimSpace(n) != "" {
			next = n
		}
		if w != nil {
			r = w
		}
	}

	if s.End {
		return "", r, ErrEndStateMachine
	}

	if strings.TrimSpace(next) == "" {
		return "", nil, ErrNextStateIsBrank
	}

	return s.Next, r, nil
}

func (s *CommonState) SetLogger(v *log.Logger) {
	s.logger = v
}

func (s *CommonState) Logger(v logrus.Fields) *log.Logger {
	l := log.Logger{
		Entry: s.logger.WithFields(logrus.Fields{
			"name": s.Name,
			"type": s.Type,
			"next": s.Next,
			"end":  s.End,
			"line": log.Line(),
		}),
	}

	if v == nil {
		return &l
	}

	return &log.Logger{
		Entry: l.WithFields(v),
	}
}
