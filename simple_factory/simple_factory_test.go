/*
* @Author: Logan Chang
* @Date: 2021/6/21 3:36 下午
 */
package simple_factory

import "testing"

func TestApple_SayName(t *testing.T) {
	appleInstance := NewFruit("apple")
	appleInstance.SayName()
}
