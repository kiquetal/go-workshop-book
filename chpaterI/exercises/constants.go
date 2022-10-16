package main

import "fmt"

const GlobalLimit = 100

const MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}
func cacheSet(key string, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func getBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)

}

func setBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

type User struct {
	name  string
	email string
	age   int
}

// NewUserV returns value ... ideally for a User we should not be
// returning value
func NewUserV(name, email string, age int) User {
	return User{name, email, age}
}

// NewUserP returns pointer ...
func NewUserP(name, email string, age int) *User {
	return &User{name, email, age}
}

// ChangeEmail ...
func (u *User) ChangeEmail(newEmail string) {
	u.email = newEmail
}

func main() {
	// with value type
	usr1 := NewUserV("frank", "frank@camero.com", 22)
	fmt.Println("Before change: ", usr1)
	usr1.ChangeEmail("frank@gmail.com")
	fmt.Println("After change: ", usr1)

	// with pointer type
	usr2 := NewUserP("john", "john@liliput.com", 22)
	fmt.Println("Before change: ", usr2)
	usr2.ChangeEmail("john@macabre.com")
	fmt.Println("After change: ", usr2)

}
