package main

import "encoding/json"
import "fmt"

type People struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (p People) String() string {
	return fmt.Sprintf("{name: %s, age: %d}\n", p.Name, p.Age)
}

func main() {
	js := `{
		"name":"11",
		"age": 21
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}