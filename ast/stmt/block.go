package stmt

import "github.com/viant/velty/ast"

type Block struct {
	Stmt []ast.Statement
}

func (b *Block) AddStatement(statement ast.Statement) {
	b.Stmt = append(b.Stmt, statement)
}
