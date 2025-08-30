# Length vs Capacity in Go

In Go (Golang), length and capacity are properties of slices and arrays, but they behave differently.

## 🔹 len() – Length

Returns the number of elements the slice currently holds.

- **For arrays**: fixed, equals the total number of elements.
- **For slices**: dynamic, changes when slicing or appending.

## 🔸 cap() – Capacity

Returns the maximum number of elements the slice can hold before allocating a new underlying array.

Capacity is from the start of the slice to the end of the underlying array.

## 📦 Example:

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:3] // slice = [2, 3]

    fmt.Println("Length:", len(slice))   // Output: 2
    fmt.Println("Capacity:", cap(slice)) // Output: 4
}
```

## 🔍 Why is capacity 4?

The slice `arr[1:3]` starts at index 1 of arr and can go up to index 5 (len(arr)), so:

```
Indexes:      0   1   2   3   4
Values:       1   2   3   4   5
                  ↑
              slice starts here
```

- `len(slice) = 3 - 1 = 2`
- `cap(slice) = 5 - 1 = 4` (from start index to end of arr)

## 📈 When is this useful?

- **When appending**: Go can reuse the underlying array if there's remaining capacity.
- **When managing memory and performance**.

## 🚀 Dynamic Arrays in Go

Go doesn't have traditional "dynamic arrays" like some other languages, but **slices** provide dynamic array functionality.

### Key Features of Dynamic Arrays (Slices):

- **Automatic resizing**: When you append beyond capacity, Go automatically allocates a larger underlying array
- **Memory management**: Go handles memory allocation and deallocation automatically
- **Growth strategy**: Typically doubles the capacity when reallocation is needed

### Example of Dynamic Behavior:

```go
package main

import "fmt"

func main() {
    // Start with empty slice
    var numbers []int
    fmt.Printf("Initial: len=%d, cap=%d\n", len(numbers), cap(numbers))
    
    // Add elements dynamically
    for i := 1; i <= 10; i++ {
        numbers = append(numbers, i)
        fmt.Printf("After adding %d: len=%d, cap=%d\n", i, len(numbers), cap(numbers))
    }
    
    fmt.Println("Final slice:", numbers)
}
```

### Output demonstrates capacity growth:
```
Initial: len=0, cap=0
After adding 1: len=1, cap=1
After adding 2: len=2, cap=2
After adding 3: len=3, cap=4
After adding 4: len=4, cap=4
After adding 5: len=5, cap=8
After adding 6: len=6, cap=8
After adding 7: len=7, cap=8
After adding 8: len=8, cap=8
After adding 9: len=9, cap=16
After adding 10: len=10, cap=16
```

### Performance Tips:

- **Pre-allocate when size is known**: Use `make([]int, 0, expectedSize)` to avoid multiple reallocations
- **Batch operations**: Append multiple elements at once when possible
- **Monitor capacity**: Understanding growth patterns helps optimize memory usage

## 🧠 Summary

| Property | Description | Function | Example Result |
|----------|-------------|----------|----------------|
| Length | Number of elements in the slice | `len()` | 2 |
| Capacity | Max size from start index to end of array | `cap()` | 4 |

Let me know if you want visual diagrams or memory layout!
