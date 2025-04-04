# Memory Leak
refers to a situation where memory that is no longer needed is not released, causing the application's memory usage to grow unnecessarily over time. Even though Go has automatic garbage collection, memory leaks can still occur due to how references are maintained.

```go
var globalSlice []*SomeStruct

func leakyFunction() {
    s := &SomeStruct{}
    globalSlice = append(globalSlice, s) // s is never garbage collected
}
```