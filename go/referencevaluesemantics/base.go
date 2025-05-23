package main

import (
	"fmt"
)

type User struct {
	Name  string
	Count int
}
type user struct {
	name string
}

func update(bookings []string, x string) []string {
	bookings = append(bookings, x)
	return bookings
}
func addTo(user *User) {
	user.Count += 1
}
func main() {

	//pointers - shared , not copied
	//values - copied , not shared

	//common use of pointers
	//some objects can't be copied safely(mutex)
	//some objects are too large to copy effeciently
	//some methods need to change the receiver (Read Methods)

	//Two variable range
	//The value returned by range is always a copy
	things := make([]string, 0, 5)
	things = append(things, "My")
	things = append(things, "name")
	things = append(things, "is")
	things = append(things, "Rishav")
	for i, thing := range things {
		//thing is a copy
		fmt.Println(i, " ", thing)
	}

	//One variable range
	for i := range things {
		fmt.Println(i, " ", things[i])
	}

	//slice are pass by reference
	//Anytime a function mutates a slice that's passed in, we must return a copy
	//because slice backing array may be reallocated to grow
	bookings := make([]string, 0, 10)
	bookings = update(bookings, "Goa Trip")
	bookings = update(bookings, "LA Trip")
	fmt.Printf("%#v\n", bookings)

	//we know that taking an address of a value in a  map will throw a compile error because
	//map rearranges itself dynamically refer: map/base.go

	//now for slice
	//keeping a pointer to an an element of a slice is risky, let's see
	//even though you updated Alice counter to 1, it is showing 0 because the underlying array was realloacted
	//when the threshold capacity was hit.
	users := make([]User, 0, 2)
	users = append(users, User{"Alice", 0})
	users = append(users, User{"Bob", 0})
	alice := &users[0]
	users = append(users, User{"Amy", 0})
	addTo(alice)
	fmt.Println(users)

	//converting slice of array to slice of slice
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	//a := make([][]byte, 0, 10)
	a := [][]byte{}

	for _, item := range items {
		a = append(a, item[:])
		//a = append(a, []byte{item[0], item[1]})
	}
	fmt.Println(items)
	fmt.Println(a)

	//Mutating Loop Variable
	//You might expect userPtrs to contain pointers to each User, but they all end up pointing to the same last u value.
	//fix : Starting with Go 1.22 , Allocates a new memory location for the loop variable on each iteration
	//So &u (even though it's still a copy) gets a unique address
	users2 := []user{{"Alice"}, {"Bob"}, {"Charlie"}}
	userPtrs := []*user{}

	for _, u := range users2 {
		userPtrs = append(userPtrs, &u)
	}
	for i, p := range userPtrs {
		fmt.Printf("userPtrs[%d]: address=%p, value=%+v\n", i, p, *p)
	}
}

/*output
0   My
1   name
2   is
3   Rishav
0   My
1   name
2   is
3   Rishav
[]string{"Goa Trip", "LA Trip"}
[{Alice 0} {Bob 0} {Amy 0}]
[[1 2] [3 4] [5 6]]
[[1 2] [3 4] [5 6]]
userPtrs[0]: address=0x140001380e0, value={name:Alice}
userPtrs[1]: address=0x140001380f0, value={name:Bob}
userPtrs[2]: address=0x14000138110, value={name:Charlie}
*/
