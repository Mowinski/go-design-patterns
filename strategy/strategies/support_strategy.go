package strategies

import "fmt"

type SupportStrategy struct {

}

func (ss SupportStrategy) TakeAction() {
	fmt.Println("I am support")
}