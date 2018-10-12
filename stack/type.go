package stack

type Stack struct {}

func(s *Stack) Push() string {
	return "pushed"
}

func(s *Stack) Pop() string {
	return "popped"
}