package main


type Stack struct {
	items []interface{}

}

func NewStack() *Stack {
	s := Stack{}
	s.items = make([]interface{}, 0)
	return &s
}

func (s *Stack) push(item interface{}){
	s.items = append(s.items, item)
	// in append func we append a new item that pushed to s.items
}

// return an element of type anything so we use interfaces
func (s *Stack) pop() interface{}{
	if len(s.items) == 0{
		return nil
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack) peek() interface{}{
	if len(s.items) == 0{
		return nil
	}

	last := s.items[len(s.items)-1]
	// just locking what is top of it
	return last
}
