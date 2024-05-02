package main


import (
	"fmt"
	"testing"
)


func TestStack(t *testing.T){
	s := NewStack()
	if s.items == nil{
		t.Error("expected stack to be initialized")
	}
}


func TestPush(t *testing.T){
	s  := NewStack()
	s.push(1)
	s.push(2)
	s.push(3)

	actualTop := s.peek()
	if actualTop != 3{
		t.Error("Expect last number in stack lifo")
	}else {
		fmt.Println("Done ya man")
	}
}


func TestPop(t *testing.T){
	s := NewStack()

	s.push(1)
	s.push(2)
	s.push(3)

	last := s.pop()
	if last != 3 {
		t.Error("expect 3 here")
	}
}


func TestPeekEmptyStack(t *testing.T){
	s := NewStack()

	last := s.peek()
	if last != nil {
		t.Error("expect last item in empty stack to be nil")
	}
}

func TestPopEmpty(t *testing.T)  {
	s := NewStack()

	last := s.pop()
	if last != nil{
		t.Error("expect last item in empty stack to be nil")
	}


}
