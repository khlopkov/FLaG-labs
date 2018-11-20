package lab2

type State interface {
	AddTransission(symbol rune, state State)
	doTransition(symbol rune) State
}

type state struct {
	transissions map[rune]State
}

func NewState() State {
	return &state{
		transissions: make(map[rune]State),
	}
}

func (s *state) AddTransission(symbol rune, state State) {
	s.transissions[symbol] = state
}

func (s state) doTransition(symbol rune) State {
	return s.transissions[symbol]
}
