package stmt

import (
	"fmt"
	"github.com/viant/velty/internal/est"
	"github.com/viant/velty/internal/est/op"
	"reflect"
	"unsafe"
)

type directAppender struct {
	x *op.Operand
}

func (a *directAppender) appendString(state *est.State) unsafe.Pointer {
	ptr := state.Pointer(*a.x.Offset)
	state.Buffer.AppendString(*(*string)(ptr))
	return ptr
}

func (a *directAppender) appendInt(state *est.State) unsafe.Pointer {
	ptr := state.Pointer(*a.x.Offset)
	state.Buffer.AppendInt(*(*int)(ptr))
	return ptr

}

func (a *directAppender) appendBool(state *est.State) unsafe.Pointer {
	ptr := state.Pointer(*a.x.Offset)
	state.Buffer.AppendBool(*(*bool)(ptr))
	return ptr
}

func (a *directAppender) appendFloat64(state *est.State) unsafe.Pointer {
	ptr := state.Pointer(*a.x.Offset)
	state.Buffer.AppendFloat(*(*float64)(ptr))
	return ptr
}

func (a *directAppender) newAppendStringIndirect() est.Compute {
	upstream := op.Upstream(a.x.Sel)

	return func(state *est.State) unsafe.Pointer {
		ret := upstream(state)
		state.Buffer.AppendString(*(*string)(ret))
		return ret
	}
}

func (a *directAppender) newAppendIntIndirect() est.Compute {
	return func(state *est.State) unsafe.Pointer {
		ret := a.x.Exec(state)
		state.Buffer.AppendInt(*(*int)(ret))
		return ret
	}
}

func (a *directAppender) newAppendBoolIndirect() est.Compute {
	return func(state *est.State) unsafe.Pointer {
		ret := a.x.Exec(state)
		state.Buffer.AppendBool(*(*bool)(ret))
		return ret
	}
}

func (a *directAppender) newAppendFloatIndirect() est.Compute {
	return func(state *est.State) unsafe.Pointer {
		ret := a.x.Exec(state)
		state.Buffer.AppendFloat(*(*float64)(ret))
		return ret
	}
}

func Selector(expr *op.Expression) est.New {
	return func(control est.Control) (est.Compute, error) {
		x, err := expr.Operand(control)
		if err != nil {
			return nil, err
		}

		result := &directAppender{x: x}
		switch expr.Type.Kind() {
		case reflect.Int:
			if !x.Sel.Indirect {
				return result.appendInt, nil
			}
			return result.newAppendIntIndirect(), nil

		case reflect.String:
			if !x.Sel.Indirect {
				return result.appendString, nil
			}
			return result.newAppendStringIndirect(), nil

		case reflect.Bool:
			if !x.Sel.Indirect {
				return result.appendBool, nil
			}
			return result.newAppendBoolIndirect(), nil

		case reflect.Float64:
			if !x.Sel.Indirect {
				return result.appendFloat64, nil
			}
			return result.newAppendFloatIndirect(), nil
		}
		return nil, fmt.Errorf("unsupported append selector: %s", expr.Type.String())
	}
}
