package lab2

import "testing"

func taskParser() Parser {
	state8 := NewState()
	state7 := NewState()
	state6 := NewState()
	state5 := NewState()
	state4 := NewState()
	state3 := NewState()
	state2 := NewState()
	state1 := NewState()
	state1.AddTransission('a', state1)
	state1.AddTransission('b', state2)
	state1.AddTransission('c', state3)
	state2.AddTransission('a', state4)
	state2.AddTransission('b', state2)
	state2.AddTransission('c', state3)
	state3.AddTransission('a', state7)
	state3.AddTransission('c', state6)
	state4.AddTransission('b', state5)
	state5.AddTransission('c', state6)
	state6.AddTransission('a', state7)
	state7.AddTransission('c', state8)
	return NewParser(state1, state8)
}

var testValues = []struct {
	str string
	res bool
}{
	{"cac", true},
	{"acac", true},
	{"aaaaacac", true},
	{"aaaabcac", true},
	{"bcac", true},
	{"abcac", true},
	{"abcac", true},
	{"ababcac", true},
	{"abccac", true},
	{"ccac", true},
	{"ababccac", false},
	{"abababcac", false},
	{"abcabcac", false},
	{"aaaaa", false},
	{"bbbb", false},
	{"", false},
	{"cacc", false},
}

func TestParse(t *testing.T) {
	parser := taskParser()
	for _, v := range testValues {
		if parser.Parse(v.str) != v.res {
			t.Errorf("%s should be %v", v.str, v.res)
		}
	}
}
