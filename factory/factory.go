/*
* @Author: Logan Chang
* @Date: 2021/6/21 3:57 下午
 */
package factory

//Operator 是被封装的实际类接口
type Operator interface {
	Operate(a, b int) int
}

//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

//PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{}
}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct{}

func (p PlusOperator) Operate(a, b int) int {
	return a + b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{}
}

//MinusOperator Operator 的实际减法实现
type MinusOperator struct{}

func (m MinusOperator) Operate(a, b int) int {
	return a - b
}

func SelectOperator(op string) Operator {
	switch op {
	case "add":
		return PlusOperatorFactory{}.Create()
	case "minus":
		return MinusOperatorFactory{}.Create()
	default:
		return nil
	}
}
