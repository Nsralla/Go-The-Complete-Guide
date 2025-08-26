# Go: The Complete Guide 🧭

Welcome to **Go: The Complete Guide**, a structured and hands-on companion repository for learning the Go programming language from the ground up.

This repo includes complete module-wise exercises, source code, and examples that cover Go essentials, packages, pointers, and more advanced concepts — perfect for learners and developers seeking to master Go through practice.

---

## 📚 Topics Covered

Each folder corresponds to a progressive module in the course:

| Module | Title                          | Highlights                           |
|--------|--------------------------------|--------------------------------------|
| 01     | GO Essentials                  | Hello world, variables, functions    |
| 02     | Working with Packages          | Modularization, file operations      |
| 03     | Understanding Pointers         | Pointer logic with diagrams          |
| ...    | *(More modules included)*      | *(See folder structure below)*       |

---

## 🧩 Folder Structure

Go-The-Complete-Guide-main/
│
├── 01- GO Essentials/
│ ├── Go Essentials/
│ ├── Go Essentials Exercise/
│ └── Go Essentials Exercise 2/
│
├── 02- Working with Packages/
│ ├── bank.go
│ ├── communication.go
│ └── utilities/
│ └── fileops.go
│
├── 03- Understanding Pointers/
│ ├── image-1.png ... image-5.png
│ └── readme.md
│
├── ... more modules ...


---

## 🚀 Getting Started

### ✅ Prerequisites

- **Go 1.18+** installed  
  [Install Go →](https://go.dev/doc/install)

### 📦 Running Examples

You can run any module using:

```bash
cd "01- GO Essentials/Go Essentials"
go run app.go
```

To initialize dependencies for folders with go.mod:
``` bash
go mod tidy
go run app.go
```

⚠️ Some folders may contain .exe files (compiled on Windows). You can ignore or rebuild them from the .go source.
