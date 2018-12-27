package strategypattern

import "fmt"
/*实体类接口*/
type Staplefood interface {
	Eat()
}

type RiceStaplefood struct {

}

type NoodleStaplefood struct {

}
/*策略类*/
type EatContext struct {
	staplefood Staplefood
}

type BreadStaplefood struct {

}

func (RiceStaplefood) Eat(){
	fmt.Println("吃米饭")
}

func (NoodleStaplefood) Eat() {
	fmt.Println("吃面条")
}

func (BreadStaplefood) Eat() {
	fmt.Println("吃面包")
}
/*策略类操作方法*/
func (context EatContext) EatFood(){
	context.staplefood.Eat()
}
/*策略类构造函数*/
func NewEatContext(staplefood Staplefood) *EatContext{
	return &EatContext{
		staplefood:staplefood,
	}
}

func Excute(){
	eat := NewEatContext(new(RiceStaplefood))
	eat.EatFood()
	eat = NewEatContext(new(NoodleStaplefood))
	eat.EatFood()
	eat = NewEatContext(new(BreadStaplefood))
	eat.EatFood()
}