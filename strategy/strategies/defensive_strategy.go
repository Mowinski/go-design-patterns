package strategies

import "fmt"

type DefensiveStrategy struct {

}


func (ds DefensiveStrategy) TakeAction() {
	fmt.Println("I am defensive")
}