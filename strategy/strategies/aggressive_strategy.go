package strategies

import "fmt"

type AggressiveStrategy struct {

}

func (as AggressiveStrategy) TakeAction() {
	fmt.Println("I am aggressive")
}