## How to start in golang after installation today 16/10/2024

Create a path and run in terminal **go mod init** any_package_name with golang version 1.23.2, the standard is create a package name like you github repository, example.: **go mod init github.com/rafamaxber/hello-word**

To install all packages we can run in terminal **go mod tidy** inside the same path are your **go.mod** file.

### Interesting Docs
https://gobyexample.com
https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md

### First example

Create a **main.go** and inside the file write:

```go
package main

import (
	"fmt"
)

func main() {
  fmt.Println("Hello word")
}
```

To see the result, in terminal run **go run .** or **go run main.go**.
To develop using live reload, we can use **nodemon** a npm package running in terminal or we can use [**air**](https://github.com/cosmtrek/air) a go module to help develop websites in go.

```bash
nodemon --exec go run main.go --signal SIGTERM
```

We should to use `defer` when we need to call some `function` on final parent function execution.

```go
defer db.Close()
```

Example:

```go
func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}

  defer rows.Close() // would be executed after return categories, nil

	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}
```
