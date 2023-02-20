package main

import (
	"fmt"
	//"formal-languages-and-automata-theory-go/fsm"
	// "formal-languages-and-automata-theory-go/fsm"
	"formal-languages-and-automata-theory-go/fsm"
)

func main() {

	var number_to_check string
	fmt.Scanf("%s", &number_to_check)

	fsm2 := fsm.Fsm2()
	fsm2.Autorun(number_to_check, "")
}
