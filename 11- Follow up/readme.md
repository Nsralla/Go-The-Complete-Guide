# Basics
1. To run a Go file : go run main.go
2. To create a module: go mod init 'example.com/first-app'.
** This tells go that this current folder and all subfolders bleong to this module **.
3. To build a module: go build.
4. To run module: ./first-app, or whatever the .exe file title.


## Difference Between go run and go build

### go run main.go (Direct Execution)
* Compiles and runs immediately - Go compiles the code in memory and executes it directly
* No executable file is created - The compilation happens temporarily
* Slower for repeated runs - Recompiles every time you run the command
* Best for development and testing - Quick iteration during development
* Memory usage - Compilation happens in memory, then discarded



### go build (Build then Execute)
* Creates an executable file - Produces a binary file (.exe on Windows)
* Two-step process - First build, then run the executable
* Faster for repeated runs - Once built, you can run the executable multiple times without recompilation
* Better for production - You get a standalone executable that can be distributed
* Persistent binary - The executable file remains on disk

**During Development:**
```bash
go run main.go    # Quick iteration, always uses latest code
```

**For Production/Distribution:**
```bash
go build          # Create final executable once
```


## Variables lesson 1
``` go
    fmt.Println("Hellow, World!!")
	fmt.Println("Go" + "Lang")
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("7.0 / 3.0 = ", fmt.Sprintf("%.2f",  7.0 / 3.0))
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(false && false)
```


## Variables lesson 2

This section demonstrates different ways to declare and initialize variables in Go:

### Variable Declaration Methods

1. **Type Inference with `var`**: Go automatically detects the type
   ```go
   var a = "application"  // Go infers this is a string
   ```

2. **Multiple Variable Declaration**: Declare multiple variables in one line
   ```go
   var x, y = 34, 65  // Both are integers
   ```

3. **Boolean Variables**: Simple boolean declaration
   ```go
   var d = true  // Go infers this is a bool
   ```

4. **Zero Value Declaration**: Declare without initialization (gets default value)
   ```go
   var e int  // e gets zero value of int, which is 0
   ```

5. **Short Variable Declaration**: Using `:=` operator (only inside functions)
   ```go
   f := "HELLO"  // Shorthand syntax, Go infers string type
   ```

### Complete Example
```go
var a  = "application"
fmt.Println("Value of a is: ", a)

var x, y = 34, 65
fmt.Println("Value of x is: ", x)
fmt.Println("Value of y is: ", y)

var d = true
fmt.Println("Value of d is: ", d)

var e int
fmt.Println("Value of e is: ", e)  // Prints: 0 (zero value)

f := "HELLO"
fmt.Println("Value of f is: ", f)
```

### Key Points
- **`:=` vs `var`**: `:=` can only be used inside functions, `var` can be used anywhere
- **Zero Values**: Uninitialized variables get their type's zero value (0 for numbers, "" for strings, false for booleans)
- **Type Inference**: Go automatically determines the type based on the assigned value

## Constants

Constants in Go are immutable values that are known at compile time. They provide a way to give names to fixed values and ensure they cannot be changed during program execution.

### Key Features of Constants

1. **Immutable**: Once declared, constants cannot be modified
2. **Compile-time evaluation**: Values must be known when the program is compiled
3. **No memory address**: Constants don't have memory addresses (can't use `&` operator)
4. **Arbitrary precision**: Constant expressions can use unlimited precision arithmetic
5. **Type flexibility**: Untyped constants can be used with compatible types

### Constant Declaration Methods

1. **Basic String Constant**: Declare at package level
   ```go
   const statement = "Application"  // Available throughout the package
   ```

2. **Numeric Constants**: Mathematical values and calculations
   ```go
   const pi = 3.14              // Float constant
   const angle = 90 / pi        // Calculated at compile time
   ```

3. **Large Number Constants**: Go handles big numbers with arbitrary precision
   ```go
   const n = 500000000          // Large integer
   const d = 3e20 / n          // Scientific notation with division
   ```

### Complete Example with Explanations
```go
const statement = "Application"  // Package-level constant

func main(){
    // Using package-level constant
    fmt.Println("This is my first " + statement)

    // Local constants within function
    const pi = 3.14
    const angle = 90 / pi        // Compile-time calculation
    
    // Mixing constants with variables
    var sin = math.Sin(angle)    // sin is a variable, angle is constant
    fmt.Println(int64(sin))      // Type conversion

    fmt.Println("Value of sin is: ", sin)

    // Large number arithmetic with arbitrary precision
    const n = 500000000
    const d = 3e20 / n          // Precise calculation at compile time
    fmt.Println(d)              // Result: 6e+11
}
```

### Constants vs Variables

| Feature | Constants | Variables |
|---------|-----------|-----------|
| **Mutability** | Immutable | Mutable |
| **Declaration** | `const` keyword | `var` or `:=` |
| **Evaluation** | Compile-time | Runtime |
| **Memory** | No memory address | Has memory address |
| **Scope** | Package or function level | Package, function, or block level |

### Best Practices
- Use constants for values that will never change (API keys, mathematical constants, configuration values)
- Declare constants at package level when they're used across multiple functions
- Use meaningful names that describe what the constant represents
- Group related constants together for better organization

## For Loop

Go has only one looping construct: the `for` loop. However, it's flexible enough to handle all common looping scenarios that other languages might use `while`, `do-while`, or `foreach` loops for.

### Types of For Loops in Go

1. **Condition-only Loop** (like `while` in other languages)
   ```go
   i := 1
   for i < 3 {                    // Loop while condition is true
       fmt.Println("i = ", i)
       i++                        // Increment manually
   }
   ```

2. **Traditional For Loop** (with initialization, condition, and increment)
   ```go
   for j := 0; j <= 5; j++ {      // init; condition; increment
       fmt.Println("j= ", j)
   }
   ```

3. **Range-based Loop** (iterate over a range of numbers)
   ```go
   for k := range 9 {             // k goes from 0 to 8
       fmt.Println("k= ", k)
   }
   ```

4. **Infinite Loop** (runs forever until broken)
   ```go
   for {                          // No condition = infinite loop
       fmt.Println("Infinite loop")
       break                      // Must use break to exit
   }
   ```

