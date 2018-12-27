package proxypattern

import "fmt"

/*接口*/
type ISeller interface {
	/*定义接口实现方法*/
	deal(int)
}
/*代理类*/
type ProxySeller struct {
	seller Seller
}
/*实体类*/
type Seller struct {

}

/*代理类构造方法*/
func newProxySeller(seller Seller) *ProxySeller {
	return &ProxySeller{
		seller:seller,
	}
}
/*代理类实现接口及业务逻辑
	中介代替买家完成工作
*/
func (proxySeller ProxySeller) deal(money int) {
	if proxySeller.seller.deal(money) {
		proxySeller.seller.connectMoney(money)
		proxySeller.seller.transfer()
		fmt.Println("交易成功")
	} else {
		fmt.Println("交易失败")
	}
	fmt.Println("交易结束")
}

/*实体类实现接口
	卖家交易函数
 */
func (Seller) deal(money int) bool{
	if money > 100 {
		return true
	} else {
		return false
	}
}
/*实体类
	卖家收钱函数
*/
func (Seller) connectMoney(money int) {
	fmt.Println("卖家收到钱")
}

/*实体类
	卖家过户函数
*/
func (Seller) transfer(){
	fmt.Println("卖家过户")
}

func Excute(){
	var proxySeller ISeller = newProxySeller(*new(Seller))
	proxySeller.deal(50)
	proxySeller.deal(200)
}