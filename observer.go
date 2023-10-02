package main

type Observer interface {
	update(string)
	getPhone() string
}
