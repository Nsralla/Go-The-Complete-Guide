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
