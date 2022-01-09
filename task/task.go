package task

import (
	"context"
	"fmt"

	"github.com/w-haibara/kakemoti/task/fn"
)

type (
	Fn    func(context.Context, string, fn.Obj) (fn.Obj, string, error)
	FnMap map[string]Fn
)

var fnMap FnMap

func init() {
	fnMap = make(FnMap)
	RegisterDefault()
}

func RegisterDefault() {
	Register("script", fn.DoScriptTask)
}

func Register(name string, fn Fn) {
	fnMap[name] = fn
}

func Do(ctx context.Context, resourceType, resoucePath string, input interface{}) (interface{}, string, error) {
	f, ok := fnMap[resourceType]
	if !ok {
		return nil, "", fmt.Errorf("invalid resouce type: %s", resourceType)
	}

	in, ok := input.(map[string]interface{})
	if !ok {
		return nil, "", fmt.Errorf("invalid input type: %#v", input)
	}

	out, stateserr, err := f(ctx, resoucePath, in)
	if stateserr != "" || err != nil {
		return nil, stateserr, fmt.Errorf("fn() failed: %v", err)
	}

	return out, "", nil
}
