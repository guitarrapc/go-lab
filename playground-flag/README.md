## description

you can change argument order with `flag`.
However there are some basics you should know to use flag.

## simeple rule

* must : set `flaged args` then `non-flaged args`. otherwise behind non-flaged args will not recognized as flaged args.
  * OK (3 3): `go run .\flag-sample-count\main.go -int 1 -bool=true -str hogemoge a b c`
  * NO (8 0): `go run .\flag-sample-count\main.go a b c -int 1 -bool=true -str hogemoge`
* must : set flaged bool arg as `key=bool`. otherwise behind `bool` argument won't recognized as flaged args.
  * OK (0 3): `go run .\flag-sample-count\main.go -int 1 -bool=true -str hogemoge`
  * NO (3 2): `go run .\flag-sample-count\main.go -int 1 -bool true -str hogemoge`

## run example

```
go run ./noflag/main.go a b c
go run ./flag-sample/main.go -int 1 -bool=false -str fuga
go run ./flag-sample-typed/main.go -int 1 -bool=true -str hogemoge
go run ./flag-sample-count/main.go -int 1 -bool=true -str hogemoge
```
