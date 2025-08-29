# Notes & Todo Application

A command-line Go application demonstrating interfaces, generic code, and modular design patterns. This project allows users to create and manage notes and todo items, showcasing Go's interface system and clean architecture principles.

## 📋 Project Overview

This application is a learning project that demonstrates key Go programming concepts:

- **Interfaces**: Defines common behavior contracts (`Saver`, `Outputter`)
- **Generic Programming**: Uses Go's `any` type for flexible function parameters
- **Modular Design**: Separates concerns into distinct packages (`note`, `todo`)
- **JSON Serialization**: Automatically saves data as JSON files
- **Error Handling**: Comprehensive error checking and user feedback
- **File I/O**: Reads user input and writes structured data to files

The application allows users to create notes with titles and content, as well as simple todo items. All data is automatically saved to JSON files for persistence.

## 🏗️ Architecture

### Project Structure
```
05-Interfaces & Generic Code/
├── main.go          # Main application entry point
├── go.mod           # Go module definition
├── note/
│   └── note.go      # Note struct and methods
└── todo/
    └── todo.go      # Todo struct and methods
```

### Key Components

#### Interfaces
- **`Saver`**: Defines the contract for objects that can save themselves
- **`Outputter`**: Embeds `Saver` and adds display functionality

#### Data Types
- **`Note`**: Contains title, content, and creation timestamp
- **`Todo`**: Contains simple text-based tasks

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- Command line interface (Terminal/PowerShell)

### Installation

1. **Clone or download the project**
   ```bash
   git clone <repository-url>
   cd "05-Interfaces & Generic Code"
   ```

2. **Verify Go installation**
   ```bash
   go version
   ```

3. **Build the application**
   ```bash
   go build -o notes-app .
   ```

## 💻 Usage Guide

### Running the Application

1. **Start the application**
   ```bash
   go run .
   ```
   or if built:
   ```bash
   ./notes-app
   ```

2. **Follow the interactive prompts**
   ```
   Enter note title: My First Note
   Enter note content: This is the content of my first note
   ```

### Example Session

```bash
$ go run .
Enter note title: Meeting Notes
Enter note content: Discuss project timeline and milestones
Title: Meeting Notes
Content: Discuss project timeline and milestones
Created At: Thu, 29 Aug 2025 21:37:00 UTC
 --------------------
Text: Discuss project timeline and milestones
 --------------------
Hello, World!
42
{Meeting Notes Discuss project timeline and milestones 2025-08-29 21:37:00.123456 +0000 UTC}
{Discuss project timeline and milestones}
```

### Generated Files

The application automatically creates JSON files:

- **Note files**: `<note_title>.json` (e.g., `meeting_notes.json`)
- **Todo file**: `todo.json`

#### Example Note JSON Structure
```json
{
  "title": "Meeting Notes",
  "content": "Discuss project timeline and milestones",
  "created_at": "2025-08-29T21:37:00.123456Z"
}
```

#### Example Todo JSON Structure
```json
{
  "text": "Discuss project timeline and milestones"
}
```

## 🔧 Code Examples

### Creating a Custom Type with Interfaces

```go
// Any type implementing Saver must have a Save() method
type Saver interface {
    Save() error
}

// Outputter embeds Saver and adds display functionality
type Outputter interface {
    Saver
    ToString()
}
```

### Using Generic Functions

```go
// Function that accepts any type
func printAny(value any) {
    fmt.Println(value)
}

// Usage examples
printAny("Hello, World!")  // string
printAny(42)               // int
printAny(myNote)           // custom struct
```

### Interface Implementation

```go
// Note implements both Save() and ToString() methods
func (note Note) Save() error {
    fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
    json, err := json.Marshal(note)
    if err != nil {
        return err
    }
    return os.WriteFile(fileName, json, 0644)
}

func (note Note) ToString() {
    fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", 
        note.Title, note.Content, note.CreatedAt.Format(time.RFC1123))
}
```

## 🛠️ Development

### Adding New Data Types

To add a new data type that works with the existing interface system:

1. Create a new package directory
2. Implement the `Save()` method
3. Implement the `ToString()` method
4. Use the `outputData()` function in main

### Extending Functionality

- **Add validation**: Enhance the `New()` functions with more validation rules
- **Add editing**: Implement update functionality for existing notes/todos
- **Add listing**: Create functions to list and search existing items
- **Add deletion**: Implement safe deletion of notes and todos

## 📚 Learning Concepts Demonstrated

1. **Interface Segregation**: Small, focused interfaces
2. **Composition over Inheritance**: Interface embedding
3. **Error Handling**: Proper error propagation and handling
4. **Package Organization**: Logical separation of concerns
5. **JSON Marshaling**: Struct tags and serialization
6. **File Operations**: Safe file writing with error handling
7. **User Input**: Reading and processing command-line input
8. **Generic Programming**: Type-safe generic functions

## 🤝 Contributing

This is a learning project. Feel free to:
- Add new features
- Improve error handling
- Add unit tests
- Enhance the user interface
- Add configuration options

## 📄 License

This project is for educational purposes and is part of "Go The Complete Guide" course materials.

---

*This project demonstrates Go's powerful interface system and how it enables clean, modular, and extensible code design.*
