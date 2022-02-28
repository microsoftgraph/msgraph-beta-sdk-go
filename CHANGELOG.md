# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

### Changed

## [0.12.1] - 2022-02-28

### Changed

- Fixed a bug where http client configuration would impact the default client configuration for other usages.

## [0.12.0] - 2022-02-23

### Changed

- Weekly generation

## [0.11.0] - 2022-02-17

### Changed

- Updated to code 0.0.9 (fixed request body compression, added error deserialization)
- Weekly generation

## [0.10.0] - 2022-02-08

### Changed

- Fixed a deserialization issue with arrays of enums #73
- Updated to core 0.0.8 (request body compression, response body decompression, page iterator)
- Weekly generation

## [0.9.0] - 2022-02-01

### Changed

- Weekly generation
- Fixes a bug with collection properties serialization.

## [0.8.0] - 2022-01-25

### Changed

- Weekly generation
- Breaking: added types for Date-only, time-only and duration instead of using string.

## [0.7.0] - 2022-01-18

### Changed

- Weekly generation

## [0.6.0] - 2022-01-11

### Changed

- Weekly generation

## [0.5.0] - 2022-01-04

Happy new year!

### Changed

- Weekly generation

## [0.4.2] - 2021-12-07

### Changed

- Fixes an issue where escaped properties serialization would fail https://github.com/microsoftgraph/msgraph-sdk-go/issues/46

## [0.4.1] - 2021-12-01

### Changed

- Fixes an issue where select query parameters would not work https://github.com/microsoftgraph/msgraph-sdk-go/issues/31.

## [0.4.0] - 2021-11-30

### Changed

- Weekly generation

## [0.3.0] - 2021-11-23

### Changed

- Weekly generation
- Fixes doc comments

## [0.2.1] - 2021-11-19

### Changed

- Made the client options public so they can be used by consumers when customizing middleware

## [0.2.0] - 2021-11-17

### Changed

- Weekly generation

## [0.1.2] - 2021-11-12

### Changed

- Fixes #17 a bug where query parameters would be missing

## [0.1.1] - 2021-11-09

### Changed

- Fixes an issue where deserialization would fail because of nil values

## [0.1.0] - 2021-11-09

### Added

- Weekly generation

## [0.0.3] - 2021-11-08

### Added

- Added support for base url change

## [0.0.2] - 2021-11-04

### Changed

- Fixes a bug where collections of scalars would not deserialize properly

## [0.0.1] - 2021-10-22

### Added

- Initial release
