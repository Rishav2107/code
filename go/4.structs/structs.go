package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
type → Go keyword to define a new type.
Employee → the name of the new type.
struct { ... } → the underlying type.
*/
type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}
type Employee2 struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

type Student struct {
	id  int
	age int
}

type Response struct {
	//first letter of variable should be capital to make it exported or accesible to other package
	//if it is lower case which makes it private, then marshal will ignore it while encoding
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

type Request struct {
	Page  int      `json:"page"`
	words []string `json:"words"`
}

// struct tags are not only for json but can be used for db,etc
func anotherFunction() {
	r := Response{1, []string{"up", "in", "out"}}

	//converting struct to json
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	fmt.Printf("%#v\n", r)

	//json to struct
	var r2 Response
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)

	req := Request{1, []string{"up", "in", "out"}}
	k, _ := json.Marshal(req)
	fmt.Println(string(k))
}

func main() {

	//1
	e := Employee{"Abhi", 10, nil, time.Now()}
	fmt.Printf("%T %+[1]v\n", e)

	e2 := Employee2{"Akash", 5, nil, time.Now()}

	c := make(map[string]*Employee)
	c["Rishav"] = &Employee{
		"Rishav",
		1,
		nil,
		time.Now(),
	}
	c["Rishav"].Number += 1
	c["Rahul"] = &Employee{"Rahul", 3, c["Rishav"], time.Now()}

	//fmt.Printf("%T %+v\n", e, e)
	fmt.Printf("%T %+[1]v\n", c["Rishav"])
	fmt.Printf("%T %+[1]v\n", c["Rahul"])

	//2
	//anonymous struct type
	var album1 = struct {
		title string
	}{
		"The White",
	}
	var album2 = struct {
		title string
	}{
		"The Black",
	}
	fmt.Println(album1, album2)

	//assignment
	/*
		Two struct types are compatible if
		1. field have same type and names
		2. in the same order
		3. and the same tags
	*/
	album1 = album2
	fmt.Println(album1, album2)

	//Fails: cannot use e (variable of struct type Employee) as Employee2 value in assignment
	//e2 = e
	e2 = Employee2(e) //first convert
	fmt.Println(e2)

	//a struct is comparable if all fields are comparable
	s1 := Student{1, 16}
	s2 := Student{3, 16}
	check := s1 == s2
	fmt.Println(check)

	//structs are passed by values unless a pointer is used

	//struct tags and json
	anotherFunction()

}

/*output
main.Employee {Name:Abhi Number:10 Boss:<nil> Hired:2025-05-17 17:16:52.309759 +0530 IST m=+0.000200668}
*main.Employee &{Name:Rishav Number:2 Boss:<nil> Hired:2025-05-17 17:16:52.310081 +0530 IST m=+0.000522418}
*main.Employee &{Name:Rahul Number:3 Boss:0x140000300c0 Hired:2025-05-17 17:16:52.310082 +0530 IST m=+0.000523251}
{The White} {The Black}
{The Black} {The Black}
{Abhi 10 <nil> 2025-05-17 17:16:52.309759 +0530 IST m=+0.000200668}
false
{"page":1,"words":["up","in","out"]}
main.Response{Page:1, Words:[]string{"up", "in", "out"}}
main.Response{Page:1, Words:[]string{"up", "in", "out"}}
{"page":1}
*/
