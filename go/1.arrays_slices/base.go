package main

import "fmt"

func main() {

	//Arrays
	//Arrays are typed by size , which is fixed at compile time
	//Arrays are passed by value, thus elements are copied

	//all these are equivalent
	var a [3]int
	b := [3]int{1, 2, 3}
	c := [...]int{0, 0, 0} //sized by intializer

	var d [3]int
	d = b //elements copied
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	//Slice
	//Slice also has a descriptor like string which has the pointer, length and capacity
	//Slices have variable lenght and backed by some array
	//Slices are passed by reference, no copying, updating ok
	//Slices are not comparable
	//Has copy and append helpers

	var a1 []int         //nil,no storage
	var b1 = []int{1, 2} //initialized
	//a1 = append(a1, 1)   //append to nil,ok
	b1 = append(b1, 3) //[]int{1,2,3}
	a1 = b1            //overwrites descriptor of a1 and now a1 has the same descriptor as b1
	fmt.Println("a1: ", a1)
	fmt.Println("b1: ", b1)
	e1 := a1 //same storage(alias)
	fmt.Println("e1: ", e1)

	//creating a slice of length 0 and capacity 10
	students := make([]string, 0, 10)
	students = append(students, "Mohan")
	students = append(students, "Rahul")
	students = append(students, "Pragya")
	students = append(students, "Ashish")
	fmt.Printf("%#v\n", students)

	students[0] = "Akash"
	fmt.Printf("students: %#v\n", students)

	anotherStudents := students[0:2]
	fmt.Printf("anotherStudents: %#v\n", anotherStudents)

	//updating anotherStudents[1] will also reflect in students slice because they both share the same underlying array
	anotherStudents[1] = "Raghul"
	fmt.Printf("anotherStudents: %#v\n", anotherStudents)
	fmt.Printf("students: %#v\n", students)

	//you are not allowed to copy slice to array or array to slice
	//instead we can slice an array
	arr := [3]int{10, 20, 30}
	slice := arr[:] // converts [3]int to []int
	fmt.Printf("%#v\n", slice)

	t := []byte("string")
	fmt.Println(len(t), t)
	fmt.Println(t[2:])
}

/*output
a:  [0 0 0]
b:  [1 2 3]
c:  [0 0 0]
d:  [1 2 3]
a1:  [1 2 3]
b1:  [1 2 3]
[]string{"Mohan", "Rahul", "Pragya", "Ashish"}
students: []string{"Akash", "Rahul", "Pragya", "Ashish"}
anotherStudents: []string{"Akash", "Rahul"}
anotherStudents: []string{"Akash", "Raghul"}
students: []string{"Akash", "Raghul", "Pragya", "Ashish"}
[]int{10, 20, 30}
6 [115 116 114 105 110 103]
[114 105 110 103]
*/