5. **Loop with Control Statements** (`continue` and `break`)
   ```go
   for l := range 10 {            // l goes from 0 to 9
       if l % 2 == 0 {           // If l is even
           fmt.Println("even skip")
           continue               // Skip rest of iteration
       }
       fmt.Println("l= ", l)      // Only prints odd numbers
   }
   ```

### Complete Example with Explanations
```go
// 1. Condition-only loop (while-style)
i := 1
for i < 3{
    fmt.Println("i = ", i )
    i ++                          // Manual increment required
}

// 2. Traditional for loop
for j :=0; j <= 5; j++{          // Initialization; condition; increment
    fmt.Println("j= ", j)
}

// 3. Range-based loop (Go 1.22+)
for k := range 9{                // Iterates from 0 to 8
    fmt.Println("k= ", k)
}

// 4. Infinite loop
for {
    fmt.Println("Infinite loop")
    break                        // Exit the infinite loop
}

// 5. Loop with continue statement
for l := range 10 {
    if l % 2 == 0 {             // Check if number is even
        fmt.Println("even skip")
        continue                 // Skip to next iteration
    }
    fmt.Println("l= ", l)        // Only executes for odd numbers
}
```

### Key Loop Control Statements

| Statement | Purpose | Effect |
|-----------|---------|--------|
| **`break`** | Exit loop completely | Terminates the entire loop |
| **`continue`** | Skip current iteration | Jumps to next iteration |

### Loop Patterns and Use Cases

- **Condition-only**: Use when you don't know how many iterations you need
- **Traditional**: Use when you know the exact number of iterations
- **Range-based**: Use for simple counting or iterating over sequences
- **Infinite**: Use for event loops, servers, or when exit condition is complex

### Best Practices
- Always ensure loop termination conditions to avoid infinite loops
- Use `range` loops for simple iteration over numbers or collections
- Prefer `continue` over deeply nested `if` statements for cleaner code
- Use descriptive variable names even in short loops

## Switch

The `switch` statement in Go is a powerful control flow structure that allows you to execute different code blocks based on the value of an expression. It's a cleaner alternative to long `if-else` chains and offers several unique features.

### Types of Switch Statements

1. **Basic Switch** (value-based comparison)
   ```go
   i := 3
   switch i {
   case 1:
       fmt.Println("one")
   case 2:
       fmt.Println("two")
   case 3:
       fmt.Println("three")
       fmt.Println("this is also printed")  // Multiple statements allowed
   default:
       fmt.Println("unknown")
   }
   ```

2. **Multiple Values per Case** (comma-separated values)
   ```go
   switch time.Now().Weekday() {
   case time.Friday, time.Saturday:        // Multiple values in one case
       fmt.Println("It's weekend")
   default:
       fmt.Println("it's weekday")
   }
   ```

3. **Expression Switch** (no expression after switch keyword)
   ```go
   t := time.Now()
   switch {                                // No expression = condition-based
   case t.Hour() < 12:
       fmt.Println("Morning")
   default:
       fmt.Println("Afternoon")
   }
   ```

4. **Type Switch** (determine the type of an interface)
   ```go
   whatAmI := func(data any) {
       switch data.(type) {               // Type assertion switch
       case int:
           fmt.Println("I am int")
       case string:
           fmt.Println("I am string")
       case float32:
           fmt.Println("I am float32")
       case []int:
           fmt.Println("I am slice of int")
       case []string:
           fmt.Println("I am slice of string")
       case map[string]any:
           fmt.Println("I am map of string to any")
       default:
           fmt.Println("I am something else")
       }
   }
   ```

### Complete Example with Explanations
```go
func main() {
    // 1. Basic switch with integer value
    i := 3
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
        fmt.Println("this is also printed")    // Multiple statements in case
    default:
        fmt.Println("unknown")
    }

    // 2. Switch with multiple values per case
    switch time.Now().Weekday() {
    case time.Friday, time.Saturday:           // Check multiple values
        fmt.Println("It's weekend")
    default:
        fmt.Println("it's weekday")
    }

    // 3. Expression switch (condition-based)
    t := time.Now()
    switch {                                   // No variable after switch
    case t.Hour() < 12:
        fmt.Println("Morning")
    default:
        fmt.Println("Afternoon")
    }
    
    // 4. Type switch function
    whatAmI := func(data any) {
        switch data.(type) {                   // Check the type of data
        case int:
            fmt.Println("I am int")
        case string:
            fmt.Println("I am string")
        case float32:
            fmt.Println("I am float32")
        case []int:
            fmt.Println("I am slice of int")
        case []string:
            fmt.Println("I am slice of string")
        case map[string]any:
            fmt.Println("I am map of string to any")
        default:
            fmt.Println("I am something else")
        }
    }

    // Testing the type switch with different data types
    whatAmI(3.14)                             // float64 -> "something else"
    whatAmI("hello")                          // string -> "I am string"
    whatAmI([]int{1,2,3})                     // []int -> "I am slice of int"
    whatAmI(map[string]any{"name":"John", "age":30})  // map -> "I am map..."
}
```

### Key Features of Go Switch

| Feature | Description | Example |
|---------|-------------|---------|
| **No fallthrough** | Cases don't fall through by default | Each case executes independently |
| **Multiple values** | One case can handle multiple values | `case 1, 2, 3:` |
| **Any type** | Can switch on any comparable type | strings, numbers, booleans, etc. |
| **Expression switch** | Switch without a variable | `switch { case condition: }` |
| **Type switch** | Determine interface type | `switch v.(type)` |

### Switch vs If-Else

| Aspect | Switch | If-Else Chain |
|--------|--------|---------------|
| **Readability** | Cleaner for multiple conditions | Can become messy |
| **Performance** | Potentially faster | Sequential evaluation |
| **Flexibility** | Limited to equality/type checks | Any boolean expression |
| **Fallthrough** | No automatic fallthrough | N/A |

### Best Practices
- Use `switch` instead of long `if-else` chains when checking equality
- Always include a `default` case when not all possibilities are covered
- Use expression switch for condition-based logic
- Use type switch when working with interfaces and `any` types
- Group related cases using comma separation
- Keep case logic simple; extract complex logic into functions 


## Arrays

