package regex

// Action ...
type Action string

//State ...
type State struct {
	ID   string
	Next map[Action]*State
}

//Fragment ...
type Fragment struct {
	Start *State
	Out   *State
}
