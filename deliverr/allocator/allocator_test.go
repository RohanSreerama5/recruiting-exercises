package allocator_test

import (
	"deliverr/allocator"
	"reflect"
	"testing"
)

func TestInventoryAllocator(t *testing.T) {
	test_cases := []struct {
		name           string
		inputOrder     map[string]int
		inventory      []allocator.InventoryDist
		expectedOutput []allocator.InventoryDist
	}{
		{
			name: "Happy Path",
			inputOrder: map[string]int{
				"apple":      1,
				"watermelon": 1,
				"banana":     1,
				"pear":       1,
				"orange":     1,
				"grapes":     1,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist{
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
			},
		},

		//insert new cases here

		{
			name: "Order can be shipped using one warehouse",
			inputOrder: map[string]int{
				"apple": 1,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist{
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
			},
		},
		{
			name: "Order can be shipped using multiple warehouses",
			inputOrder: map[string]int{
				"apple": 17,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist{
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
			},
		},
		{
			name: "Order cannot be shipped because there is not enough inventory",
			inputOrder: map[string]int{
				"apple": 100,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist(nil),
		},
		{
			name: "Asking for too many apples, but right amount of watermelon",
			inputOrder: map[string]int{
				"apple":      100,
				"watermelon": 3,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist(nil),
		},
		{
			name: "Reverse: asking for too many watermelon, but correct amount of apples",
			inputOrder: map[string]int{
				"apple":      1,
				"watermelon": 400,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist(nil),
		},
		{
			name: "Zero case: Input of 0 for 'good' count should be ignored",
			inputOrder: map[string]int{
				"apple":      0,
				"watermelon": 5,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist{
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
			},
		},
		{
			name: "Zero case (reverse): Input of 0 for 'good' count should be ignored",
			inputOrder: map[string]int{
				"apple":      5,
				"watermelon": 0, //testing to ensure results do not change because watermelon is coming from a different distributor
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist{
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
			},
		},
		{
			name: "Comprehensive zero case: Mix of zeroes",
			inputOrder: map[string]int{
				"apple":      0,
				"watermelon": 0, //testing to ensure results do not change because watermelon is coming from a different distributor
				"pear":       3,
				"grapes":     0,
				"orange":     1,
			},
			inventory: []allocator.InventoryDist{
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
			},
			expectedOutput: []allocator.InventoryDist{
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
			},
		},
	}

	//cases end here

	for _, test := range test_cases {

		output := allocator.InventoryAllocator(test.inputOrder, test.inventory)
		expectedOutput := test.expectedOutput

		if !reflect.DeepEqual(output, expectedOutput) { //complaining because of wrong ordering
			t.Errorf("\n Test Name: %v \n \n Inputted: %v \n \n Expected: %#v \n \n Got: %#v \n", test.name, test.inputOrder, test.expectedOutput, output)
			//t.Error("Test Failed: {} inputted, {} expected, received: {}", test.name, test.inputOrder, test.expectedOutput, output)
		}

	}
}
