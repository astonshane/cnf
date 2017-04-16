package main

import (
	"cnf/clause"
	"cnf/literal"
	"fmt"
)

func main() {
	fmt.Printf("Hello World!\n")

	clause1 := clause.Clause{Symbol: "^"}
	clause1 = clause1.Append(literal.Literal{Name: "A"})
	clause1 = clause1.Append(literal.Literal{Name: "B"})

	clause2 := clause.Clause{Symbol: "~"}
	clause2 = clause2.Append(clause1)

	fmt.Println(clause1)
	fmt.Println(clause2)
}
