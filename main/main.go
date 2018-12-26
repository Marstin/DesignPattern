package main

import (
	"./factorypattern"
)

/*工厂模式*/
func factory() {
	factorypattern.Excute()
	/*
		测试类的实例化
	*/
	/*var a factorypattern.Animal = new(factorypattern.Dog)
	a.Cry()
	b := new(factorypattern.Dog)
	b.Ask()*/
	var cat * factorypattern.Cat = new (factorypattern.Cat)
	cat.Cry();
}

func main(){
	factory()
}