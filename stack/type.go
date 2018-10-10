package stack

type Stack struct {}

func(s *Stack) push() string {
	return "pushed"
}

func(s *Stack) pop() string {
	return "popped"
}