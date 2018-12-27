package main

import (
	"./factorypattern"
	"./proxypattern"
	"fmt"
	"./strategypattern"
)

/*工厂模式*/
func factory() {
	fmt.Println("工厂模式")
	factorypattern.Excute()
	/*
		测试类的实例化
	*/
	/*var a factorypattern.Animal = new(factorypattern.Dog)
	a.Cry()
	b := new(factorypattern.Dog)
	b.Ask()*/
	/*var cat *factorypattern.Cat = new (factorypattern.Cat)
	cat.Cry()*/
}

/*静态代理模式*/
func proxy() {
	fmt.Println("静态代理模式")
	proxypattern.Excute()
}

func strategy() {
	fmt.Println("策略模式")
	strategypattern.Excute()
}

func main(){
	factory()
	proxy()
	strategy()
}