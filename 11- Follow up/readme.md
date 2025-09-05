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