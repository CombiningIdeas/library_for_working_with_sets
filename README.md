# Library for working with sets
### Library for working with sets via bitwise operations on Golang :white_check_mark:

+ [Building and running the main program](#Building-and-running-the-main-program)
+ [Project file structure](#Project-file-structure)
+ [Project details](#Project-details)
+ [List of available functions](#List-of-available-functions)

---

## Building and running the main program

**Compiling a program** into a binary file:
```Linux
make
OR
make build
OR
make all
```

**Starting the program**:
```Linux
make run
```

**Code formatting command** for the entire Golang project:
```Linux
make fmt
```

**Reviews the code in the entire project**, if it finds errors or 
unclean code, it displays them in the console:
```Linux
make vet
```

**Testing**(but this library doesn't have them, 
so as not to burden the readability):
```Linux
make test
```

**Installing dependencies**, allows to automatically bring go.mod 
and go.sum files into line with the actual imports in the source: 
it adds missing dependencies, removes unused ones, and updates checksums:
```Linux
make deps
```

**Rebuild project**:
```Linux
make rebuild
```

**Clean the project** from the program :
```Linux
make clean
```

---

## Project file structure

```
/libraryForWorkingWithSets/
├── build/
│   └── package/ 
│       └── libraryForWorkingWithSets
├── cmd/
│   └── libraryForWorkingWithSets/     
│       └── main.go
├── internal/
│   └── app/                          
│       └── app.go
├── pkg/
│   └── library_for_working_with_sets/
│       ├── expressionChecking.go
│       ├── input.go
│       ├── output.go
│       ├── library.go
│       └── services.go
├── .gitignore
├── go.mod
│── LICENSE.md
│── Makefile
└── README.md
```

---

## Project details

Sets are represented as arrays of numbers, where each number is a 
bit mask in which each bit of the binary notation corresponds to a 
separate element of the set. This representation helps to use memory 
efficiently and speeds up operations with sets.

---

## List of available functions

| method                        | Description                                            |
|-------------------------------|--------------------------------------------------------|
| `BitsetCreate()`              | Creating a set (constructor)                           |
| `BitsetAdd()`                 | Adding one element                                     |
| `BitsetAddMany ()`            | Adding multiple elements                               |
| `BitsetRemove()`              | Removing one element                                   |
| `BitsetRemoveMany()`          | Removing multiple elements                             |
| `Contains()`                  | Checking if an element is in a set                     |
| `CompareSets()`               | Checking if sets are the same                          |
| `SetIsSubset`                 | Checking whether set A is a subset of set B            |
| `SetIsStrongSubset()`         | Checking whether a set A is a strict subset of a set B |
| `GetUnionSet()`               | Union of sets A and B                                  |
| `GetIntersectionSet()`        | Intersection of sets A and B                           |
| `GetDifferenceSet()`          | Difference between sets A and B                        |
| `GetSymmetricDifferenceSet()` | Symmetric difference of sets A and B                   |
| `GetComplementSet`            | Complement (or negation)                               |
| `PrintSetElements()`          | Output the set in normal form to the console           |
| `ReadSetFromConsole`          | Reading a set from the console in normal form          |