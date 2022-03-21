package main

import "fmt"

type Action struct {
	Name string
}
type Human struct {
	Surname string
	Name string
	Action
}

func (a *Action) SetName(name string) {
	a.Name = name
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func main() {

	var h Human = Human{
		Surname:   "PonPon",
		Name: "Li",
		Action: Action{
			Name: "engineer",
		},
	}
	//При встраивании одной структуры в другую метод наследуется.
	//Структура Human может иметь доступ ко всем методам структур,
	//которые в нее встроены.
	h.SetName("actor") //Вызов метода SetName у структуры Action изменит "engineer" на "actor"

	//Если дополнительно создать метод (h *Human)SetName, то вызов метода
	//будет изменять имя "PonPon", т.к. метод главной структуры Human
	//имеет больший приоритет, чем метод встроенной структуры Action.
	h.SetName("Jackie") //Теперь вызов метода SetName у структуры Human изменит имя "PonPon" на "Jackie", 
			     //если есть методы (a *Action)SetName и (h *Human)SetName.

	//Если требуется вызвать метод непосредственно у встроенной структуры, 
	//то обратиться к нему можно, используя полный селектор.
	h.Action.SetName("developer")
	fmt.Println(h)
}
