Counter
=======

## Description

A set of counters which helps avoid reinvented wheels.

## Usage

### Installation

`go get "github.com/sam-caldwell/counter/v2"`

### Tests

`make test`

### Features

#### `simpleCounter()`

> Just a simple counter with an increment and decrement capability.
> Returns current value then increments the internal state.

```golang
var count SimpleCounter
count.Increment() //output: 0
count.Increment() //output: 1
count.Increment() //output: 2
```

`countOk(ok bool)` - returns incremented count if input is `true` or `-1` if false