Arrays in Go are fixed-size sequences of elements of the same type. The size is part of the array's type, making arrays in Go quite different from arrays in many other languages. Arrays provide the foundation for slices, which are more commonly used in Go programming.

### Key Characteristics of Arrays

1. **Fixed Size**: The length is determined at compile time and cannot be changed
2. **Type Specification**: Both element type and length are part of the array type
3. **Zero Values**: Uninitialized arrays contain zero values of the element type
4. **Value Types**: Arrays are passed by value, not by reference
5. **Index-based**: Elements are accessed using zero-based indexing

### Array Declaration and Initialization Methods

1. **Zero-valued Array Declaration**
   ```go
   var a [5]int                    // Array of 5 integers, all initialized to 0
   fmt.Println(a)                  // Output: [0 0 0 0 0]
   ```

2. **Array with Partial Initialization**
   ```go
   var arr = [7]int{1,11,111,1111} // Specify some values, rest are zero
   fmt.Println(arr)                // Output: [1 11 111 1111 0 0 0]
   ```

3. **Array with Automatic Size Detection**
   ```go
   var arr2 = [...]int{100,1000,2000}  // ... lets Go count the elements
   fmt.Println(arr2)                    // Output: [100 1000 2000]
   fmt.Println(len(arr2))               // Output: 3
   ```

4. **Array with Index-specific Initialization**
   ```go
   b := [...]int{100, 3: 400, 500, 150}  // 100 at index 0, 400 at index 3
   fmt.Println("idx:", b)                 // Output: [100 0 0 400 500 150]
   ```

5. **Multi-dimensional Arrays**
   ```go
   // 2D array with literal initialization
   var twoD = [2][2]int{
       {6,11},     // First row
       {87,32},    // Second row
   }
   fmt.Println(twoD)  // Output: [[6 11] [87 32]]

   // 2D array with separate declaration and initialization
   var twoD2 [2][3]int
   twoD2 = [...][3]int{           // First dimension can use ..., but not others
       {10,100,1000},
       {20,200,2000},
   }
   fmt.Println(twoD2)  // Output: [[10 100 1000] [20 200 2000]]
   ```

### Complete Example with Explanations
```go
func main() {
    // 1. Zero-valued array declaration
    var a [5]int
    fmt.Println(a)                      // Output: [0 0 0 0 0]
    
    // Populate array using loop
    for i := range 5 {
        a[i] = i + 1                    // Assign values 1, 2, 3, 4, 5
    }
    fmt.Println(a)                      // Output: [1 2 3 4 5]
    fmt.Println(len(a))                 // Output: 5

    // 2. Partial initialization (rest are zero values)
    var arr = [7]int{1,11,111,1111}     // Only first 4 elements specified
    fmt.Println(arr)                    // Output: [1 11 111 1111 0 0 0]
    fmt.Println(len(arr))               // Output: 7

    // 3. Automatic size detection
    var arr2 = [...]int{100,1000,2000}  // Go counts: 3 elements
    fmt.Println(arr2)                   // Output: [100 1000 2000]
    fmt.Println(len(arr2))              // Output: 3

    // 4. Index-specific initialization
    b := [...]int{100, 3: 400, 500, 150}  // Index 0=100, 3=400, 4=500, 5=150
    fmt.Println("idx:", b)              // Output: [100 0 0 400 500 150]

    // 5. 2D arrays
    var twoD = [2][2]int{
        {6,11},                         // Row 0
        {87,32},                        // Row 1
    }
    fmt.Println(twoD)                   // Output: [[6 11] [87 32]]

    // 2D array with separate initialization
    var twoD2 [2][3]int
    twoD2 = [...][3]int{               // First dimension uses ..., inner must be explicit
        {10,100,1000},
        {20,200,2000},
    }
    fmt.Println(twoD2)                  // Output: [[10 100 1000] [20 200 2000]]
}
```

### Array vs Slice Comparison

| Feature | Arrays | Slices |
|---------|--------|--------|
| **Size** | Fixed at compile time | Dynamic, can grow/shrink |
| **Type** | `[n]Type` (size included) | `[]Type` (no size) |
| **Memory** | Value type (copied) | Reference type (pointer to array) |
| **Usage** | Less common | More common in Go |
| **Performance** | Predictable memory layout | Slight overhead for flexibility |

### Important Notes

1. **Index-specific Initialization**: `{100, 3: 400, 500, 150}` means:
   - Index 0: 100
   - Index 1-2: 0 (zero values)
   - Index 3: 400
   - Index 4: 500  
   - Index 5: 150

2. **Multi-dimensional Arrays**: Only the first dimension can use `...` for auto-sizing
   ```go
   // ✅ Valid: [...][3]int
   // ❌ Invalid: [2][...]int
   ```

3. **Array Assignment**: Arrays are value types, so assignment copies all elements
   ```go
   var x [3]int = [3]int{1,2,3}  // Creates a copy
   ```

### Best Practices
- Use slices instead of arrays for most cases (more flexible)
- Use arrays when you need fixed-size collections with compile-time guarantees
- Prefer array literals over manual element assignment when possible
- Use `len()` function to get array length programmatically
- Consider memory implications: large arrays are copied when passed to functions


## Slices

Slices are one of the most important and commonly used data structures in Go. Unlike arrays, slices are dynamic and can grow or shrink during runtime. A slice is a reference type that provides a flexible interface to arrays, making them much more practical for everyday programming.

### Key Characteristics of Slices

1. **Dynamic Size**: Can grow and shrink during runtime
2. **Reference Type**: Points to an underlying array, doesn't store data directly
3. **Three Components**: Pointer to array, length, and capacity
4. **Zero Value**: `nil` (unlike arrays which have zero values of their element type)
5. **Flexible**: More commonly used than arrays in Go programming

### Slice vs Array Comparison

```go
var a = []int{1,2,3,4}    // Slice - no size specified
var b = [4]int{1,2,3,4}   // Array - size explicitly defined

fmt.Printf("Type of a is %T\n", a)  // []int (slice)
fmt.Printf("Type of b is %T\n", b)  // [4]int (array)
```

