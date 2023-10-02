package main

import "fmt"

type Customer struct {
	name  string
	phone string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Item %s is available, dear %s\n", itemName, c.name)
}

func (c *Customer) getPhone() string {
	return c.phone
}
