/*
* @Author: Logan Chang
* @Date: 2021/6/21 3:16 下午
 */
package simple_factory

import "fmt"

type Fruit interface {
	SayName()
}

type Apple struct{}

func (apple Apple) SayName() {
	fmt.Println("apple")
}

type Orange struct{}

func (orange Orange) SayName() {
	fmt.Println("orange")
}

func NewFruit(kind string) Fruit {
	switch kind {
	case "apple":
		return &Apple{}
	case "orange":
		return &Orange{}
	default:
		return nil
	}

}