| Feature | Slice (`[]int`) | Array (`[4]int`) |
|---------|-----------------|------------------|
| **Declaration** | `[]int{...}` (no size) | `[4]int{...}` (size defined) |
| **Type** | Reference type | Value type |
| **Size** | Dynamic | Fixed |
| **Memory** | Points to array | Stored directly |
| **Usage** | More common | Less common |

### Slice Declaration and Initialization Methods

1. **Nil Slice Declaration**
   ```go
   var c []string                    // Creates a nil slice
   fmt.Println(len(c))              // Output: 0
   fmt.Println(cap(c))              // Output: 0
   fmt.Println(c == nil)            // Output: true
   ```

2. **Slice Literal**
   ```go
   var z = []int{1,10,20,30}        // Initialize with values
   fmt.Println(z)                   // Output: [1 10 20 30]
   fmt.Println(len(z))              // Output: 4 (length)
   fmt.Println(cap(z))              // Output: 4 (capacity)
   ```

3. **Using `make()` Function**
   ```go
   var m = make([]int, 5)           // Create slice with length 5
   fmt.Println(m)                   // Output: [0 0 0 0 0]
   fmt.Println(len(m))              // Output: 5
   fmt.Println(cap(m))              // Output: 5
   
   // Assign values to specific indices
   m[0] = 9
   m[1] = 1987
   fmt.Println(m)                   // Output: [9 1987 0 0 0]
   ```

### Slice Operations

#### Copying Slices
```go
copyM := make([]int, len(m))        // Create destination slice
copy(copyM, m)                      // Copy elements from m to copyM
copyM[4] = 1000                     // Modify copy independently
fmt.Println("m is ", m)             // Original unchanged
fmt.Println("copyM is ", copyM)     // Copy has modifications
```

#### Slice Slicing (Sub-slices)
```go
fmt.Println(copyM[:2])              // [start:end] - first 2 elements
fmt.Println(copyM[2:])              // [start:] - from index 2 to end
fmt.Println(copyM[:5])              // [:end] - from start to index 5
```

#### Comparing Slices
```go
if slices.Equal(m, copyM) {         // Use slices.Equal() function
    fmt.Println("m and copyM are equal")
} else {
    fmt.Println("m and copyM are not equal")
}
```

### Multi-dimensional Slices

```go
twoD := make([][]int, 3)            // Create slice with 3 rows
for i := range 3 {
    columns := i + 1                // Variable number of columns
    twoD[i] = make([]int, columns)  // Create each row
    for j := range columns {
        twoD[i][j] = i + j          // Populate with values
    }
}
fmt.Println(twoD)                   // Output: [[0] [1 2] [2 3 4]]

// Add another row dynamically
twoD = append(twoD, []int{1, 2, 3})
fmt.Println(twoD)                   // Output: [[0] [1 2] [2 3 4] [1 2 3]]
```

### Complete Example with Explanations
```go
func main() {
    // 1. Slice vs Array comparison
    var a = []int{1,2,3,4}           // Slice - reference type
    var b = [4]int{1,2,3,4}          // Array - value type
    
    fmt.Printf("Type of a is %T\n", a)  // []int
    fmt.Printf("Type of b is %T\n", b)  // [4]int

    // 2. Nil slice
    var c []string
    fmt.Println(len(c))              // 0
    fmt.Println(cap(c))              // 0
    fmt.Println(c == nil)            // true

    // 3. Slice literal
    var z = []int{1,10,20,30}
    fmt.Println(z)                   // [1 10 20 30]
    fmt.Println(len(z))              // 4
    fmt.Println(cap(z))              // 4

    // 4. Using make()
    var m = make([]int, 5)
    fmt.Println(m)                   // [0 0 0 0 0]
    m[0] = 9
    m[1] = 1987
    fmt.Println(m)                   // [9 1987 0 0 0]

    // 5. Copying slices
    copyM := make([]int, len(m))
    copy(copyM, m)
    copyM[4] = 1000
    fmt.Println("m is ", m)          // [9 1987 0 0 0]
    fmt.Println("copyM is ", copyM)  // [9 1987 0 0 1000]

    // 6. Slicing operations
    fmt.Println(copyM[:2])           // [9 1987]
    fmt.Println(copyM[2:])           // [0 0 1000]
    fmt.Println(copyM[:5])           // [9 1987 0 0 1000]

    // 7. String slice example
    t := make([]string, 3)
    t[0] = "Hello"
    t[1] = "World"
    t[2] = "!"
    fmt.Println(t)                   // [Hello World !]

    // 8. Slice comparison
    if slices.Equal(m, copyM) {
        fmt.Println("m and copyM are equal")
    } else {
        fmt.Println("m and copyM are not equal")
    }

    // 9. Multi-dimensional slices
    twoD := make([][]int, 3)
    for i := range 3 {
        columns := i + 1
        twoD[i] = make([]int, columns)
        for j := range columns {
            twoD[i][j] = i + j
        }
    }
    fmt.Println(twoD)                // [[0] [1 2] [2 3 4]]
    
    // Adding rows dynamically
    twoD = append(twoD, []int{1, 2, 3})
    fmt.Println(twoD)                // [[0] [1 2] [2 3 4] [1 2 3]]
}
```

### Slice Internals

A slice has three components:
1. **Pointer**: Points to the first element of the underlying array
2. **Length**: Number of elements in the slice (`len()`)
3. **Capacity**: Number of elements from the start of the slice to the end of the underlying array (`cap()`)

### Common Slice Operations

| Operation | Syntax | Description |
|-----------|--------|-------------|
| **Create empty** | `make([]int, length)` | Create slice with specified length |
| **Create with capacity** | `make([]int, len, cap)` | Create slice with length and capacity |
| **Copy** | `copy(dest, src)` | Copy elements from source to destination |
| **Append** | `append(slice, elements...)` | Add elements to end of slice |
| **Sub-slice** | `slice[start:end]` | Create sub-slice from existing slice |
| **Compare** | `slices.Equal(a, b)` | Compare two slices for equality |

### Key Differences from Arrays

- **Dynamic sizing**: Slices can grow using `append()`
- **Reference semantics**: Multiple slices can point to the same underlying array
- **Nil comparison**: Slices can be `nil`, arrays cannot
- **More flexible**: Better for most use cases than arrays

