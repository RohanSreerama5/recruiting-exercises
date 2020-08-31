package allocator

type Good struct { //represents 1 good (ie. apple) and the count of apple
	Name   string
	Number int
}

type InventoryDist struct { //represents 1 distributor (ie. owd) and its name and what inventory the distributor has
	Name      string
	Inventory []Good
}

//finds duplicate distributors in slice and removes them
func findDuplicates(c InventoryDist, outputList []InventoryDist) []InventoryDist {
	myFlag := 0
	for _, z := range outputList {
		if z.Name == c.Name {
			myFlag = 1 //found that its already in the list
		}
	}
	if myFlag != 1 {
		outputList = append(outputList, c) //append this inventory distributor
	}

	return outputList
}

func InventoryAllocator(inputOrder map[string]int, inventory []InventoryDist) []InventoryDist { //[]InventoryDist is a slice of inventory distributors
	//[ { name: owd, inventory: { apple: 5, orange: 10 } }, { name: dm:, inventory: { banana: 5, orange: 10 } } ]
	//list of distributors/warehouses is presorted based on cost

	//First warehouse is less expensive to ship from than the second

	//gonna get a map of Good objects. We need to go thru and turn them each into an actual Good object

	var outputList []InventoryDist
	var emptyList []InventoryDist
	var flag int
	var zeroFlag int
	for a, b := range inputOrder { //loop thru each order
		flag = 0
		order := Good{
			Name:   a,
			Number: b,
		}
		for _, c := range inventory { //loop thru each inventory distributor
			if order.Number == 0 {
				zeroFlag = 1
				break
			}
			s := c.Inventory
			for _, v := range s { //looping thru slice (inventory of a single distributor)
				if order.Name == v.Name {
					if order.Number <= v.Number && order.Number > 0 {
						outputList = findDuplicates(c, outputList)
						flag = 1
						break //Done. break out of this current for loop and move onto next order in inputOrder
					} else if order.Number > v.Number {
						outputList = findDuplicates(c, outputList)
						//We want to empty this distributor and fill the order with the same good from the next guy
						order.Number = order.Number - v.Number //handles case where we need to get 10 apples form 1 guy and 10 from another guy
						break                                  //breaks out of this for loop and loops to next inventory distributor
					}
				}
			}

			if flag == 1 { //breaks and moves to next item in inputOrder list
				break
			}
		}

		if flag == 0 && zeroFlag != 1 { //if the order is too big (inventory cannot supply), just output an empty list
			outputList = emptyList
			break
		}
	}

	var temp InventoryDist
	for i := 0; i < len(outputList)-1; i++ { //sorting outputList lexicographically to ensure same response every time
		if outputList[i].Name > outputList[i+1].Name {
			temp = outputList[i]
			outputList[i] = outputList[i+1]
			outputList[i+1] = temp
		}

	}

	return outputList
}
