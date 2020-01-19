package main

import "fmt"

// Person .
type Person struct {
	Name string
	Age  int
}

// SetName .
func (p Person) SetName(newName string) {
	p.Name = newName
}

// SetAge .
func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

// IA .
type IA interface {
	Get() int
	Set()
}

// B .
type B struct {
	Age int
}

// Get .
func (b *B) Get() int {
	return b.Age
}

// Set .
func (b B) Set() {
	b.Age = 16
}

// GetSomething .
func GetSomething(ia IA) int {
	return 18
}
func main() {
	{
		p := Person{"zhangsan", 15}
		fmt.Printf("%+v\n", p) // {Name:zhangsan Age:15}
		p.SetName("lisi")
		fmt.Printf("%+v\n", p) // {Name:zhangsan Age:15}
		p.SetAge(16)
		fmt.Printf("%+v\n", p) // {Name:zhangsan Age:16}

		p1 := &Person{"zhangsan", 15}
		fmt.Printf("%+v\n", p1) // {Name:zhangsan Age:15}
		p1.SetName("lisi")
		fmt.Printf("%+v\n", p1) // {Name:zhangsan Age:15}
		p1.SetAge(16)
		fmt.Printf("%+v\n", p1) // {Name:zhangsan Age:16}
	}

	{
		b := B{Age: 19}
		br := GetSomething(&b)
		fmt.Println(br) // 18
		/*
			如果接口的实现方法的接收者是 地址和对象，那么可以用地址继承接口
		*/

		f := func(i interface{}) interface{} {
			v, ok := i.(int)
			if ok {
				return v
			}
			return "not int"

		}

		fmt.Printf("%+v\n", f(12))                          // 12
		fmt.Printf("%+v\n", f("zhangsan"))                  // not int
		fmt.Printf("%+v\n", f(Person{Name: "kk", Age: 18})) // not int

		// f1 := func(i interface{}) interface{} {
		// 	v := i.(int)
		// 	return v

		// }

		// fmt.Printf("%+v\n", f1(12))                          // 12
		// fmt.Printf("%+v\n", f1("zhangsan"))                  // panic
		// fmt.Printf("%+v\n", f1(Person{Name: "kk", Age: 18})) // panic

		f2 := func(i interface{}) interface{} {
			switch v := i.(type) {
			case int, string, Person:
				return v
			default:
				return "unkwon type"
			}

		}

		fmt.Printf("%+v\n", f2(12))                          // 12
		fmt.Printf("%+v\n", f2("zhangsan"))                  // zhangsan
		fmt.Printf("%+v\n", f2(Person{Name: "kk", Age: 18})) // {Name:kk Age:18}
	}

}
