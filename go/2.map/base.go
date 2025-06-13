package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

func main() {

	//Maps are dictionaries:indexed by key,returing a value
	//Maps are passed by reference,no copying,updating ok
	//The type used for the key must have == & != defined(not slices,maps or funcs)
	//You can read from a nil map,but inserting will panic
	//Map also has a descriptor that points to internal hash table

	//var mp map[string]int        //nil,no storage
	m := make(map[string]string) //non-nil but empty

	//mp["and"] = 1 //PANIC-nil map

	m["Alen"] = "25"
	m["Rishav"] = "20"
	fmt.Println("map:", m)

	//You can read from nil map and it will return the default value
	fmt.Println("fetch:", m["Raj"])

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

	//Maps can't be compared to one another,maps can be compared only to nil as a special case
	rk := map[string]int{
		"apple":  5,
		"banana": 10,
		"cherry": 15,
	}
	var sk map[string]int

	//fails - map can only be compared to nil
	// if sk == rk {
	// }
	fmt.Println("is rk nil?", rk == nil)
	fmt.Println("is sk nil?", sk == nil)

	//Maps have a special two-result lookup function
	//The second variable tells you if the key was there
	// Access a non-existing key
	p := map[string]int{}
	a := p["the"] // returns 0
	fmt.Println("a =", a)

	// Access another non-existing key with comma-ok idiom
	b, ok := p["and"] // 0, false
	fmt.Println("b =", b, "ok =", ok)

	p["and"]++
	fmt.Println("p[\"and\"] =", p["and"])

	// Access again after modification
	c, ok := p["and"] // 1, true
	fmt.Println("c =", c, "ok =", ok)

	//Make nil useful *
	//Nil is a type of zero,it indicates absence of something

	//printing top 3 frequent words
	words := make(map[string]int)
	words["the"] = 4
	words["hello"] = 1
	words["world"] = 1
	words["I"] = 7
	words["am"] = 3
	type kv struct {
		key string
		val int
	}
	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}
	//read about function literal & closures
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})
	fmt.Println(ss[0:3])
}

/*output
map: map[Alen:25 Rishav:20]
fetch:
map: map[Rishav:20]
map: map[]
map[string]main.Student{"Alen":main.Student{Name:"Alen", Age:20}, "Rishav":main.Student{Name:"Rishav", Age:25}}
is rk nil? false
is sk nil? true
a = 0
b = 0 ok = false
p["and"] = 1
c = 1 ok = true
[{I 7} {the 4} {am 3}]
*/
