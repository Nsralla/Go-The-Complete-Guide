# Go: The Complete Guide 🧭

Welcome to **Go: The Complete Guide**, a structured and hands-on companion repository for learning the Go programming language from the ground up.

This repo includes complete module-wise exercises, source code, and examples that cover Go essentials, packages, pointers, and more advanced concepts — perfect for learners and developers seeking to master Go through practice.

---

## 📚 Topics Covered

Each folder corresponds to a progressive module in the course:

| Module | Title                                    | Highlights                                      |
|--------|------------------------------------------|-------------------------------------------------|
| 01     | GO Essentials                           | Hello world, variables, functions, basics       |
| 02     | Working with Packages                   | Modularization, file operations, utilities      |
| 03     | Understanding Pointers                  | Pointer logic with visual diagrams              |
| 04     | Practice Project                        | Note management system                          |
| 04     | Structs And Custom Types               | Custom types, struct methods, user models       |
| 05     | Generic                                | Generic programming concepts                     |
| 05     | Interfaces & Generic Code              | Interface design, generic implementations        |
| 06     | Exercise                               | Hands-on coding challenges                       |
| 06     | Arrays, Slices and Maps                | Data structures, collections management         |
| 07     | Functions Deep Dive                    | Advanced functions, closures, recursion         |
| 08     | Project                                | Tax calculator with concurrent processing       |
| 09     | Concurrency                            | Goroutines, channels, parallel processing       |
| 10     | REST API Project                       | Full-stack API with authentication & database   |
| 11     | Follow up                              | Advanced topics and best practices              |

---

## 🧩 Folder Structure

```
Go-The-Complete-Guide/
│
├── 01- GO Essentials/
│   ├── Go Essentials/                    # Basic Go syntax and concepts
│   ├── Go Essentials Exercise/           # Practice exercises
│   └── Go Essentials Exercise 2/         # Additional challenges
│
├── 02- Working with Packages/
│   ├── bank.go                          # Banking operations
│   ├── communication.go                 # I/O operations
│   └── utilities/fileops.go             # File management utilities
│
├── 03- Understanding Pointers/
│   ├── pointers.go                      # Pointer concepts and examples
│   └── [image-1.png ... image-5.png]   # Visual pointer diagrams
│
├── 04- Practice Project/
│   └── note/                            # Note management system
│
├── 04- Structs And Custom Types/
│   ├── structs.go                       # Struct definitions and methods
│   └── user/user.go                     # User model implementation
│
├── 05- Generic/
│   └── genric.go                        # Generic programming examples
│
├── 05- Interfaces & Generic Code/
│   ├── main.go                          # Interface implementations
│   ├── note/note.go                     # Note interface
│   └── todo/todo.go                     # Todo interface
│
├── 06- Exercise/
│   └── main.go                          # Coding challenges
│
├── 06- Managing Related Data with Arrays, Slices and Maps/
│   ├── lists.go                         # Array and slice operations
│   ├── maps.go                          # Map data structure examples
│   └── lists/                           # Advanced list operations
│
├── 07- Functions Deep Dive/
│   ├── main.go                          # Main function examples
│   ├── anonyms/                         # Anonymous functions
│   ├── closure/                         # Closure concepts
│   ├── funcs/                           # Function variations
│   └── recursion/                       # Recursive algorithms
│
├── 08- Project/
│   ├── main.go                          # Tax calculator main
│   ├── cmdmanager/                      # Command management
│   ├── conversion/                      # Data conversion utilities
│   ├── filemanager/                     # File I/O operations
│   ├── iom/                             # Input/output management
│   └── job/                             # Concurrent job processing
│
├── 09- Concurrency Running Tasks In Parallel/
│   ├── main.go                          # Concurrency examples
│   └── Practice/                        # Hands-on concurrency exercises
│
├── 10- Course Project Build a REST API incl Authentication SQL Database/
│   ├── main.go                          # API server entry point
│   ├── api.db                           # SQLite database
│   ├── db/                              # Database operations
│   ├── middleware/                      # Authentication middleware
│   ├── models/                          # Data models (Event, User)
│   ├── routes/                          # API route handlers
│   ├── utils/                           # JWT, hashing utilities
│   └── api-test/                        # HTTP test files
│
└── 11- Follow up/
    └── main.go                          # Advanced Go concepts
```


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
```bash
go mod tidy
go run main.go
```

### 🏗️ Key Projects

**Tax Calculator (Module 08)**
```bash
cd "08- Project"
go run main.go
```

**REST API Server (Module 10)**
```bash
cd "10- Course Project Build a REST API incl Authentication SQL Database"
go mod tidy
go run main.go
```

### 🧪 Testing the API

Use the provided HTTP test files in `api-test/` folder:
- `create-user.http` - User registration
- `login.http` - User authentication
- `create-event.http` - Event creation
- `get-events.http` - Fetch all events

⚠️ **Note:** Some folders may contain .exe files (compiled on Windows). You can ignore or rebuild them from the .go source.

---

## 🎯 Learning Path

1. **Start with basics** → GO Essentials (Module 01)
2. **Learn modularization** → Working with Packages (Module 02)
3. **Master memory management** → Understanding Pointers (Module 03)
4. **Build data structures** → Structs and Custom Types (Module 04)
5. **Advanced concepts** → Interfaces & Generics (Module 05)
6. **Data handling** → Arrays, Slices, Maps (Module 06)
7. **Function mastery** → Functions Deep Dive (Module 07)
8. **Real-world project** → Tax Calculator (Module 08)
9. **Concurrency** → Parallel Processing (Module 09)
10. **Full-stack development** → REST API Project (Module 10)
11. **Advanced topics** → Follow up (Module 11)

---

## 🛠️ Technologies Used

- **Go 1.18+** - Main programming language
- **SQLite** - Database for REST API project
- **JWT** - Authentication tokens
- **Goroutines** - Concurrency implementation
- **JSON** - Data serialization
- **HTTP** - REST API development

---

## 📝 License

This project is for educational purposes. Feel free to use and modify for learning Go programming.
