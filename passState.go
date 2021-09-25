package main

import (
	"bytes"
	"strings"
)

type PassState struct {
	CommonState
	Result     string `json:"Result"`
	ResultPath string `json:"ResultPath"`
	Parameters string `json:"Parameters"`
}

func (s PassState) Transition(r, w *bytes.Buffer) (next string, err error) {
	if _, err := r.WriteTo(w); err != nil {
		return "", err
	}

	if s.End {
		return "", EndStateMachine
	}

	if strings.TrimSpace(s.Next) == "" {
		return "", NextStateIsBrank
	}

	return s.Next, nil
}
