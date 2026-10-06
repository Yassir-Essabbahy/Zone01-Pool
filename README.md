# Zone01 Piscine — Algorithms, Data Structures & Low-Level Go

[![Language](https://img.shields.io/badge/Language-Go%201.25+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![School](https://img.shields.io/badge/Campus-Zone01%20Oujda-7B2CBF?style=for-the-badge)](https://zone01oujda.ma/)
[![Framework](https://img.shields.io/badge/01--Edu-z01-brightgreen?style=for-the-badge)](https://github.com/01-edu/z01)
[![Pedagogy](https://img.shields.io/badge/Method-Peer--to--Peer-orange?style=for-the-badge)](https://01-edu.org/)

A curated repository of 60+ algorithmic solutions, low-level data structures, and command-line utilities developed during the intensive 4-week **Piscine at Zone01 Oujda** (part of the global **01Edu** network).

All implementations are coded in **Go (Golang)** strictly following low-level computational constraints—implementing memory structures, sorting routines, and pointer mechanics from first principles without standard library shortcuts.

---

## 🎯 About the Zone01 Piscine

Zone01 operates on the revolutionary **01Edu / Ecole 42** peer-learning model:
* **Zero Teachers, 100% Peer-to-Peer**: Automated test suites evaluate memory efficiency, edge-case resilience, and algorithmic complexity.
* **First-Principles Thinking**: High-level convenience methods (e.g. `strings.Split`, `sort.Slice`) are strictly prohibited in core modules; solutions must be implemented from memory allocations, byte slices, and pointer manipulation.
* **Intensive Immersion**: Over 200 hours of algorithmic problem solving covering recursion, pointers, linked lists, bitwise operations, and UNIX command-line utilities.

---

## 🧩 Categorized Syllabus & Implementations

### 1. Data Structures: Linked Lists
Implementation of custom singly-linked lists from scratch without external dependencies:

| File | Concept / Functionality |
| :--- | :--- |
| `listpushback.go` | Appends a new node to the tail of the linked list |
| `listpushfront.go` | Prepends a new node to the head of the linked list |
| `listreverse.go` | In-place reversal of node pointers without allocating extra memory |
| `listsize.go` | Counts total elements in $O(n)$ time |
| `listlast.go` | Retrieves pointer to the terminal node |
| `listclear.go` | Clears all references to enable garbage collection |
| `listat.go` | Returns the $n$-th node of the linked list |
| `listforeach.go` | Higher-order iterator executing a function on every node |
| `listforeachif.go` | Conditional iterator executing callbacks matching specific predicates |

---

### 2. Mathematics, Recursion & Complexity

| Exercise | Description | Time Complexity |
| :--- | :--- | :--- |
| `fibonacci.go` | Recursive and iterative generation of Fibonacci numbers | $O(n)$ |
| `recursivefactorial.go` | Computes $n!$ recursively with integer overflow protection | $O(n)$ |
| `iterativepower.go` | Fast exponentiation using iterative multiplication | $O(n)$ |
| `sqrt.go` | Calculates square root for natural squares without floating-point libraries | $O(\sqrt{n})$ |
| `isprime.go` | Primality test optimized up to $\sqrt{n}$ | $O(\sqrt{n})$ |
| `findnextprime.go` | Computes the smallest prime greater than or equal to $n$ | $O(k\sqrt{n})$ |
| `collatzcountdown.go` | Simulates the $3n + 1$ Collatz sequence and returns steps to reach $1$ | $O(\text{steps})$ |
| `divmod.go` / `swap.go` | Direct pointer dereferencing for division, modulus, and memory swapping | $O(1)$ |

---

### 3. String Manipulation & Parsing (Byte-Level)

| File | Description |
| :--- | :--- |
| `splitwhitespaces.go` | Custom whitespace tokenizer splitting words on spaces, tabs, and newlines |
| `strrev.go` | In-place byte and rune reversal |
| `join.go` | Joins a slice of strings with a specified separator string |
| `concatparams.go` | Concatenates multi-line parameter inputs separated by `\n` |
| `rot14.go` | Caesar cipher variation shifting alphabetic ASCII values by 14 positions |
| `jumpover.go` | Extracts and returns every third character from an input stream |
| `stringtointslice.go` | Parses formatted strings into strongly-typed integer arrays |

---

### 4. Sorting & Combinatorics

| File | Description |
| :--- | :--- |
| `issorted.go` | Validates slice ordering using custom comparator function pointers |
| `descendcomb.go` | Generates all descending 2-digit number pairs |
| `shoppinglistsort.go` | Custom multi-criteria sorting based on item priority and alphabetical order |
| `sortparams/main.go` | CLI program that accepts arbitrary shell arguments and outputs them ASCII-sorted |

---

### 5. Bitwise & Low-Level Operations

| File | Description |
| :--- | :--- |
| `activebits.go` | Computes Hamming weight (counts number of 1s in an integer's binary representation) using bitwise shifts (`&` and `>>`) |
| `fixthemain/main.go` | Memory struct alignment and boolean status validation |

---

### 6. UNIX Command-Line Tools & Parsing
Standalone executables utilizing `z01.PrintRune`:
* `printprogramname/`: Displays the execution binary name.
* `printparams/`: Iterates through arguments received from `os.Args`.
* `revparams/`: Prints passed arguments in reverse order.
* `comcheck/`: Substring pattern matching on command line inputs.
* `stt/`: Terminal formatting and argument validation utility.

---

## 🛠️ Getting Started & Testing

### Prerequisites
* Go 1.25+ installed on your machine (`go version`)

### Running Solutions
```bash
# Clone the repository
git clone https://github.com/Yassir-Essabbahy/Zone01-Pool.git
cd Zone01-Pool

# Run a CLI tool (e.g., sortparams)
go run ./sortparams/main.go "banana" "apple" "cherry"
# Output:
# apple
# banana
# cherry
```

### Running Unit Tests
You can execute standard Go tests across modules:
```bash
go test ./...
```

---

## 👨‍💻 Author

**Yassir ESSABAHY**  
* Alumni, Zone01 Oujda (01Edu Network)  
* Solo Game Developer & Technical Artist  
* Portfolio: [yessirdev.vercel.app](https://yessirdev.vercel.app)  
* LinkedIn: [linkedin.com/in/yessir001](https://www.linkedin.com/in/yessir001/)  
* Instagram: [@thats_yessir](https://www.instagram.com/thats_yessir)
