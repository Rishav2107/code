package controlstatement

import (
	"fmt"
	"sort"
)

type kv struct {
	k string
	v string
}

// Named typing
type x int

func Get() int {
	return 1
}

// short declaration bug
func Read() int {
	var i int
	for {
		i := 5
		i += 1
		break
	}
	return i
}
func Display() {
	fmt.Println("The value of i is ", Read())
	for i := 0; i < 10; i++ {
		fmt.Printf("%d %d\n", i, i*i)
	}

	var s = []string{"ab", "bc", "cd"}
	//for array and slices
	// i is the index
	for i := range s {
		fmt.Println(s[i])
	}

	//this is ineffecient because value gets copied to v variable every time
	//first is index and second is value
	for _, v := range s {
		fmt.Println(v)
	}

	//maps
	var m = map[string]string{"Name": "Rishav", "City": "Pune"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	var ss []kv
	for key := range m {
		ss = append(ss, kv{key, m[key]})
	}
	//compares lexicographically
	sort.Slice(ss, func(i, j int) bool { return ss[i].k < ss[j].v })

	for _, s := range ss {
		fmt.Println(s.k, " ", s.v)
	}

	//infinite loop
	n := 0
	for {
		if n > 3 {
			break
		}
		fmt.Println(n)
		n += 1
	}

	//unlike java, in go switch case doesn't need to have a break
	a := Get()
	switch a {
	case 0, 1, 2:
		fmt.Println("underflow possible")
	case 3, 4, 5, 6, 7, 8: //do nothing
	default:
		fmt.Println("default warning")
	}

	var mm x // x is a defined type;base n
	y := 11  //y defaults to int
	// m = y   //Type Mismatch
	mm = x(y)
	fmt.Println(mm)
}
