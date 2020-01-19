package main

import "fmt"

// Person .
type Person struct {
	Name string
	Age  int
}

// Student .
type Student struct {
	Person
	Age   int
	Skill string
}

func main() {
	student := Student{Person{"zhangsan", 25}, 15, ""}
	fmt.Println(student)             // {{zhangsan 25} 15 }
	fmt.Printf("%v\n", student)      // {{zhangsan 25} 15 }
	fmt.Printf("%+v\n", student)     // {Person:{Name:zhangsan Age:25} Age:15 Skill:}
	fmt.Println(student.Age)         // 15
	fmt.Println(student.Person.Age)  // 25
	fmt.Println(student.Name)        // zhangsan
	fmt.Println(student.Person.Name) // zhangsan
}
