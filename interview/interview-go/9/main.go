package main

type Student struct {
	name string
}

func main() {
	m := map[string]*Student{"people": &Student{"zhoujielun"}}
	m["people"].name = "wuyanzu"

	n := map[string]*Student{}
	n["hello"] = &Student{"shit"}
	n["hello"].name = "fuck"
}