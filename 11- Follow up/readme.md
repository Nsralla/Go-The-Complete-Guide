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