### Best Practices
- Use slices instead of arrays for most scenarios
- Always check if a slice is `nil` before using it
- Use `make()` to create slices with known capacity to avoid reallocations
- Use `copy()` when you need independent copies of slice data
- Use `slices.Equal()` to compare slice contents (Go 1.21+)
- Be aware that sub-slices share the same underlying array

## Functions

Functions are first-class citizens in Go and are the primary way to organize and structure code. Go functions support multiple return values, named return parameters, and various parameter patterns that make them both powerful and flexible.

### Key Characteristics of Functions

1. **First-class citizens**: Functions can be assigned to variables, passed as arguments, and returned from other functions
2. **Multiple return values**: Functions can return multiple values simultaneously
3. **Named return parameters**: Return values can be named and used as variables within the function
4. **Zero values**: Named return parameters are initialized to their zero values
5. **Clean syntax**: Simple and consistent function declaration syntax

### Function Declaration Syntax

```go
func functionName(parameters) (returnTypes) {
    // function body
    return values
}
```

### Basic Function Examples

#### 1. Function with Multiple Return Values and Named Returns
```go
func getCoords() (x, y int) {          // Named return parameters
    fmt.Println(x, y)                  // Output: 0 0 (zero values)
    x = 100                            // Assign to named return
    y = 200                            // Assign to named return
    return                             // Naked return (returns x and y)
}

func main() {
    x, y := getCoords()                // Receive multiple return values
    fmt.Println(x, y)                  // Output: 100 200
}
```

### Function Parameter and Return Patterns

#### 1. No Parameters, No Return
```go
func sayHello() {
    fmt.Println("Hello, World!")
}
```

#### 2. Parameters with Return Value
```go
func add(a, b int) int {               // Same type parameters can be grouped
    return a + b
}
```

#### 3. Multiple Parameters of Different Types
```go
func greet(name string, age int) string {
    return fmt.Sprintf("Hello %s, you are %d years old", name, age)
}
```

#### 4. Multiple Return Values (Explicit)
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

#### 5. Named Return Parameters
```go
func calculate(a, b int) (sum, product int) {
    sum = a + b                        // Named returns act as variables
    product = a * b
    return                             // Naked return
}
```

#### 6. Variadic Functions (Variable Arguments)
```go
func sum(numbers ...int) int {         // Accept variable number of int arguments
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage:
result := sum(1, 2, 3, 4, 5)          // Pass multiple arguments
slice := []int{1, 2, 3, 4, 5}
result2 := sum(slice...)               // Expand slice with ...
```

### Advanced Function Concepts

#### 1. Functions as Variables
```go
func main() {
    // Assign function to variable
    var operation func(int, int) int = add
    result := operation(5, 3)          // result = 8
    
    // Anonymous function
    multiply := func(a, b int) int {
        return a * b
    }
    result2 := multiply(4, 5)          // result2 = 20
}
```

#### 2. Higher-Order Functions (Functions as Parameters)
```go
func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    result := applyOperation(10, 5, add)       // Pass function as argument
    
    // Pass anonymous function
    result2 := applyOperation(10, 5, func(x, y int) int {
        return x - y
    })
}
```

#### 3. Functions Returning Functions
```go
func makeAdder(x int) func(int) int {
    return func(y int) int {           // Return a closure
        return x + y                   // Captures x from outer scope
    }
}

func main() {
    addFive := makeAdder(5)            // Create specific adder function
    result := addFive(10)              // result = 15
}
```

### Complete Example with Explanations
```go
package main

import "fmt"

func main() {
    // 1. Basic function call with multiple returns
    x, y := getCoords()
    fmt.Println(x, y)                  // Output: 100 200
    
    // 2. Function with parameters and return
    sum := add(10, 20)
    fmt.Println("Sum:", sum)           // Output: Sum: 30
    
    // 3. Function with error handling
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // Output: Result: 5
    }
    
    // 4. Variadic function
    total := sum(1, 2, 3, 4, 5)
    fmt.Println("Total:", total)       // Output: Total: 15
    
    // 5. Function as variable
    operation := add
    fmt.Println("Operation result:", operation(7, 3))  // Output: 10
}

// Named return parameters with naked return
func getCoords() (x, y int) {
    fmt.Println(x, y)                  // 0 0 (zero values)
    x = 100
    y = 200
    return                             // Returns x and y
}

// Basic function with parameters
func add(a, b int) int {
    return a + b
}

// Function with error return
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

### Function Best Practices

#### Error Handling Pattern
```go
func doSomething() error {
    // Return error as last return value
    if someCondition {
        return fmt.Errorf("something went wrong")
    }
    return nil                         // nil means no error
}

func main() {
    if err := doSomething(); err != nil {
        fmt.Println("Error:", err)
    }
}
```

#### Multiple Return Values vs Structs
```go
// Good for 2-3 related values
func getNameAndAge() (string, int) {
    return "Alice", 30
}

// Better for complex data
type Person struct {
    Name string
    Age  int
    City string
}

