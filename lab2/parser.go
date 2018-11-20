package lab2

type Parser interface {
	Parse(str string) bool
}

type parser struct {
	initialState State
	finalState   State
}

func NewParser(initial State, final State) Parser {
	return &parser{
		initialState: initial,
		finalState:   final,
	}
}

func (p parser) Parse(str string) bool {
	var state = p.initialState
	if state == nil {
		return false
	}
	for _, ch := range str {
		state = state.doTransition(ch)
		if state == nil {
			return false
		}
	}
	if state == p.finalState {
		return true
	}
	return false
}
