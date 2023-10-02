package main

func main() {

	desirableItem := newItem("Adidas Sneakers")

	customer1 := &Customer{name: "Arsen", phone: "87073047764"}
	customer2 := &Customer{name: "Sultan", phone: "87778765677"}

	desirableItem.register(customer1)
	desirableItem.register(customer2)

	desirableItem.updateAvailability()
}
