package velty

import (
	"github.com/viant/velty/internal/ast/expr"
	"github.com/viant/velty/internal/est"
	"github.com/viant/velty/internal/est/op"
	"github.com/viant/velty/internal/est/stmt"
	"github.com/viant/xunsafe"
)

func (p *Planner) selectorExpr(selector *expr.Select) (*op.Expression, error) {
	var err error
	expr := &op.Expression{}
	expr.Selector, err = p.selector(selector)
	if err != nil {
		return nil, err
	}

	if expr.Selector == nil {
		id := selector.ID
		expr.Selector = op.NewSelector(id, selector.ID, nil, nil)
	}
	expr.Type = expr.Selector.Type
	return expr, nil
}

func (p *Planner) compileStmtSelector(actual *expr.Select) (est.New, error) {
	selExpr, err := p.selectorExpr(actual)
	if err != nil {
		return nil, err
	}

	p.updateFieldOffset(selExpr.Field, selExpr.Selector)
	return stmt.Selector(selExpr), nil
}

func (p *Planner) updateFieldOffset(field *xunsafe.Field, selector *op.Selector) {
	//If selector doesn't have a parent, it means that it was added to the p.Type as primitive field
	if selector.Parent == nil || selector.Indirect {
		return
	}

	temp := selector
	for temp.Parent != nil {
		temp = temp.Parent
	}
	field.Offset += temp.Offset
}
