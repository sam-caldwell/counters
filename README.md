Counter
=======

# Description

A set of simple counters that we can use rather than reinventing wheels.

`simpleCounter()` - just a counter. It increments and returns the new value.

```golang
counter:= simpleCounter()

counter() //output: 1
counter() //output: 2
counter() //output: 3
```


`countOk(ok bool)` - returns incremented count if input is `true` or `-1` if false
