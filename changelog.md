# Changelog

## v1.0.0

* Raised the minimum supported Go version to 1.21, which offers much better
  type inference, and make the library easier to maintain.
* `HasLength` has been removed.
* The `B`, `F`, and `I` functions have been removed.

## v0.4.0

* Added support for `RandomAccessRange` across the board.
  * The addition of random access ranges may still break code due to weak
    type inference in Go versions below 1.21.
* The unused `HasLength` interface is deprecated.
* `B`, `F`, and `I` functions are not marked deprecated, but will be removed
  when the library is updated to support only Go 1.21.

## v0.3.0

* Added `Get` functions for turning any fixed tuple type into a native tuple.

## v0.2.0

* Added support for `BidirectionalRange` across the board.
* Now features `StripRight` and `Strip` algorithms, that require
  `BidirectionalRange`.

## v0.1.1

* Fixed package layout

## v0.1.0

* Initial version with only `InputRange` and `ForwardRange` supported.
