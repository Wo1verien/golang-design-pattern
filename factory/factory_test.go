/*
* @Author: Logan Chang
* @Date: 2021/6/21 5:26 下午
 */
package factory

import (
	"fmt"
	"testing"
)

func TestSelectOperator(t *testing.T) {
	op := SelectOperator("add")
	fmt.Println(op.Operate(1, 2))
}

//
func compute(factory OperatorFactory) int {
	op := factory.Create()
	return op.Operate(3, 1)
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)

	factory = PlusOperatorFactory{}
	if compute(factory) != 4 {
		t.Fatal("error with factory method pattern")
	}

	factory = MinusOperatorFactory{}
	if compute(factory) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
