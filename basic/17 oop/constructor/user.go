package main

type User struct {
	Name string //""表示未知
	Age  int    //-1表示未知
	Sex  byte   //1男，2女，3未知
}

func NewDefaultUser() *User {
	return &User{
		Name: "",
		Age:  -1,
		Sex:  3,
	}
}
