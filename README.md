# Culer
This is a cli program that listens to std in and if the text matches a
substring then we wrap the text with a color and print to std out

## Example usage
Pipe a running program to Culer
```bash
go run data/runner.go | go run main.go
```

Pipe a file to Culer

```bash
cat main.go | go run main.go
```
