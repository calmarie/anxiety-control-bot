package state

type State string

var StateMachine = make(map[int64][]State)
