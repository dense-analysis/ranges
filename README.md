# ranges

[![Go](https://img.shields.io/badge/pkg-%2311AB00.svg?style=for-the-badge&logo=go&labelColor=555555&logoColor=white)](https://pkg.go.dev/github.com/dense-analysis/ranges) [![CI](https://img.shields.io/github/actions/workflow/status/dense-analysis/ranges/go.yml?branch=master&style=for-the-badge&label=ci&logo=github)](https://github.com/dense-analysis/ranges/actions/workflows/go.yml?query=event%3Apush+branch%3Amaster)

`ranges` implements the nearest implementation of D range-based algorithms in
Go, for fun and to experiment with what is possible in Go 1.18 and above.

Ranges are a concept popular in D and C++ for creating a language by which
generic algorithms can be define which operate on potentially any type of
container or sequence of data. Instead of redundantly defining the same
algorithms over and over again for different types, you instead define how to
create a range that lazily evaluates a given container or sequence in
potentially many directions, and existing generic algorithms can be applied to
the range.

## Primitives

The library defines the following ranges primitives, as Go interfaces.

* `OutputRange[T any]` - Any type you can write the following operations:
  * `Put(element T) error` Write to the output range.
* `InputRange[T any]` - Anything iterable, with the following operations:
  * `Empty() bool` - Check if a range is empty.
  * `Front() T` - Get the current element. May panic if an empty range.
  * `PopFront()` - Remove the front element. May panic if an empty range.
* `ForwardRange[T any]` - An `InputRange[T]` with additional operations:
  * `Save() ForwardRange[T]` - Copy the range and its position.
* `BidirectionalRange[T any]` - `ForwardRange[T]` with additional operations:
  * `Back() T` - Get the current end element. May panic if an empty range.
  * `PopBack() T` - Remove the back/end element. May panic if an empty range.
  * `SaveB() BidirectionalRange[T]` - Save the position in both directions.
* `RandomAccessRange[T any]` - A `BidirectionalRange[T]` with additional
  operations:
  * `Get(index int) T` Return an element of the range. May panic if out of
    bounds.
  * `Len() int` - Return the length of the range.
  * `SaveR() RandomAccessRange[T]` - Save the position with random access.

The ranges library defines the following types, which may be used as constraints
for generic functions.

* `Signed` - Any signed integer primitive type.
* `Unsigned` - Any signed integer primitive type.
* `Integer` - Any integer primitive type.
* `Float` - Any floating point primitive type.
* `RealNumber` - Any integer or floating point primitive type.
* `Ordered` - Any primitive type that be ordered.

### Tuple Types

To express tuples in Go, there are different types declared for different
numbers of items. This library has the following:

* `Pair` - Holds 2 values of any mix of types.
* `Triplet` - Holds 3 values of any mix of types.
* `Quartet` - Holds 4 values of any mix of types.
* `Quintet` - Holds 5 values of any mix of types.
* `Sextet` - Holds 6 values of any mix of types.
* `Septet` - Holds 7 values of any mix of types.
* `Octet` - Holds 8 values of any mix of types.
* `Ennead` - Holds 9 values of any mix of types.
* `Decade` - Holds 10 values of any mix of types.

For convenience, every tuple has a `Make` function for creating them, and a
`Get()` function for returning the values as a Go native tuple, so tuples
can be and split into multiple values without redundantly naming types.

```go
// Inferred as Pair[int, string]
pair := MakePair(1, "hello")
// Split into int and string values.
num, str := pair.Get()
```

## Error Handling

Ranges other than `OutputRange` do not include `error` values as part of their
types. Algorithms in this library do not result in runtime errors. They may only
panic when input to the functions producing the ranges is invalid, or when
attempting to access elements that do not exist. When you wish to place values
that result in errors, make the errors an explicit part of the type of your
range such as `InputRange[Pair[T, error]]`.

Remember that in Go you can forward Go's native return tuples as arguments and
return values, and this integrates well with the ranges library tuple types.
This can make forwarding errors easier.

```go
func OtherFunc() int, error { /* ... */ }

func ReturnsValueAndError() int, error {
  // Create Pair[int, error]
  // Go can spread the multiple return into the arguments for us.
  pair := MakePair(OtherFunc())

  // Return both values.
  return pair.Get()
}
```

## Lazy Evaluation

All ranges are lazily-evaluated and compute values anew on demand. This means
each call to `Front()` or `Back()` on a range will return a new value of `T`.

Modifications to the returned objects may not be reflected in a container, such
as returning a copy of a struct instead of a pointer to it. When modification
to elements of an underlying container is necessary, you should create a range
of pointers, as in `*T`. This will permit modification of underlying values.

Because values are computed on the fly, a call to a callback function may be
executed each time `Front()` or `Back()` are called for ranges. You may wish to
cache the results of computation in a chain of ranges by calling one the `Cache`
functions, including `Cache`, `CacheF`, `CacheB`, and `CacheR`.

## Algorithms

Nearly all algorithms accepting or producing `InputRange` can accept or produce
a `ForwardRange` instead by calling a variation of the function with an `F`
suffix. Most algorithms that can accept or produce a `ForwardRange` can accept
or produce a `BidirectionalRange` with a `B` suffix. Some functions such as `Map` can be called with shortcuts for slices
with an `S` suffix.

* `operators`
  * `Lt` - Implements `a < b` for all orderable types.
  * `Le` - Implements `a <= b` for all orderable types.
  * `Eq` - Implements `a == b` for all comparable types.
  * `Ne` - Implements `a != b` for all comparable types.
  * `Ge` - Implements `a >= b` for all orderable types.
  * `Gt` - Implements `a > b` for all orderable types.
* `functional`
  * `Compose*` - Composes several functions in a sequence, such that
    `Compose*(f1, ..., fn)(x)` produces `f1(...(fn(x)))`
  * `Pipe*` - Pipes several functions in a sequence, such that
    `Pipe*(f1, ..., fn)(x)` produces `fn(...(f1(x)))`
* `ranges`
  * `Chain` - Produces an `InputRange` iterating over a slice of ranges.
  * `Chunks` - Takes `InputRange` chunks of a given size from an `InputRange`.
  * `Cycle` - Repeats a `ForwardRange` infinitely.
  * `Drop` - Drops up to a number of elements from the start of a range.
  * `Enumerate` - Yields elements with indexes (`Pair[int, T]`) from `0`.
  * `EnumerateN` - `Enumerate` with a provided start index.
  * `Flatten` - Produces an `InputRange` from a range of ranges.
  * `FlattenSB` - A special variation to flatten a slice of
    `BidirectionalRange`s into a `BidirectionalRange.
  * `FlattenSS` - A special variation to flatten a slice of slices into a
    `BidirectionalRange.
  * `FrontTransversal` - Produces an `InputRange` iterating over the first value
    of every non-empty range in a range of ranges.
  * `Generate` - Creates an infinite `BidirectionalRange` by calling a function
    repeatedly.
  * `Iota` - A `BidirectionalRange` producing values from `0` value up to and
    excluding an `end` value, incrementing by `1`.
  * `IotaStart` - `Iota` with a given `start` value to use in place of `0`.
  * `IotaStep` - `IotaStart` with a `step` value to use in place of `1`.
  * `Null` - Returns a `BidirectionalRange` that is always empty and consumes
    zero memory.
  * `Only` - Returns a `BidirectionalRange` through the arguments provided.
  * `PadRight` - Produces an `InputRange` with up to `count` items by padding
    the range with a given value.
  * `Repeat` - Creates an infinite `BidirectionalRange` repeating a value.
  * `Retro` - Returns a reversed `BidirectionalRange`.
  * `RoundRobin` - Produces an `InputRange` iterating over the first value of
    every non-empty range in a range of ranges in a cycle until all ranges are
    exhausted.
  * `Slide` - Produces a `ForwardRange` of chunks of a given `windowSize`
    stepping forward `1` element at a time.
  * `SlideStep` - `Slide` with a `stepSize` for stepping over elements.
  * `Stride` - Produces every `step` element in an `InputRange`.
  * `Take` - Takes up to a number of elements from a range.
  * `Tee` - Produces an `InputRange` that produces elements from a given range
    and outputs values to a given `OutputRange` when elements are popped.
  * `ZipN` - Produce a range stepping over `N` ranges in parallel. There are
    several `Zip` functions for different numbers of arguments.
* `output`
  * `AssignSink` - Creates an `OutputRange` that assigns values to a given
    `InputRange` of pointers by dereferencing the pointers.
  * `NullSink` - Creates an `OutputRange` that discards all data.
  * `SliceSink` - Creates an `OutputRange` that appends to the given slice.
* `slices`
  * `Bytes` - Produces `BidirectionalRange[byte]` from a `string` like `[]byte(s)`
  * `Runes` - Produces `BidirectionalRange[rune]` from a `string` like `[]rune(s)`
  * `SliceRange` - Produces `BidirectionalRange[T]` from `[]T`
  * `SliceRetro` - Produces `BidirectionalRange[T]` from `[]T` in reverse.
  * `SlicePtrRange` - Produces a `BidirectionalRange[*T]` from `[]T`
  * `SlicePtrRetro` - Produces a `BidirectionalRange[*T]` from `[]T` in reverse.
  * `Slice` - Produces `[]T` from `InputRange[T]`
  * `String` - Produces `string` from `InputRange[rune]`
* `comparison`
  * `Among` - Returns `true` if a `value` is equal to any of the `values`
     according to an `eq` callback.
  * `AmongEq` - Returns `true` if a `value` is equal to any of the `values`
     using a simple `==` comparison.
  * `Cmp` - Steps through two ranges comparing values with a `cmp` function and
    returns the result if it's nonzero. Returns `-1` or `1` if ranges are
    different lengths.
  * `CmpFunc` - Produces a comparison function for all types that support `<`
    and `>`.
  * `Equal` - Returns `true` if two ranges are equal according to a comparison
    defined in a callback.
  * `EqualComparable` - Returns `true` if two ranges are equal, element by
    element, for all comparable types.
  * `IsPermutation` - Returns `true` if two ranges are permutations of each
    other in `O(m + n)` time by allocating a temporary map.
  * `IsPermutationNoAlloc` - Returns `true` if two ranges are permutations of
    each other in `O(m * n)` without allocating a map.
  * `IsSameLength` - Checks if two ranges are the same length in `O(n)` time.
  * `Max` - Returns the maximum value among all values given as arguments.
  * `Min` - Returns the minimum value among all values given as arguments.
  * `Mismatch` - Eagerly advances all ranges until the first element is found
    where any two elements are not equal according to a callback.
* `iteration`
  * `Cache` - Caches results in an `InputRange` so `Front()` will be called only
    once per element on the original range.
  * `CacheF` - Caches results so `Front()` will only be called once per element
    on the original range, unless the range is saved and traversed over multiple
    times.
  * `CacheB` OR `CacheR` - Caching of `BidirectionalRange` and
    `RandomAccessRange` elements.
  * `ChunkBy` - Returns an `InputRange` that splits a range into sub-ranges
    when `cb(a, b)` returns `false`.
  * `ChunkByValue` Returns an `InputRange` that splits a range into sub-ranges
    where `cb(a) == c(b)`.
  * `Each` - Calls a callback with each value of a range.
  * `Exhaust` - Steps through every element of a range until it's empty.
    Similar to `Each` with an empty callback function, only `Front()` will never
    be called for the range.
  * `Filter` - Filter any `InputRange` with a callback.
  * `FilterB` - Filter a `BidirectionalRange`, producing a range that can be
    advanced in both directions. Less efficient for moving forwards, as it
    requires priming the range in both directions.
  * `Group` - Yields pairs of `(value, size)` counting how many values are equal
    in each group according to `cb(a, b)`.
  * `GroupComparable` - `Group` where `a == b` for any comparable value.
  * `Joiner` - Joins ranges with a `separator` `ForwardRange` between ranges.
  * `JoinerSS` - A special variation to join a slice of slices into a
    `ForwardRange`.
  * `JoinStrings` - A convenience function for creating a `string` from a
    `ForwardRange` of `string` values with a `separator` `string`.
  * `Map` - Map elements in any `InputRange` with a callback. The result of
    calling the callback is not stored, so use `Cache` when generating ranges
    with `Map`.
  * `Permutations` - Given a slice of values, produce a `ForwardRange` of
    all permutations of the given slice.
  * `ReduceNoSeed` - Eagerly reduces a range to a single value without a seed
    value. Panics when a range is empty.
  * `SplitWhen` - Splits a range where `cb(a, b) == true` for adjacent elements.
  * `Splitter` - Splits forward ranges with a `separator` `ForwardRange` between
    ranges where `cb(a, b) == true`
  * `SplitterSS` - A special variation to split a slice with a slice.
  * `SplitterComparable` - `Splitter` where `a == b`.
  * `SplitString` - A convenience function for splitting a `string` with
    a `string.
  * `Uniq` - Yields unique adjacent elements where `cb(a, b) == true`.
  * `UniqComparable` - `Uniq` where `a == b`.
* `mutation`
  * `Copy` - Copies all values from an `InputRange` into an `OutputRange`.
  * `Fill` - Assigns a value to all locations in a range of pointers.
  * `FillPattern` - Assigns a pattern of values from a `ForwardRange` to all
    locations in a range of pointers.
  * `StripLeft` - Removes elements where `cb(a) == true` from the front of a
    range.
  * `StripLeftComparable` - `StripLeft` where `a == value`, given a provided
    value.
  * `StripRight` - Removes elements where `cb(a) == true` from the back of a
    range.
  * `StripRightComparable` - Removes elements where `a == value` from the back
    of a range.
  * `Strip` - Removes elements where `cb(a) == true` from the front and back of
    a range.
  * `StripComparable` - Removes elements where `a == value` from the front and
    back of a range.
* `searching`
  * `All` - Checks if all elements in an `InputRange` satisfy a callback.
  * `Any` - Checks if any elements in an `InputRange` satisfy a callback.
  * `CanFind` - `Find`, but simply returns `true` if anything can be found.
  * `CanFindComparable` - `FindComparable` but simply returns `true` if anything
    can be found.
  * `Length` - Returns the number of elements in an `InputRange` in `O(n)` time.
  * `Count` - Returns the number of elements in an `InputRange` where the
    callback returns `true`
  * `CountUntil` - Returns the number of elements in an `InputRange` until the
    callback returns `true`.
  * `CommonPrefix` - Returns a `FowardRange` over the common prefix of two
    ranges. The first range must be a `ForwardRange`.
  * `DropWhile` - Advances a range while a callback returns `true`.
  * `Find` - Advances a range until `cb(x, needle)` returns `true`, comparing
    elements with a `needle`.
  * `FindComparable` - Advances a range until `x == needle`.
  * `FindEqual` - Advances a range until `cb(a, b)` returns `true` for every
    element of a `needle` range.
  * `FindEqualComparable` - Advances a range until `a == b` is satisfied for
    all elements of a `needle` range.
  * `FindAdjacent` - Advances a range until `cb(a, b)` returns `true` for two
    adjacent elements in a `ForwardRange`.
  * `FindAdjacentComparable` - Advances a range until `a == b` is satisfied for
    two adjacent elements in a `ForwardRange`.
  * `FindAmong` - Advances a range until `cb(a, b)` returns `true` for any
    element of a `needle` range.
  * `FindAmongComparable` - Advances a range until `a == b` is satisfied for any
    element of a `needle` range.
  * `SkipOver` - Skips over elements in a `haystack` `ForwardRange` if the range
    starts with a `needle` range, according to `cb(a, b) == true`.
  * `StartsWith` - Checks if one `InputRange` starts with another one, through a
    callback.
  * `TakeWhile` - Advances a range until the callback returns `true`.
* `setops`
  * `CartesianProduct` - Computes the `Cartesian` product of a series of forward
    ranges.

## Implementing New Ranges

New Ranges can be implemented by creating a struct with pointer receivers for
all of the methods needed to satisfy a type of range. For example:

```go
// A pointer to this is an InputRange[T] because of the implementation below.
type newRangeType[T any] struct {
  // Add whatever data you need.
}

func (r *newRangeType[T]) Empty() bool { /* ... */ }
func (r *newRangeType[T]) Front() T    { /* ... */ }
func (r *newRangeType[T]) PopFront()   { /* ... */ }

func MakeNewRangeType[T any](/* ... */) InputRange[T] {
  return &newRangeType{/* ... */}
}
```

It's important to return pointers to your structs so ranges can be freely
passed around as a reference type.

## Performance

Wherever possible, algorithms will attempt to minimize allocations. The `Len()`
of any `RandomAccessRange` objects at runtime, or the `len()` of slices may be
used to determine how much memory to allocate, or to optimize algorithms to
avoid unnecessary operations.

Go's compiler is capable of inlining many range function calls, which will
reduce overhead. Ranges are likely to be stored in the garbage-collected heap.
The performance of using ranges will therefore be hard to predict. You should
benchmark and debug your code, and add optimizations where appropriate.
Significant improvements easily obtained are never premature.

## Limitations

### No function overloading for generic types

It's not possible to implement a `Filter` function that returns
`ForwardRange[T]` if given `ForwardRange[T]` and returns `InputRange[T]` if
given `InputRange[T]`.

The only available alternative in Go is to write multiple function names for
each range type, such as `FilterF` for `ForwardRange[T]` and `Filter`
for `InputRange[T]`.

### No covariant return types in Go

We need `SaveB` and `SaveR` methods because there are no covariant return types
in Go. This means a user can be led to choosing a suboptimal save function.
There's no way around this in Go. In D, saving a range automatically carries
across the more specific details of the range type, such as `RandomAccessRange`
`save` method returning a `RandomAccessRange`.

### No tuple types and variadic generics

Variadic generic types are difficult to implement correctly, and Go does not
have generic tuple types to represent any combination of values. This means that
just as in classic Java generic interfaces, functions that take different
combinations of types must be redundantly defined up to some maximum number of
arguments in a programmer's imagination of how many arguments someone will call
a function with. For example:

```go
func Args0()[] { }
func Args1[V1 any](v1 V1) { }
func Args2[V1 any, V2](v1 V1, v2 V2) { }
func Args3[V1 any, V2 any, V3 any](v1 V1, v2 V2, v3 V3) { }
func Args3[V1 any, V2 any, V3 any, V4 any](v1 V1, v2 V2, v3 V3, v4 V4) { }
```

It's not possible to write something like the following for all of the above:

```go
func Args[V any...](v V...) { }
```

Because of this, multiple variations of the `Zip` function are required.

## Advantages of Go

### Exceptions must be an explicit part of the type of ranges

Go does not have exceptions outside of `panic`. Ranges that can yield errors at
any point have to explicitly include them in the type of result returned by
`Front()`, so all ranges should be considered exception safe as long as the
contract of calling `Empty()` before `Front()` or `PopFront()` is not broken.
