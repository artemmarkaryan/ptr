# Ptr — useful pointer-related funcs, which go lacks so much

I got tired of implementing these two method in every project of mine.

## Usage

```go
// get pointer to time.Now()
now := ptr.P(time.Now()) // -> *time.Time
```

```go
// get value of pointer
nilValue := (nil).(*int)
zeroValue := ptr.V(nilValue)  // -> 0
```

