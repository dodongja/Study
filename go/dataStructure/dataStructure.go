package main

import (
	"fmt"
)

/*func main() {
	v := list.New() //리스트 생성
	e4 := v.PushBack(4) // 리스트 뒤에 요소 추가
	e1 := v.PushFront(1) //리스트 앞에 요소 추가
	v.InsertBefore(3, e4)  // e4 요소 앞에 요소 삽입
	v.InsertAfter(2, e1) //e1 요소 뒤에 요소 삽입

	for e := v.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
			fmt.Print(e.Value, " ")
	}
}*/

//queue
/*
type Queue struct {
	v * list.List
}

func (q *Queue) Push(val interface{}){
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{ list.New()}
}

func main() {
	queue := NewQueue()

	for i:= 1; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v ->", v)
		v = queue.Pop()
	}
}*/

//map

func main() {
	m := make(map[string]string)
	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "청주시 상당구"

	fmt.Println(m["송하나"]) //서울시 강남구

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "이화랑")
	v, ok := m["이화랑"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}
}
