package main
import (
	"fmt"
	"reflect"
)

type person struct {
	id int
	name string
}

func main() {

	p := person{}
	ch := make(chan int)
	refType(1)
	refType("tyty")
	refType(true)
	refType(ch)
	refType(p)

}

func refType(v interface{}) {

	actualType := reflect.TypeOf(v) //функция reflect.TypeOf() возвращает объект типа reflect.Type
					//reflect.Type, содержит фактический тип переданного аргумента interface{}
	specificKind := actualType.Kind() // метод Kind() объекта reflect.Type возвращает конкретный вид аргумента reflect.Kind
	fmt.Printf("%v has actualType = %v, specificKind = %v\n", v, actualType, specificKind)
}
