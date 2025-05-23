package main

import "fmt"

func main() {

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

	arr := [3]int{10, 20, 30}
	slice := arr[:] // converts [3]int to []int
	fmt.Printf("%#v\n", slice)
}

/*output
[]string{"Mohan", "Rahul", "Pragya", "Ashish"}
students: []string{"Akash", "Rahul", "Pragya", "Ashish"}
antoherStudents: []string{"Akash", "Rahul"}
anotherStudents: []string{"Akash", "Raghul"}
students: []string{"Akash", "Raghul", "Pragya", "Ashish"}
[]int{10, 20, 30}
*/
