package library_for_working_with_sets

// You can write the expression you need in such a file yourself and pull
// all the other methods from my library on mt github account.

// This function is the left side of the expression: A − (X ∩ B) ∪ (C ∩ ¬X)
func LeftPart(setA, setB, setC, setX *BitSet) *BitSet {
	NotX := setX.GetComplementSet(setX)         // ¬X
	op1 := setC.GetIntersectionSet(setC, &NotX) // C ∩ ¬X
	op2 := setB.GetIntersectionSet(setB, setX)  // X ∩ B
	op3 := setA.GetDifferenceSet(setA, &op2)    // A - (X ∩ B)
	result := op3.GetUnionSet(&op1, &op3)       // A − (X ∩ B) ∪ (C ∩ ¬X)

	return &result
}

// This function is the right side of the expression: (X − A) ∩ (X − B) − C
func RightPart(setA, setB, setC, setX *BitSet) *BitSet {
	op1 := setX.GetDifferenceSet(setX, setA)   // X - A
	op2 := setX.GetDifferenceSet(setX, setB)   // X - B
	op3 := op1.GetIntersectionSet(&op2, &op1)  // (X - A) ∩ (X - B)
	result := op3.GetDifferenceSet(&op3, setC) // (X − A) ∩ (X − B) - C

	return &result
}
