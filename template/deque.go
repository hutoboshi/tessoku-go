package template

import "container/list"

// deque
type Deque struct {
	list *list.List
}

func NewDeque() *Deque {
	return &Deque{
		list: list.New(),
	}
}

func (d *Deque) PushFront(item interface{}) {
	d.list.PushFront(item)
}

func (d *Deque) PushBack(item interface{}) {
	d.list.PushBack(item)
}

func (d *Deque) PopFront() interface{} {
	if d.list.Len() == 0 {
		return nil
	}
	front := d.list.Front()
	d.list.Remove(front)
	return front.Value
}

func (d *Deque) PopBack() interface{} {
	if d.list.Len() == 0 {
		return nil
	}
	back := d.list.Back()
	d.list.Remove(back)
	return back.Value
}

func (d *Deque) Len() int {
	return d.list.Len()
}

/* dequeの使い方
func main() {
	deque := NewDeque()
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(3)

	fmt.Println("Deque length:", deque.Len()) // Output: Deque length: 3
	fmt.Println(deque.PopFront())             // Output: 3
	fmt.Println(deque.PopBack())              // Output: 2
	fmt.Println(deque.PopFront())             // Output: 1
	fmt.Println("Deque length:", deque.Len()) // Output: Deque length: 0
}
*/
