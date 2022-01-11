package fn

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func DoStatesFormat(ctx context.Context, args []interface{}) (interface{}, error) {
	var ErrStatesFormatFailed = errors.New("DoStatesFormat() failed")
	const (
		str1 = "{}"
		str2 = "\\{}"
	)

	if len(args) < 2 {
		return nil, ErrStatesFormatFailed
	}

	f, ok := args[0].(string)
	if !ok {
		return nil, ErrStatesFormatFailed
	}

	if strings.Count(f, str1)-strings.Count(f, str2) != len(args)-1 {
		return nil, ErrStatesFormatFailed
	}

	result := f
	for _, arg := range args[1:] {
		switch reflect.ValueOf(arg).Kind() {
		case reflect.Map, reflect.Struct:
			return nil, ErrStatesFormatFailed
		}

		str := result
		n := 0
		for {
			n1 := strings.Index(str, str2)
			n2 := strings.Index(str, str1)
			if n1 < 0 || n2 < n1 {
				break
			}
			n = n1 + 3
			str = str[n:]
		}

		result = result[:n] + strings.Replace(result[n:], str1, fmt.Sprint(arg), 1)
	}

	return result, nil
}
