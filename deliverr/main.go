package main

import (
	"fmt"

	"deliverr/allocator"
)

func main() {
	inputOrder := map[string]int{
		"apple":      1,
		"watermelon": 1,
		"avocado":    9,
		"hamburger":  1,
		"banana":     1,
		"orange":     1,
	}

	inventory := []allocator.InventoryDist{
		allocator.InventoryDist{
			Name: "owd",
			Inventory: []allocator.Good{
				allocator.Good{
					Name:   "apple",
					Number: 15,
				},
				allocator.Good{
					Name:   "pear",
					Number: 10,
				},
				allocator.Good{
					Name:   "grapes",
					Number: 5,
				},
				allocator.Good{
					Name:   "banana",
					Number: 5,
				},
			},
		}, //This one is cheaper than the second one
		allocator.InventoryDist{
			Name: "dm",
			Inventory: []allocator.Good{
				allocator.Good{
					Name:   "apple",
					Number: 3,
				},
				allocator.Good{
					Name:   "pear",
					Number: 8,
				},
				allocator.Good{
					Name:   "watermelon",
					Number: 40,
				},
				allocator.Good{
					Name:   "orange",
					Number: 10,
				},
			},
		},
		allocator.InventoryDist{
			Name: "rohanDistributor",
			Inventory: []allocator.Good{
				allocator.Good{
					Name:   "almonds",
					Number: 3,
				},
				allocator.Good{
					Name:   "hamburger",
					Number: 8,
				},
				allocator.Good{
					Name:   "orange",
					Number: 40,
				},
				allocator.Good{
					Name:   "pear",
					Number: 10,
				},
				allocator.Good{
					Name:   "avocado",
					Number: 12,
				},
			},
		},
	}

	distributors := allocator.InventoryAllocator(inputOrder, inventory)
	fmt.Println(distributors)
}
