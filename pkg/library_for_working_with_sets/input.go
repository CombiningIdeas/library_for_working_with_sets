package library_for_working_with_sets

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadSetFromConsole - Reading a set from the console (in normal form)
func ReadSetFromConsole(capacity int) BitSet {
	result := BitsetCreate(capacity)

	fmt.Printf("Enter the elements of the set separated by a space (ranging from 1 to %d):\n",
		capacity-1)

	//Creates a new scanner to read from standard input (os.Stdin),
	//i.e. from the keyboard or redirected input.
	scanner := bufio.NewScanner(os.Stdin)
	//bufio.NewScanner, Used to conveniently read line by line.

	if scanner.Scan() == true { //scanner.Scan() - Attempts to read one line from the input.
		line := scanner.Text()             //the line variable stores the result of the read line
		inputArray := strings.Fields(line) //Splits the string line into arrays based on spaces and tabs
		for _, currentElement := range inputArray {
			num, err := strconv.Atoi(currentElement) //Converts a string to an integer
			if err == nil && num >= 1 && num <= capacity {
				result.BitsetAdd(&result, num)
			}
		}
	}

	return result
}
