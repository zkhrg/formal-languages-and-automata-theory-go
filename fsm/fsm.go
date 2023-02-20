package fsm

import (
	"fmt"
	"strings"
)

type Transition struct {
	start_state string
	values      []string
	end_state   string
}

func NewTransition(start_state string, values []string, end_state string) *Transition {
	t := Transition{
		start_state: start_state,
		values:      values,
		end_state:   end_state,
	}
	return &t
}

// структура, реализующая конечный автомат
// автомат может быть исключительно ДКА
type FiniteStateMachine struct {
	States        []string
	Initial_state string
	Current_state string
	Transitions   map[string]map[string]string
	Label         string
	Final_states  map[string]string
	Alphabet      []string
}

func NewFiniteStateMachine(states []string, initial_state string, transtions []Transition, label string, final_states map[string]string, alphabet []string) *FiniteStateMachine {
	fsm := FiniteStateMachine{
		States:        states,
		Initial_state: initial_state,
		Current_state: initial_state,
		Transitions:   set_transitions(transtions),
		Label:         label,
		Final_states:  final_states,
		Alphabet:      alphabet,
	}
	return &fsm
}

func (fsm *FiniteStateMachine) Run(element string) {
	fsm.Current_state = fsm.Transitions[fsm.Current_state][element]
	fmt.Printf("Пришел символ: %s | Автомат принял состояние: %s\n", element, fsm.Current_state)
}

func (fsm *FiniteStateMachine) Autorun(start_string, split_by string) {
	fmt.Printf("Автомат запущен \"%s\"\n", fsm.Label)
	list_of_values := strings.Split(start_string, "")

	for _, element := range list_of_values {
		fsm.Run(element)
	}

	value, exist := fsm.Final_states[fsm.Current_state]
	fmt.Println(exist)
	if exist {
		fmt.Printf("%s\n\n", value)
	} else {
		fmt.Println("Автомат не попал в конечное состояние")
	}
}

// функция выставления переходов в виде вложенных хэш-таблиц в экземпляр класса
// FiniteStateMachine
func set_transitions(transitions []Transition) map[string]map[string]string {
	b := map[string]map[string]string{}
	for _, transition := range transitions {
		for _, value := range transition.values {
			_, exist := b[transition.start_state]
			if exist {
				b[transition.start_state][value] = transition.end_state
			} else {
				b[transition.start_state] = map[string]string{value: transition.end_state}
			}
		}
	}
	return b
}

func Set_v_transitions(transitions []Transition) map[string][][]string {
	b := map[string][][]string{}
	for _, transition := range transitions {
		for _, value := range transition.values {
			_, exist := b[transition.start_state]
			if exist {
				b[transition.start_state] = append(b[transition.start_state], []string{value, transition.end_state})
			} else {
				b[transition.start_state] = [][]string{{value, transition.end_state}}
			}
		}
	}
	return b
}

// функция которая вернет нам ссылку уже готовый конечный автомат делимости на 2
func Fsm2() *FiniteStateMachine {
	alphabet := strings.Split("0123456789", "")
	fsm2 := NewFiniteStateMachine([]string{"q0", "q1"}, "q0", []Transition{
		*NewTransition("q0", []string{"0", "2", "4", "6", "8"}, "q0"),
		*NewTransition("q0", []string{"1", "3", "5", "7", "9"}, "q1"),
		*NewTransition("q1", []string{"0", "2", "4", "6", "8"}, "q0"),
		*NewTransition("q1", []string{"1", "3", "5", "7", "9"}, "q1")},
		"Делимость на 2",
		map[string]string{"q0": "Делится на 2!"},
		alphabet,
	)

	return fsm2
}
