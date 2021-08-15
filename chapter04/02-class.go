package main

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

func NewStudent(id uint, name string, score float64) *Student {
	return &Student{id: id, name: name, score: score}
}

func (s Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}

func main() {
	student := NewStudent(1, "测试一下", 100)
	fmt.Println(student)
	student.SetName("测试一下1")
	fmt.Println(student)
}
