# Understanding Scope in Go

In simple terms, **scope** refers to the permission or accessibility of a variable in different parts of the program.

In Go, variables and functions declared outside the `main` function are in the **global scope**. However, Go defines three primary scopes:

---

## 1. **Global Scope**
Variables and functions declared outside the `main` function or any other function/block fall under the **global scope**. These can be accessed throughout the program.

---

## 2. **Local Scope**
Variables declared inside a block `{ }` are considered to be in the **local scope** (or block scope). These are accessible only within that specific block and are referred to as **local variables**.

---

## 3. **Package Scope**
**Package scope** allows sharing code (functions or variables) across multiple files within the same package. To achieve this, you need to:
- Create a new package in a separate folder.
- Ensure that functions/variables in the package that you want to access in other files start with an uppercase letter.

### Example: Using Package Scope
Let’s see how package scope works with an example:

#### Folder Structure: project/ ├── main.go └── mathlib/ └── math.go
#### **main.go**
```go
package main

import (
	"fmt"
	"example.com/mathlib" // Importing the custom package
)

func main() {
	fmt.Println("Showing custom package")
	mathlib.Add(3, 4) // Accessing the Add function from the mathlib package
}

```
#### **mathlib/math.go**
```go
package mathlib

import "fmt"

// Function to add two numbers
func Add(x int, y int) {
	z := x + y
	fmt.Println("The total sum is", z)
}
```


