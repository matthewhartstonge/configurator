# Changelog

## [0.2.0](https://github.com/matthewhartstonge/configurator/compare/v0.1.0...v0.2.0) (2024-09-13)


### Features

* **flag:** cli flag support. ([#13](https://github.com/matthewhartstonge/configurator/issues/13)) ([692ceef](https://github.com/matthewhartstonge/configurator/commit/692ceefb0d66e9f63756e3ae2b9788f1e8872619))

## 0.1.0 (2024-06-19)


### Features

* **.github:** `ConfigParser` adds `Type() string` for informing the user which parser is being used. ([#4](https://github.com/matthewhartstonge/configurator/issues/4)) ([e2b675e](https://github.com/matthewhartstonge/configurator/commit/e2b675ef25130b11443593444faf8eed0f667935))
* a lot of new goodies. ([f322f89](https://github.com/matthewhartstonge/configurator/commit/f322f898aabe3752710ce69b58032b4e58f5fe66))
* adds MIT license. ([98b86f3](https://github.com/matthewhartstonge/configurator/commit/98b86f3465af26e1150b5f70f4eae67c9d43fd67))
* **ConfigFileType:** implements `ConfigParser` return of filepath for `Parse`. ([e0b0b9b](https://github.com/matthewhartstonge/configurator/commit/e0b0b9b9ce02f62cc2d0450440b01726e3b709c9))
* **ConfigFileType:** makes naming consistent and simplifies implementation due to functionality provided by `ConfigType`. ([2b711f6](https://github.com/matthewhartstonge/configurator/commit/2b711f6b90653bfcfab0c03b9cee5c940b875d42))
* **ConfigImplementer:** `Validate` now takes in a `diag.Component`. ([247a51e](https://github.com/matthewhartstonge/configurator/commit/247a51e39bc17468c4b2b2c9493f0c3af8a9ef40))
* **ConfigParser:** `Parse` now returns filepath and `Values` returns the state of the parsed configuration. ([ee7e7d0](https://github.com/matthewhartstonge/configurator/commit/ee7e7d07ed539b9a0cb6188d6a043a1169d9d463))
* **ConfigTypeable:** now requires implementing `fmt.Stringer`. ([a5d73d7](https://github.com/matthewhartstonge/configurator/commit/a5d73d79671dd72f9f909a6f915a4c0e679d2e37))
* **ConfigType:** implements `ConfigParser.Values`. ([4e7a860](https://github.com/matthewhartstonge/configurator/commit/4e7a860bd88634d82a15d5c67f54079e54067194))
* **ConfigType:** simplifies dependant use by implementing base functionality. ([e76bdbc](https://github.com/matthewhartstonge/configurator/commit/e76bdbc51677440d396fb3b08db3cbac96cc4fd2))
* **configurator:** `FileParser` renamed to `ConfigFileParser`, implements `Stat` with diagnostic support. ([23a6378](https://github.com/matthewhartstonge/configurator/commit/23a6378a03ad7599bbfa2dffad15c69be330d42c))
* **diag:** adds `FromComponent` to enable logging on a passed in component. ([87723d4](https://github.com/matthewhartstonge/configurator/commit/87723d495a84e29be28c6f8828812678ecae7f1d))
* **diag:** adds diag builder for a developer friendly API. ([07cc7cb](https://github.com/matthewhartstonge/configurator/commit/07cc7cb68e94e2917db383e858b97c7482a71dea))
* **diag:** builder now returns `Diagnostics` instead of `*Diagnostics`. ([c76038d](https://github.com/matthewhartstonge/configurator/commit/c76038d5801d12e68a27584a10eab4dcd2377b9b))
* **diag:** reverses the parameters for builder as summary should be smaller than detail. ([35b507e](https://github.com/matthewhartstonge/configurator/commit/35b507e80b43e8827b8010a380e2e4243046f753))
* **diags:** initial diagnostics implementation. ([68f3319](https://github.com/matthewhartstonge/configurator/commit/68f33196de8f2cd473aad21ff025fbdcf851083d))
* **env/envconfig:** hoists env to a top level package, simplifies envconfig implementation. ([b5fbc38](https://github.com/matthewhartstonge/configurator/commit/b5fbc381adc0f75580b8e7d9246dc2a1cd3d4582))
* **env/envconfig:** implements `ConfigTypeable` `fmt.Stringer` and updates `Parse` to return the environment variable prefix. ([efd3036](https://github.com/matthewhartstonge/configurator/commit/efd3036117a3f7ca7a041aa8dce70090cc41db62))
* **file:** implements `fmt.Stringer`. ([4894201](https://github.com/matthewhartstonge/configurator/commit/4894201ae7ea118252510bb73d443ed5ff65b947))


### Bug Fixes

* **diag:** `Diagnostics` no longer attempts to append if supplied diags are empty. ([4306601](https://github.com/matthewhartstonge/configurator/commit/430660145717b578036e594bc430ac71699e85bc))
* **diag:** if `Diagnostics` is nil, builder will now return an empty `Diagnostics`. ([7466466](https://github.com/matthewhartstonge/configurator/commit/7466466cca5821ccd00b346a29aa3977bf2ce0bb))


### Miscellaneous Chores

* release v0.1.0 ([#10](https://github.com/matthewhartstonge/configurator/issues/10)) ([b6c523b](https://github.com/matthewhartstonge/configurator/commit/b6c523bef8aef1e6790a30c338410b9cb6cc5d8c))
