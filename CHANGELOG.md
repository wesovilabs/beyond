# Changelog

## [v0.0.1] - 24/11/2019 

🎉 first release!

## [v1.0.0] - 

### Added
- [#22](Goa should show the help when introduced commands are not valid) `Beyond usage` is
shown when the entered command is not valid. 
- [#7](https://github.com/wesovilabs/beyond/issues/7) This new version provide a method named `Exclude`
which can be used to ignore methods to be intercepted.
```go
func Beyond() *api.Beyond {
  return api.New().
    WithBefore(advice.NewTracingAdvice, "*.*(...)...").
    WithBefore(advice.NewTracingAdviceWithPrefix("[beyond]"), "greeting.Bye(...)...").
    Exclude("advice.*(...)...")
}
```
-[#5](https://github.com/wesovilabs/beyond/issues/5) `Beyond configuration` can be loaded from a
file. By default `beyond.toml` will be loaded. It can be overrided 
**beyond.toml**

```toml
project="github.com/wesovilabs/beyond-examples/settings"
outputDir="generated"
verbose=true
work=true
excludes=[
    "go.sum",
    "vendor",
    ".git"
]
```
```sh
beyond --config my-beyond.toml run cmd/main.go
```
- [#3](https://github.com/wesovilabs/beyond/issues?q=is%3Aissue+is%3Aclosed+milestone%3Av0.0.2+label%3Aenhancement)
Dokcer image `wesovilabs/beyond:1.0.0` is provided.

### Changed
- [#26](https://github.com/wesovilabs/beyond/issues/26) Library has been renamed from `goa` to `beyondd`. 
It's due an existing and known library already had that name.
- [#10](https://github.com/wesovilabs/beyond/issues/10) This change improves the `Beyond performance` since
the list of functions to be evaluated is less than It used to be required in older versions.

### Removed

### Fixed
- [#23](https://github.com/wesovilabs/beyond/issues/23) Fix build command when flag -o is not provided.


### Deprecated

 
