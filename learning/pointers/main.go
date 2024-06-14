package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
	fmt.Println("-----------------")

	user := User{
		email:    "deepraj@mail.com",
		username: "deepraj",
		age:      23,
	}
	user.UpdateEmail("foo@me.com")
	fmt.Println(user.Email())
}

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u User) updateEmail(email string) {
	u.email = email
}


/// changes the state (with pointer)
func (u *User) UpdateEmail(email string) {
	u.email = email
}