func getPerson() Person {
    return Person{Name: "Alice", Age: 30, City: "NYC"}
}
```

### Key Function Features Summary

| Feature | Example | Use Case |
|---------|---------|----------|
| **Named Returns** | `func f() (x int)` | Clean code, documentation |
| **Multiple Returns** | `func f() (int, error)` | Return value + error |
| **Variadic** | `func f(args ...int)` | Variable number of arguments |
| **Anonymous** | `func(x int) int { return x*2 }` | Short, one-time functions |
| **Closures** | Functions that capture outer variables | State preservation |

### Best Practices
- Use named returns for documentation, but prefer explicit returns for clarity
- Always return errors as the last return value
- Keep functions small and focused on a single responsibility
- Use variadic functions when you need flexible argument lists
- Prefer explicit returns over naked returns for better readability
- Use function types for callbacks and configurable behavior
- Consider using structs instead of many return values (>3)


## Maps

Maps in Go are key-value data structures similar to hash tables, dictionaries, or associative arrays in other languages. They provide fast lookups, additions, and deletions based on keys. Maps are reference types and are one of the most commonly used data structures in Go.

### Key Characteristics of Maps

1. **Key-Value Storage**: Store data as key-value pairs for fast retrieval
2. **Reference Type**: Maps are reference types, passed by reference
3. **Dynamic Size**: Can grow and shrink during runtime
4. **Zero Value**: The zero value of a map is `nil`
5. **Unordered**: Maps don't maintain insertion order (Go 1.0+)

### Map Declaration and Initialization Methods

#### 1. Using `make()` Function (Empty Map)
```go
a := make(map[string]int)            // Create empty map
a["one"] = 1                         // Add key-value pairs
a["two"] = 2
a["three"] = 3
fmt.Println(a)                       // Output: map[one:1 two:2 three:3]
```

#### 2. Map Literal (Pre-populated)
```go
n := map[string]int{                 // Create and initialize in one step
    "one": 1,
    "two": 2,
}
fmt.Println(n)                       // Output: map[one:1 two:2]
n["three"] = 3                       // Can still add more elements
n["four"] = 4
fmt.Println(n)                       // Output: map[one:1 two:2 three:3 four:4]
```

### Map Operations

#### 1. Accessing Values
```go
fmt.Println(a["one"])                // Output: 1 (existing key)
fmt.Println(a["five"])               // Output: 0 (non-existent key returns zero value)
```

#### 2. Safe Key Lookup (Comma OK Idiom)
```go
value, ok := a["six"]                // Check if key exists
if ok {
    fmt.Println("Value:", value)     // Key exists
} else {
    fmt.Println("Key not present", ok) // Key doesn't exist, ok = false
}
```

#### 3. Iterating Over Maps
```go
for key, value := range a {          // Iterate over key-value pairs
    fmt.Println("Key:", key, ", Value:", value)
}
// Note: Iteration order is not guaranteed
```

#### 4. Map Length
```go
fmt.Println(len(a))                  // Output: 3 (number of key-value pairs)
```

#### 5. Deleting Elements
```go
delete(a, "two")                     // Remove key-value pair
fmt.Println(a)                       // Output: map[one:1 three:3]
```

#### 6. Clearing All Elements (Go 1.21+)
```go
clear(a)                             // Remove all key-value pairs
fmt.Println(a)                       // Output: map[] (empty map)
```

### Complete Example with Explanations
```go
func main() {
    // 1. Create empty map using make()
    a := make(map[string]int)
    
    // 2. Add key-value pairs
    a["one"] = 1
    a["two"] = 2
    a["three"] = 3
    fmt.Println(a)                   // map[one:1 two:2 three:3]

    // 3. Iterate over map
    for key, value := range a {
        fmt.Println("Key:", key, ", Value:", value)
    }

    // 4. Get map length
    fmt.Println(len(a))              // 3

    // 5. Access non-existent key (returns zero value)
    fmt.Println(a["five"])           // 0

    // 6. Safe key lookup with comma ok idiom
    value, ok := a["six"]
    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not present", ok)  // Key not present false
    }

    // 7. Delete element
    delete(a, "two")
    fmt.Println(a)                   // map[one:1 three:3]

    // 8. Clear all elements
    clear(a)
    fmt.Println(a)                   // map[]

    fmt.Println("----------------------------------------")
    
    // 9. Map literal initialization
    n := map[string]int{
        "one": 1,
        "two": 2,
    }
    fmt.Println(n)                   // map[one:1 two:2]
    
    // 10. Add more elements
    n["three"] = 3
    n["four"] = 4
    fmt.Println(n)                   // map[one:1 two:2 three:3 four:4]
    
    // 11. Access non-existent key
    fmt.Println(n["five"])           // 0
}
```

### Map Declaration Patterns

#### 1. Different Key-Value Types
```go
// String keys, int values
userAges := map[string]int{
    "Alice": 30,
    "Bob":   25,
}

// Int keys, string values
statusCodes := map[int]string{
    200: "OK",
    404: "Not Found",
    500: "Internal Server Error",
}

// String keys, slice values
groups := map[string][]string{
    "admins": {"alice", "bob"},
    "users":  {"charlie", "david"},
}
```

#### 2. Nested Maps
```go
nestedMap := map[string]map[string]int{
    "fruits": {
        "apple":  5,
        "banana": 3,
    },
    "vegetables": {
        "carrot": 10,
        "potato": 8,
    },
}
```

### Key Map Operations Summary

| Operation | Syntax | Description |
|-----------|--------|-------------|
| **Create empty** | `make(map[KeyType]ValueType)` | Create empty map |
| **Create with data** | `map[KeyType]ValueType{k1: v1, k2: v2}` | Create with initial data |
| **Add/Update** | `m[key] = value` | Add new or update existing |
| **Access** | `value := m[key]` | Get value (zero value if not exists) |
| **Safe access** | `value, ok := m[key]` | Check existence while accessing |
| **Delete** | `delete(m, key)` | Remove key-value pair |
| **Clear** | `clear(m)` | Remove all elements (Go 1.21+) |
| **Length** | `len(m)` | Get number of key-value pairs |
| **Iterate** | `for k, v := range m` | Iterate over all pairs |

### Important Map Behaviors

#### 1. Zero Values for Missing Keys
```go
m := make(map[string]int)
fmt.Println(m["missing"])            // 0 (zero value for int)

m2 := make(map[string]string)
fmt.Println(m2["missing"])           // "" (zero value for string)
```

#### 2. Maps are Reference Types
```go
original := map[string]int{"a": 1}
copy := original                     // Both point to same underlying data
copy["b"] = 2
fmt.Println(original)                // map[a:1 b:2] - original is modified!
```

#### 3. Key Requirements
- Keys must be **comparable types**: strings, numbers, booleans, arrays, structs with comparable fields
- **Cannot use**: slices, maps, or functions as keys
- **Can use**: any comparable type including custom structs

### Best Practices

1. **Always use comma ok idiom** when you're not sure if a key exists
2. **Initialize maps** before use (either with `make()` or map literal)
3. **Check for nil maps** before operations to avoid panics
4. **Use meaningful key names** for better code readability
5. **Consider using structs** instead of nested maps for complex data
6. **Remember maps are unordered** - don't rely on iteration order

### Common Patterns

#### Safe Map Access
```go
if value, exists := myMap[key]; exists {
    // Key exists, use value
    fmt.Println("Found:", value)
} else {
    // Key doesn't exist
    fmt.Println("Key not found")
}
```

#### Map as Set (using bool values)
```go
visited := map[string]bool{
    "page1": true,
    "page2": true,
}

