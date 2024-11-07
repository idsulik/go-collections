# Contributing to go-collections

First off, thank you for considering contributing to go-collections! It's people like you who make go-collections such a great tool.

## Code of Conduct

By participating in this project, you are expected to uphold our Code of Conduct:

- Use welcoming and inclusive language
- Be respectful of differing viewpoints and experiences
- Gracefully accept constructive criticism
- Focus on what is best for the community
- Show empathy towards other community members

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the issue list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

1. **Use a clear and descriptive title**
2. **Describe the exact steps to reproduce the problem**
3. **Provide specific examples to demonstrate the steps**
4. **Describe the behavior you observed after following the steps**
5. **Explain which behavior you expected to see instead and why**
6. **Include your Go version (`go version`)**
7. **Include any relevant code snippets or test cases**

```go
// Example of a good bug report code snippet
ds := disjointset.New[int]()
ds.MakeSet(1)
ds.MakeSet(2)
ds.Union(1, 2)

// Expected: ds.Connected(1, 2) == true
// Actual: ds.Connected(1, 2) == false
```

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

1. **Use a clear and descriptive title**
2. **Provide a step-by-step description of the suggested enhancement**
3. **Provide specific examples to demonstrate the steps**
4. **Describe the current behavior and explain which behavior you expected to see instead**
5. **Explain why this enhancement would be useful**
6. **List some other libraries or applications where this enhancement exists**

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests
3. If you've changed APIs, update the documentation
4. Ensure the test suite passes
5. Make sure your code follows the existing style
6. Issue that pull request!

## Development Setup

1. Install Go 1.18 or higher (for generics support)
2. Fork and clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-collections.git
   cd go-collections
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run tests:
   ```bash
   go test ./...
   ```

## Styleguides

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

### Go Code Styleguide

* Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
* Write idiomatic Go code:
  ```go
  // Good
  var items []string
  for _, item := range collection {
      items = append(items, item)
  }

  // Bad
  items := make([]string, 0)
  for i := 0; i < len(collection); i++ {
      items = append(items, collection[i])
  }
  ```
* Use meaningful variable names
* Add comments for exported functions and types
* Maintain existing code style
* Use `go fmt` before committing
* Handle errors appropriately
* Write table-driven tests

### Documentation Styleguide

* Use [Markdown](https://guides.github.com/features/mastering-markdown/) for documentation
* Reference functions and types with backticks: `Queue.Enqueue()`
* Include code examples for new features
* Document performance characteristics
* Update README.md with any necessary changes
* Add doc comments for all exported functions and types:
  ```go
  // Push adds an item to the top of the stack.
  // Returns false if the stack is full.
  func (s *Stack[T]) Push(item T) bool {
      // Implementation
  }
  ```

## Testing

* Write test cases for all new functionality
* Use table-driven tests where appropriate:
  ```go
  func TestStack(t *testing.T) {
      tests := []struct {
          name     string
          input    []int
          expected int
      }{
          {"single item", []int{1}, 1},
          {"multiple items", []int{1, 2, 3}, 3},
          // More test cases...
      }

      for _, tt := range tests {
          t.Run(tt.name, func(t *testing.T) {
              // Test implementation
          })
      }
  }
  ```
* Aim for high test coverage
* Include both positive and negative test cases
* Test edge cases and error conditions

## Questions?

Feel free to [open an issue](https://github.com/idsulik/go-collections/issues/new) if you have any questions about contributing!

## License

By contributing, you agree that your contributions will be licensed under the MIT License.