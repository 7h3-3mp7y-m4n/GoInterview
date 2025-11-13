package stack

import "fmt"

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	n := len(*s)
	if n == 0 {
		fmt.Println("sorry bro stack is already empty")
	}
	top := (*s)[n-1]
	*s = (*s)[:n-1]
	return top
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func main() {
	var st Stack
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	fmt.Println(st.Top())
	fmt.Println(st.Empty())

}
