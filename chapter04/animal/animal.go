package animal

type Animal struct {
	name string
	age  int
}

func NewAnimal(name string, a int) Animal {
	return Animal{name: name, age: a}
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.name
}