if visited["page1"] {
    fmt.Println("Already visited page1")
}
```


## Advanced Slice Operations & Variadic Functions

This section covers advanced slice manipulation techniques, variadic functions, and important slice behaviors including copying, appending, and element removal. Understanding these concepts is crucial for effective slice usage in Go.

### Variadic Functions

Variadic functions can accept a variable number of arguments of the same type using the `...` syntax.

#### Variadic Function Declaration
```go
func sum(nums ...int) (total_sum int) {    // Accept variable number of int arguments
    for _, num := range nums {
        total_sum += num
    }
    return                                 // Named return with naked return
}
```

#### Calling Variadic Functions
```go
// Method 1: Pass individual arguments
fmt.Println("Sum:", sum(1,2,3,4,5))       // Output: Sum: 15
fmt.Println("Sum:", sum(10,20,30))        // Output: Sum: 60

// Method 2: Expand slice with ... operator
nums := []int{1,2,3,4,5}
fmt.Println("Sum:", sum(nums...))         // Expand slice to individual arguments
```

### Advanced Slice Copying Techniques

Understanding different ways to copy slices and their implications on the underlying array.

#### 1. Slice Expression Copying (Shallow Copy)
```go
nums := []int{1,2,3,4,5}
nums = append(nums, 5,5,5)               // Add more elements

new_nums := nums[:]                      // Creates a new slice header
                                        // BUT shares the same underlying array
new_nums = append(new_nums, 100, 200)
fmt.Println("new_nums:", new_nums)       // [1 2 3 4 5 5 5 5 100 200]
fmt.Println("nums:", nums)               // [1 2 3 4 5 5 5 5] - original unchanged
```

#### 2. Direct Assignment (Shared Reference)
```go
// new_nums := nums                      // Both slices point to same underlying array
                                        // Changes to new_nums WILL affect nums
```

#### 3. Deep Copy with make() and copy()
```go
// new_nums := make([]int, len(nums))    // Create new slice with same length
// copy(new_nums, nums)                 // Copy elements to new underlying array
                                        // Changes to new_nums won't affect nums
```

### Slice Element Removal Techniques

#### Removing Element at Specific Index
```go
nums := []int{1,2,3,4,5,5,5,5}
// Remove element at index 2 (value 3)
nums = append(nums[:2], nums[3:]...)     // Combine slices before and after index 2
fmt.Println("nums after removing index 2:", nums)  // [1 2 4 5 5 5 5]
```

**How `append(nums[:2], nums[3:]...)` works:**
- `nums[:2]` = `[1, 2]` (elements before index 2)
- `nums[3:]` = `[4, 5, 5, 5, 5]` (elements after index 2)
- `nums[3:]...` expands the slice to individual elements
- `append()` combines them: `[1, 2] + [4, 5, 5, 5, 5] = [1, 2, 4, 5, 5, 5, 5]`

### Complete Example with Explanations
```go
package main

import "fmt"

func main() {
    // 1. Variadic function calls
    fmt.Println("Sum:", sum(1,2,3,4,5))    // Individual arguments
    fmt.Println("Sum:", sum(10,20,30))     // Different number of arguments
    
    // 2. Slice preparation
    nums := []int{1,2,3,4,5}
    nums = append(nums, 5,5,5)             // Add multiple elements
    
    // 3. Slice copying (shallow copy - shares underlying array)
    new_nums := nums[:]                    // New slice header, same underlying array
    new_nums = append(new_nums, 100, 200) // Append to copy
    fmt.Println("new_nums:", new_nums)     // [1 2 3 4 5 5 5 5 100 200]
    fmt.Println("nums:", nums)             // [1 2 3 4 5 5 5 5] - original unchanged
    
    // 4. Element removal at specific index
    nums = append(nums[:2], nums[3:]...)   // Remove element at index 2
    fmt.Println("nums after removing index 2:", nums)  // [1 2 4 5 5 5 5]
}

// Variadic function with named return
func sum(nums ...int) (total_sum int) {
    for _, num := range nums {
        total_sum += num
    }
    return                                 // Naked return
}
```

### Slice Copying Comparison

| Method | Syntax | Underlying Array | Use Case |
|--------|--------|------------------|----------|
| **Slice Expression** | `new := original[:]` | Shared | When you want a new view of same data |
| **Direct Assignment** | `new := original` | Shared | When you want the same slice reference |
| **Deep Copy** | `copy(new, original)` | Independent | When you need completely separate data |

### Element Removal Patterns

#### Remove First Element
```go
slice = slice[1:]                        // Remove first element
```

#### Remove Last Element
```go
slice = slice[:len(slice)-1]             // Remove last element
```

#### Remove Element at Index i
```go
slice = append(slice[:i], slice[i+1:]...)  // Remove element at index i
```

#### Remove Multiple Elements (Range)
```go
slice = append(slice[:start], slice[end:]...)  // Remove elements from start to end-1
```

### Key Insights

1. **Slice Expression `[:]`**: Creates a new slice header but shares the underlying array
2. **Append Behavior**: When capacity is sufficient, append modifies the underlying array
3. **Element Removal**: Uses slice expressions and append to "skip over" unwanted elements
4. **Variadic Expansion**: Use `...` to expand slices into individual arguments
5. **Named Returns**: Can be used with naked returns for cleaner function signatures

### Best Practices

- **Use deep copy** when you need independent slice data
- **Be aware of shared arrays** when using slice expressions
- **Check capacity** before appending to avoid unexpected behavior
- **Use variadic functions** for flexible APIs that accept variable arguments
- **Understand slice expressions** for efficient element removal
- **Prefer explicit returns** over naked returns for complex functions

### Common Pitfalls

1. **Shared underlying arrays** can lead to unexpected modifications
2. **Index out of bounds** when removing elements without proper bounds checking
3. **Memory leaks** when removing elements from large slices (removed elements may still be referenced)
4. **Capacity changes** during append operations can cause slice copies


## Closure

Closures are functions that capture and access variables from their surrounding (outer) scope, even after the outer function has returned. This creates a powerful mechanism for maintaining state and creating specialized functions in Go.

### What is a Closure?

A closure is formed when:
1. **Inner function** is defined inside an outer function
2. **Inner function** references variables from the outer function's scope
3. **Outer function** returns the inner function
4. **Variables** from outer scope remain accessible to the inner function

### Key Characteristics of Closures

1. **State Preservation**: Variables from outer scope are "captured" and persist
2. **Encapsulation**: Each closure maintains its own copy of captured variables
3. **Function Factory**: Outer functions can create customized inner functions
4. **Memory Persistence**: Captured variables stay in memory as long as closure exists
5. **Lexical Scoping**: Inner function has access to outer function's variable scope

### Basic Closure Example

```go
func outer() func() int {              // Returns a function that returns int
    i := 0                             // Variable in outer scope
    return func() int {                // Anonymous function (closure)
        i++                            // Accesses and modifies outer variable
        return i                       // Returns the updated value
    }
}
```

### How Closures Work

When you call `outer()`, it:
1. **Creates** a local variable `i` with value `0`
2. **Defines** an anonymous function that references `i`
3. **Returns** that anonymous function
4. **Preserves** the variable `i` even after `outer()` finishes

### Complete Example with Explanations

```go
package main

