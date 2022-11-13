package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Class에 반, 학생수, 학생들 / Student에 번호, 점수
// 성적 입력 메서드 for문 반 인원수 만큼 돌면서 i = 번호 score= rand(0<= score <=100)
// 반 1등부터 10등 뽑기 번호:점수 뽑기
type Class struct {
	number          int
	numberOfStudent int
	students        []Student
}

type Student struct {
	num   int
	score int
}

// 애네들 리시버 * 포인터 변수로 못 받는가?
func (c Class) Len() int           { return len(c.students) }
func (c Class) Less(i, j int) bool { return c.students[i].score < c.students[j].score }
func (c Class) Swap(i, j int)      { c.students[i], c.students[j] = c.students[j], c.students[i] }

func (c *Class) setScore() *Class {
	var student Student
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < c.numberOfStudent; i++ {
		student.num = i

		student.score = rand.Intn(101)
		c.students = append(c.students, student)
	}

	return c
}

func main() {
	c := &Class{number: 1, numberOfStudent: 30}
	c.setScore()

	sort.Sort(sort.Reverse(Class(*c)))

	fmt.Println(c.students[:11])

}
