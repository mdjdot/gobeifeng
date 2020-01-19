package main

import "encoding/json"
import "fmt"

// Person .
type Person struct {
	Name string `json:"user_name"`
	Age  int    `json:"user_age"`
}

func main() {
	arr := []string{"a", "b", "中文"}
	arrJSON, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println(arrJSON)         // [91 34 97 34 44 34 98 34 44 34 228 184 173 230 150 135 34 93]
	fmt.Println(string(arrJSON)) // ["a","b","中文"]

	p := Person{Name: "张天", Age: 18}
	pJSON, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pJSON)         // [123 34 117 115 101 114 95 110 97 109 101 34 58 34 229 188 160 229 164 169 34 44 34 117 115 101 114 95 97 103 101 34 58 49 56 125]
	fmt.Printf("%+v\n", string(pJSON)) // {"user_name":"张天","user_age":18}

	var pp Person
	err = json.Unmarshal(pJSON, &pp)
	fmt.Printf("%+v", pp) // {Name:张天 Age:18}
}
