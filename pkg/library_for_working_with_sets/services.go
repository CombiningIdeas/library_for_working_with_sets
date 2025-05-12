package library_for_working_with_sets

import "math"

type BitSetInterface interface {
	BitsetAdd(setA *BitSet, element int)        // Adding one element
	BitsetAddMany(setA *BitSet, array []int)    // Adding multiple elements
	BitsetRemove(setA *BitSet, element int)     // Removing one element
	BitsetRemoveMany(setA *BitSet, array []int) // Deleting multiple elements

	Contains(setA *BitSet, element int) bool   // Checking if an element is in a set
	CompareSets(setA, setB *BitSet) bool       // Checking if sets are the same
	SetIsSubset(setA, setB *BitSet) bool       // Checking whether a set A is a subset of a set B
	SetIsStrongSubset(setA, setB *BitSet) bool // Checking whether a set A is a strict subset of a set B

	GetUnionSet(setA, setB *BitSet) BitSet               // Union of sets A and B
	GetIntersectionSet(setA, setB *BitSet) BitSet        // Intersection of sets A and B
	GetDifferenceSet(setA, setB *BitSet) BitSet          // Difference between sets A and B
	GetSymmetricDifferenceSet(setA, setB *BitSet) BitSet // Symmetric difference of sets A and B
}

// BitSet — a structure containing a bit set
type BitSet struct {
	Bits []uint64 // it is a dynamic array of bits (block of bits)
	// or a slice of bits of 64 bits each
	Size     int // current set size
	Capacity int // potential set size
}

// Constructor
func BitsetCreate(capacity int) BitSet {
	//this way, we find out the number of bit blocks we need:
	blocks := int(math.Ceil(float64(capacity) / 64.0))
	//math.Ceil(float64) - Round up to the nearest integer

	return BitSet{
		Bits:     make([]uint64, blocks),
		Size:     0,
		Capacity: capacity,
	}
}
