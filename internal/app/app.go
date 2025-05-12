package app

import (
	"fmt"
	"libraryForWorkingWithSets/pkg/library_for_working_with_sets"
)

func Tests() {
	fmt.Println("Testing the written code: ")

	var capacity int

	fmt.Print("Enter the value of the maximum element of the set, that is, capacity since our set \n" +
		"is a sorted non-repeating array of numbers, so we need to enter the maximum value of our set: ")
	_, err := fmt.Scan(&capacity)
	if err != nil {
		panic(err)
	}

	capacity++
	// Read the set from the console
	set := library_for_working_with_sets.ReadSetFromConsole(capacity)

	// Output the resulting set
	fmt.Println("Set read from console:")
	library_for_working_with_sets.PrintSetElements(&set)

	fmt.Printf("\n\n\n")
}

func Library() {
	fmt.Println("Main program: ")

	const lengthCapacity int = 10 + 1

	Universe := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Empty := []int{}
	InitialElementsSetA := []int{4, 8, 9}
	InitialElementsSetB := []int{1, 2, 4, 6, 8, 9}
	InitialElementsSetC := []int{4, 5, 7, 9}

	setA := library_for_working_with_sets.BitsetCreate(lengthCapacity)
	setB := library_for_working_with_sets.BitsetCreate(lengthCapacity)
	setC := library_for_working_with_sets.BitsetCreate(lengthCapacity)
	setUniverse := library_for_working_with_sets.BitsetCreate(lengthCapacity)
	setEmpty := library_for_working_with_sets.BitsetCreate(lengthCapacity)

	setA.BitsetAddMany(&setA, InitialElementsSetA)
	setB.BitsetAddMany(&setB, InitialElementsSetB)
	setC.BitsetAddMany(&setC, InitialElementsSetC)
	setUniverse.BitsetAddMany(&setUniverse, Universe)
	setEmpty.BitsetAddMany(&setEmpty, Empty)

	library_for_working_with_sets.PrintSetElements(&setA)
	library_for_working_with_sets.PrintSetElements(&setB)
	library_for_working_with_sets.PrintSetElements(&setC)
	library_for_working_with_sets.PrintSetElements(&setUniverse)
	library_for_working_with_sets.PrintSetElements(&setEmpty)

	fmt.Println("\n\nFor example, we work with the expression: A − (X ∩ B) ∪ (C ∩ ¬X) = (X − A) ∩ (X − B) − C")

	fmt.Println("Let's calculate the value: ϕ^∅ ")

	op1 := setC.GetIntersectionSet(&setC, &setUniverse)
	op2 := setB.GetIntersectionSet(&setB, &setEmpty)
	op3 := setA.GetDifferenceSet(&setA, &op2)
	op4 := op3.GetUnionSet(&op1, &op3)
	op5 := setEmpty.GetDifferenceSet(&setEmpty, &setA)
	op6 := setEmpty.GetDifferenceSet(&setEmpty, &setB)
	op7 := op6.GetIntersectionSet(&op5, &op6)
	op8 := op7.GetDifferenceSet(&op7, &setC)
	fi0 := op8.GetSymmetricDifferenceSet(&op4, &op8)

	fmt.Print("We get ϕ^∅: ")

	library_for_working_with_sets.PrintSetElements(&fi0)

	fmt.Println("\n\nNow let's calculate the value: ¬ϕ^U ")

	opU1 := setC.GetIntersectionSet(&setC, &setEmpty)
	opU2 := setB.GetIntersectionSet(&setB, &setUniverse)
	opU3 := setA.GetDifferenceSet(&setA, &opU2)
	opU4 := opU3.GetUnionSet(&opU1, &opU3)
	opU5 := setUniverse.GetDifferenceSet(&setUniverse, &setA)
	opU6 := setUniverse.GetDifferenceSet(&setUniverse, &setB)
	opU7 := opU6.GetIntersectionSet(&opU5, &opU6)
	opU8 := opU7.GetDifferenceSet(&opU7, &setC)
	fiU := opU8.GetSymmetricDifferenceSet(&opU4, &opU8)
	notFiU := fiU.GetComplementSet(&fiU)

	fmt.Print("We get ¬ϕ^U ")

	library_for_working_with_sets.PrintSetElements(&notFiU)

}

// TestingExpressions using unknown sets
func TestingExpressions() {
	fmt.Println("Testing work with expressions: ")

	const lengthCapacity int = 10 + 1
	found := false // Flag to track whether a suitable set X was found

	//From these combinations of sets of different elements, you can choose any
	//Among all the variations of sets A, B and C there are different ones,
	//a small part of them is presented below
	//InitialElementsSetA := []int{}
	//InitialElementsSetB := []int{1, 4, 5}
	//InitialElementsSetC := []int{1, 2, 3, 4}

	//	InitialElementsSetA := []int{}
	//	InitialElementsSetB := []int{1, 2, 3, 4, 5}
	//	InitialElementsSetC := []int{2, 3}

	//InitialElementsSetA := []int{3}
	//InitialElementsSetB := []int{1, 2, 3, 5}
	//InitialElementsSetC := []int{1, 3, 4}

	//InitialElementsSetA := []int{2}
	//InitialElementsSetB := []int{2, 3, 4}
	//InitialElementsSetC := []int{3}

	InitialElementsSetA := []int{1, 2, 3}
	InitialElementsSetB := []int{1, 2, 3, 5}
	InitialElementsSetC := []int{1, 4}

	setA := library_for_working_with_sets.BitsetCreate(lengthCapacity)
	setB := library_for_working_with_sets.BitsetCreate(lengthCapacity)
	setC := library_for_working_with_sets.BitsetCreate(lengthCapacity)

	setA.BitsetAddMany(&setA, InitialElementsSetA)
	setB.BitsetAddMany(&setB, InitialElementsSetB)
	setC.BitsetAddMany(&setC, InitialElementsSetC)

	library_for_working_with_sets.PrintSetElements(&setA)
	library_for_working_with_sets.PrintSetElements(&setB)
	library_for_working_with_sets.PrintSetElements(&setC)

	// Enumeration of all possible subsets of the universal set U (from 1 to 10)
	for num := 0; num < 1024; num++ { // 2^10 = 1024 possible subsets
		setX := library_for_working_with_sets.BitsetCreate(lengthCapacity)
		for x := 1; x <= 10; x++ {
			if num&(1<<(x-1)) != 0 {
				setX.BitsetAdd(&setX, x)
			}
		}

		left := library_for_working_with_sets.LeftPart(&setA, &setB, &setC, &setX)
		right := library_for_working_with_sets.RightPart(&setA, &setB, &setC, &setX)

		if left.CompareSets(left, right) == true {
			fmt.Println("\nA set X has been found that satisfies the expression:")
			library_for_working_with_sets.PrintSetElements(&setX)
			fmt.Print("Get the left side of the expression  A − (X ∩ B) ∪ (C ∩ ¬X) ")
			library_for_working_with_sets.PrintSetElements(left)
			fmt.Print("Get the right side of the expression  (X − A) ∩ (X − B) − C ")
			library_for_working_with_sets.PrintSetElements(right)
			found = true
			// if in both cases left and right nothing is output, but setX is output,
			// then this means that both parts are equal to the empty set
		}
	}

	//This check will work if none of the variants of setX satisfy the original equality
	if !found {
		fmt.Println("\nThere is no set X that satisfies the equality:")
		fmt.Println("A − (X ∩ B) ∪ (C ∩ ¬X) = (X − A) ∩ (X − B) − C")
	}

}
