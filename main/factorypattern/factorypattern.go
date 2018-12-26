package factorypattern

import "fmt"

/*定义接口*/
type Animal interface {
	Cry()
}

/*接口实现类1*/
type Dog struct {

}

/*接口实现类2*/
type Cat struct {


}

/*实现类2实现接口的方法*/
func (Cat) Cry() {
	fmt.Println("汪汪汪")
}

/*实现类1实现接口的方法*/
func (Dog) Cry() {
	fmt.Println("喵喵喵")
}

/*实现类自定义方法*/
func (Dog) Ask() {
	fmt.Println("汪汪汪")
}

/*工厂类*/
type AnimalFactory struct {

}

/*工厂类方法实现*/
func (AnimalFactory) getAnimalObj(animalType string) (Animal){
	var animal Animal
	switch animalType {
	case "cat":
		animal = new(Cat)
		break
	case "dog":
		animal = new(Dog)
		break
	default:
		panic("no this kind animal")
	}
	return animal
}

/*测试*/
func Excute(){
	new(AnimalFactory).getAnimalObj("cat").Cry()
	new(AnimalFactory).getAnimalObj("dog").Cry()
}