import "fmt"

func main() {
    // 1. Create first closure instance
    call_func := outer()               // call_func now holds the returned function
    fmt.Println(call_func())          // Output: 1 (i becomes 1)
    fmt.Println(call_func())          // Output: 2 (i becomes 2)
    fmt.Println(call_func())          // Output: 3 (i becomes 3)

    // 2. Create second closure instance (independent state)
    call2 := outer()                  // call2 has its own separate 'i' variable
    fmt.Println(call2())              // Output: 1 (new i, starts from 0)
}

func outer() func() int {
    i := 0                            // Local variable captured by closure
    return func() int {               // Return anonymous function
        i++                           // Increment captured variable
        return i                      // Return current value
    }
}
```

### Closure Behavior Analysis

#### Independent State
```go
call_func := outer()    // Creates closure with its own 'i'
call2 := outer()        // Creates another closure with separate 'i'

// Each closure maintains independent state:
call_func() // Returns 1 (its own i: 0 -> 1)
call_func() // Returns 2 (its own i: 1 -> 2)
call2()     // Returns 1 (separate i: 0 -> 1)
```

#### State Persistence
```go
counter := outer()      // Create closure
// The variable 'i' from outer() stays alive
// Even though outer() function has finished executing
counter() // i is still accessible and modifiable
```

### Advanced Closure Patterns

#### 1. Closure with Parameters
```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor              // Captures 'factor' from outer scope
    }
}

// Usage:
double := multiplier(2)               // Creates closure that multiplies by 2
triple := multiplier(3)               // Creates closure that multiplies by 3
fmt.Println(double(5))                // Output: 10
fmt.Println(triple(5))                // Output: 15
```

#### 2. Closure for Configuration
```go
func configuredPrinter(prefix string) func(string) {
    return func(message string) {
        fmt.Printf("%s: %s\n", prefix, message)
    }
}

// Usage:
errorLogger := configuredPrinter("ERROR")
infoLogger := configuredPrinter("INFO")
errorLogger("Something went wrong")    // Output: ERROR: Something went wrong
infoLogger("Process completed")        // Output: INFO: Process completed
```

#### 3. Closure with Multiple Captured Variables
```go
func makeCounter(start, step int) func() int {
    current := start
    return func() int {
        current += step               // Captures both 'current' and 'step'
        return current
    }
}

// Usage:
byTwos := makeCounter(0, 2)          // Count by 2s starting from 0
byFives := makeCounter(10, 5)        // Count by 5s starting from 10
```

### Closure vs Regular Functions

| Aspect | Regular Function | Closure |
|--------|------------------|---------|
| **State** | Stateless | Maintains state |
| **Scope Access** | Only parameters and globals | Outer function's variables |
| **Memory** | No persistent local variables | Captured variables persist |
| **Instances** | Single behavior | Each instance can have different captured values |

### Practical Use Cases

#### 1. Event Handlers with Context
```go
func makeClickHandler(buttonName string) func() {
    clickCount := 0
    return func() {
        clickCount++
        fmt.Printf("%s clicked %d times\n", buttonName, clickCount)
    }
}
```

#### 2. Middleware Functions
```go
func withLogging(operation func()) func() {
    return func() {
        fmt.Println("Starting operation...")
        operation()
        fmt.Println("Operation completed")
    }
}
```

#### 3. Memoization (Caching Results)
```go
func memoize(fn func(int) int) func(int) int {
    cache := make(map[int]int)
    return func(n int) int {
        if result, exists := cache[n]; exists {
            return result
        }
        result := fn(n)
        cache[n] = result
        return result
    }
}
```

### Memory Considerations

#### What Gets Captured
- Only **referenced variables** from outer scope are captured
- **Entire variable**, not just the value at time of creation
- Variables remain in memory as long as closure exists

#### Memory Management
```go
func outer() func() int {
    largeArray := make([]int, 1000000)  // This will be kept in memory
    counter := 0                        // This will also be kept
    return func() int {
        counter++                       // Only counter is used
        return counter                  // But largeArray is still captured!
    }
}
```

### Best Practices

1. **Minimize captured variables** to reduce memory usage
2. **Be explicit** about what variables you're capturing
3. **Use closures for state** when you need persistent local variables
4. **Consider alternatives** like structs with methods for complex state
5. **Understand lifetime** - closures keep captured variables alive

### Common Patterns

#### Factory Pattern
```go
func makeValidator(rules []string) func(string) bool {
    return func(input string) bool {
        // Validate input against captured rules
        return true // Simplified
    }
}
```

#### Callback with Context
```go
func processAsync(data string, callback func(string)) {
    go func() {
        result := "processed: " + data
        callback(result)              // Closure captures callback
    }()
}
```

### Key Takeaways

- **Closures capture variables by reference**, not by value
- **Each closure instance** has independent captured variables
- **Outer function variables persist** as long as closure exists
- **Powerful for creating specialized functions** with embedded state
- **Memory overhead** - be mindful of what gets captured