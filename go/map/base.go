package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]string)
	m["Alen"] = "25"
	m["Rishav"] = "20"
	fmt.Println("map:", m)

	delete(m, "Alen")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	students := make(map[string]Student)
	students["Rishav"] = Student{"Rishav", 25}
	students["Alen"] = Student{"Alen", 20}
	fmt.Printf("%#v\n", students)

	//go/map/base.go:21:3: invalid operation: cannot take address of students["Rishav"].Age (value of type int)
	//&students["Rishav"].Age = "30"
	//solution take map of map[string]*Student
}

/*output
map: map[Alen:25 Rishav:20]
map: map[Rishav:20]
map: map[]
map[string]main.Student{"Alen":main.Student{Name:"Alen", Age:20}, "Rishav":main.Student{Name:"Rishav", Age:25}}
*/
