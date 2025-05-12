package library_for_working_with_sets

// Adding one element
func (bs *BitSet) BitsetAdd(setA *BitSet, element int) {
	if element >= setA.Capacity || element <= 0 {
		return
	}
	block := element / 64
	bit := uint(element % 64)
	if (setA.Bits[block] & (1 << bit)) == 0 {
		setA.Bits[block] |= 1 << bit
		setA.Size++
	}
}

// Adding multiple elements
func (bs *BitSet) BitsetAddMany(setA *BitSet, array []int) {
	for _, element := range array {
		bs.BitsetAdd(setA, element)
	}
}

// Removing one element
func (bs *BitSet) BitsetRemove(setA *BitSet, element int) {
	if element >= setA.Capacity || element <= 0 {
		return
	}
	block := element / 64
	bit := uint(element % 64)
	if (setA.Bits[block] & (1 << bit)) != 0 {
		setA.Bits[block] &^= 1 << bit
		setA.Size--
	}
}

// Removing multiple elements
func (bs *BitSet) BitsetRemoveMany(setA *BitSet, array []int) {
	for _, element := range array {
		setA.BitsetRemove(setA, element)
	}
}

// Checking if an element is in a set
func (bs *BitSet) Contains(setA *BitSet, element int) bool {
	if element >= setA.Capacity || element < 0 {
		return false
	}
	block := element / 64
	bit := uint(element % 64)
	return (setA.Bits[block] & (1 << bit)) != 0
}

// Checking if sets are the same
func (bs *BitSet) CompareSets(setA, setB *BitSet) bool {
	if setA.Capacity != setB.Capacity {
		return false
	}
	for ii := range setA.Bits {
		if setA.Bits[ii] != setB.Bits[ii] {
			return false
		}
	}
	return true
}

// Checking whether set A is a subset of set B
func (bs *BitSet) SetIsSubset(setA, setB *BitSet) bool {
	for ii := range setA.Bits {
		if setA.Bits[ii]&^setB.Bits[ii] != 0 {
			return false
		}
	}
	return true
}

// Checking whether a set A is a strict subset of a set B
func (bs *BitSet) SetIsStrongSubset(setA, setB *BitSet) bool {
	return bs.SetIsSubset(setA, setB) && !bs.CompareSets(setA, setB)
}

// Union of sets A and B
func (bs *BitSet) GetUnionSet(setA, setB *BitSet) BitSet {
	capacity := max(setA.Capacity, setB.Capacity)
	result := BitsetCreate(capacity)
	for ii := range result.Bits {
		if ii < len(setA.Bits) {
			result.Bits[ii] |= setA.Bits[ii]
		}
		if ii < len(setB.Bits) {
			result.Bits[ii] |= setB.Bits[ii]
		}
	}
	return result
}

// Intersection of sets A and B
func (bs *BitSet) GetIntersectionSet(setA, setB *BitSet) BitSet {
	capacity := max(setA.Capacity, setB.Capacity)
	result := BitsetCreate(capacity)
	for ii := range result.Bits {
		if ii < len(setA.Bits) && ii < len(setB.Bits) {
			result.Bits[ii] = setA.Bits[ii] & setB.Bits[ii]
		}
	}
	return result
}

// Difference between sets A and B
func (bs *BitSet) GetDifferenceSet(setA, setB *BitSet) BitSet {
	capacity := max(setA.Capacity, setB.Capacity)
	result := BitsetCreate(capacity)
	for ii := range result.Bits {
		if ii < len(setA.Bits) {
			result.Bits[ii] = setA.Bits[ii]
		}
		if ii < len(setB.Bits) {
			result.Bits[ii] &^= setB.Bits[ii]
		}
	}
	return result
}

// Symmetric difference of sets A and B
func (bs *BitSet) GetSymmetricDifferenceSet(setA, setB *BitSet) BitSet {
	capacity := max(setA.Capacity, setB.Capacity)
	result := BitsetCreate(capacity)
	for ii := range result.Bits {
		var a, b uint64
		if ii < len(setA.Bits) {
			a = setA.Bits[ii]
		}
		if ii < len(setB.Bits) {
			b = setB.Bits[ii]
		}
		result.Bits[ii] = a ^ b
	}
	return result
}

// Complement (or negation)
func (bs *BitSet) GetComplementSet(setA *BitSet) BitSet {
	result := BitsetCreate(setA.Capacity)
	for ii := range setA.Bits {
		result.Bits[ii] = ^setA.Bits[ii]
	}
	maskLast := uint(setA.Capacity % 64)
	if maskLast != 0 {
		lastIndex := len(result.Bits) - 1
		mask := uint64(1<<maskLast) - 1
		result.Bits[lastIndex] &= mask
	}
	return result
}
