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

## 🧠 Summary

| Property | Description | Function | Example Result |
|----------|-------------|----------|----------------|
| Length | Number of elements in the slice | `len()` | 2 |
| Capacity | Max size from start index to end of array | `cap()` | 4 |

Let me know if you want visual diagrams or memory layout!
