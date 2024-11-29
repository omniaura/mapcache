# mapcache

[![Godoc Reference](https://godoc.org/github.com/omniaura/mapcache?status.svg)](http://godoc.org/github.com/omniaura/mapcache)
[![Go Coverage](https://github.com/omniaura/mapcache/wiki/coverage.svg)](https://raw.githack.com/wiki/omniaura/mapcache/coverage.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/omniaura/mapcache)](https://goreportcard.com/report/github.com/omniaura/mapcache)

A type-safe, concurrent in-memory key-value cache for Go with TTL support.

## Features

- 🔒 Thread-safe operations
- 📦 Generic type support
- ⏰ TTL (Time To Live) support
- 🧹 Automatic cleanup of expired entries
- 🔄 Iterator support for cache entries
- 💪 Zero external dependencies

## Installation

```bash
go get github.com/omniaura/mapcache
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/omniaura/mapcache"
)

func main() {
    // Create a new string->int cache
    cache, _ := mapcache.New[string, int]()

    // Get or compute a value
    value, err := cache.Get("mykey", func() (int, error) {
        // This function is only called if the value isn't cached
        // or has expired
        return 42, nil
    })
    
    if err != nil {
        panic(err)
    }
    fmt.Println(value) // Output: 42
    
    // Subsequent calls will return the cached value
    value, _ = cache.Get("mykey", func() (int, error) {
        // This won't be called since the value is cached
        return 100, nil
    })
    fmt.Println(value) // Still outputs: 42
}
```

### With TTL (Time To Live)

```go
package main

import (
    "time"
    "github.com/omniaura/mapcache"
)

func main() {
    // Create cache with 5 minute TTL
    cache, _ := mapcache.New[string, int](
        mapcache.WithTTL(5 * time.Minute),
    )

    // Values will expire after 5 minutes
    cache.Get("key", func() (int, error) {
        return 42, nil
    })
    
    // You can also override TTL per-request
    cache.Get("key", func() (int, error) {
        return 42, nil
    }, mapcache.WithTTL(10 * time.Second))
}
```

### With Automatic Cleanup

```go
package main

import (
    "context"
    "time"
    "github.com/omniaura/mapcache"
)

func main() {
    ctx := context.Background()
    
    // Create cache with TTL and cleanup every minute
    cache, _ := mapcache.New[string, int](
        mapcache.WithTTL(5 * time.Minute),
        mapcache.WithCleanup(ctx, time.Minute),
    )

    // Expired entries will be automatically removed every minute
}
```

### Pre-allocated Size

```go
package main

import "github.com/omniaura/mapcache"

func main() {
    // Create cache with pre-allocated size
    cache, _ := mapcache.New[string, int](
        mapcache.WithSize(100),
    )
}
```

### Iterating Over Cache Entries

```go
package main

import (
    "fmt"
    "github.com/omniaura/mapcache"
)

func main() {
    cache, _ := mapcache.New[string, int]()
    
    // Add some values
    cache.Get("one", func() (int, error) { return 1, nil })
    cache.Get("two", func() (int, error) { return 2, nil })
    
    // Sequential iteration
    for k, v := range cache.All() {
        fmt.Printf("Key: %s, Value: %d\n", k, v.V)
    }
    
    // Parallel iteration (for concurrent processing)
    for k, v := range cache.AllParallel() {
        fmt.Printf("Processing: %s=%d\n", k, v.V)
    }
}
```

### Error Handling

```go
package main

import (
    "errors"
    "github.com/omniaura/mapcache"
)

func main() {
    cache, _ := mapcache.New[string, int]()
    
    // Handle errors from the update function
    value, err := cache.Get("key", func() (int, error) {
        return 0, errors.New("failed to compute value")
    })
    
    if err != nil {
        // Handle error
    }
}
```
