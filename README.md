![](https://img.shields.io/github/workflow/status/arata-nvm/Solitude/Go?style=for-the-badge)
![](https://img.shields.io/codecov/c/github/arata-nvm/Solitude?style=for-the-badge)
# Solitude

A compiled programming language

## Example
```
func main() {
  print(fib(41))
  return 0
}

func fib(n) {
  if n <= 1 {
    return n
  }
  return fib(n - 1) + fib(n - 2)
}
```

More examples can be found [here](https://github.com/arata-nvm/Solitude/tree/master/examples).

## Features

### Language Features
- [x] variables
- [x] functions
- [ ] modules
- [x] if / else / then
- [x] for
- [x] while

### Types
- [x] int
- [ ] string
- [ ] struct
- [ ] array
- [ ] map
- [ ] bool
- [ ] func

## Dependencies
- Clang == 9.x

## Development

### Building from source
1. `git clone https://github.com/arata-nvm/Solitude && cd Solitude`
2. `make`

### Compiling a Solitude program
1. `./bin/solitude -O <filename>`

## License
MIT
