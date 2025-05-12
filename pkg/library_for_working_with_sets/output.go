package library_for_working_with_sets

import (
	"fmt"
)

// PrintSetElements - Output the set in normal form to the console
func PrintSetElements(set *BitSet) {
	fmt.Println("set:")

	//by the condition the universe starts with 1, so here too with 1
	for ii := 1; ii < set.Capacity; ii++ {
		if set.Contains(set, ii) == true {
			fmt.Printf("%d ", ii)
		}
	}

	fmt.Println()
}
