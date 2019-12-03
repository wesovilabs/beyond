# Changelog

## [v1.0.0] - 02/12/2019

### Added
- [#22](https://github.com/wesovilabs/beyond/issues/22) `Beyond usage` is
shown when the entered command is not valid. 
- [#7](https://github.com/wesovilabs/beyond/issues/7) This new version provide a method named `Ignore`
which can be used to ignore methods to be intercepted.
- [#5](https://github.com/wesovilabs/beyond/issues/5) `Beyond configuration` can be loaded from a
toml file.
- [#6](https://github.com/wesovilabs/beyond/issues/6). It's possible to skip joinpoint invocation from the
Before method.  
- [#3](https://github.com/wesovilabs/beyond/issues?q=is%3Aissue+is%3Aclosed+milestone%3Av0.0.2+label%3Aenhancement)
Docker image `wesovilabs/beyond:1.0.0` is provided.
- [#8](https://github.com/wesovilabs/beyond/issues/8) It special character `?` is supported for optional objects

### Changed
- [#26](https://github.com/wesovilabs/beyond/issues/26) Library has been renamed from `goa` to `beyond`. 
It's due an existing and known library already had that name.
- [#10](https://github.com/wesovilabs/beyond/issues/10) This change improves the `Beyond performance` since
the list of functions to be evaluated is less than It used to be required in older versions.

### Fixed
- [#23](https://github.com/wesovilabs/beyond/issues/23) Fix build command when flag -o is not provided.
- [#14](https://github.com/wesovilabs/beyond/issues/14) Fixing windows compilation error

## [v0.0.1] - 24/11/2019 

🎉 first release! 
