# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

### Changed

## [1.0.0] - 2023-05-04

### Changed

- GA Release + Weekly generation.

## [0.62.0] - 2023-04-25

### Changed

- Weekly generation.

## [0.61.0] - 2023-04-18

### Changed

- BREAKING: `client.UsersById("id").MessagesById("id").Get(context.Background())` now becomes `client.Users().ByUserId("id").Messages().ByMessageId("id").Get(context.Background())`
- Weekly generation.

## [0.60.0] - 2023-04-11

### Changed

- Weekly generation.

## [0.59.0] - 2023-04-06

### Changed

- Restored support for APIs under education/me.
- Weekly generation.

## [0.58.0] - 2023-03-28

### Changed

- Weekly generation.

## [0.57.0] - 2023-03-21

### Changed

- Weekly generation.

## [0.56.0] - 2023-03-20

### Changed

- Weekly generation.

## [0.55.0] - 2023-03-03

### Changed

- shimmed "me" package, backing store , uuid fix and reduced file count.


## [0.54.0] - 2023-01-24

### Changed

- Weekly generation.
- Updated dependencies.

## [0.53.0] - 2023-01-17

### Changed

- Weekly generation.


## [0.52.0] - 2023-01-10

### Changed

- Weekly generation.
- BREAKING: createXXXRequestInformation methods are now named toXXXRequestInformation.

## [0.51.0] - 2022-12-28

### Changed

- Weekly generation.

## [0.50.0] - 2022-12-14

### Added

- BREAKING: added support for multi-valued request headers.

### Changed

- BREAKING: shortens the type names by removing the redudant first level path segment.

## [0.49.0] - 2022-12-02

### Changed

- BREAKING: the package structure has been flattened, making most imports much shorter.
- Significant improvements in build time.
- Weekly generation.

## [0.48.0] - 2022-11-22

### Changed

- Weekly generation.

## [0.47.0] - 2022-11-08

### Changed

- Weekly generation.

## [0.46.0] - 2022-11-01

### Changed

- Weekly generation.

## [0.45.0] - 2022-10-18

### Changed

- Weekly generation.

## [0.44.0] - 2022-10-11

### Changed

- Weekly generation.

## [0.43.0] - 2022-10-04

### Changed

- Weekly generation.
- BREAKING removed overloads to get request information.

## [0.42.1] - 2022-09-28

### Changed

- Fixed a regression by the pipeline where simple authentication initialization would be missing.

## [0.42.0] - 2022-09-27

### Changed

- Weekly generation.

## [0.41.0] - 2022-09-20

### Changed

- Weekly generation.

## [0.40.0] - 2022-09-16

### Changed

- Weekly generation.

## [0.39.1] - 2022-09-09

### Changed

- Updated references to core libraries.

## [0.39.0] - 2022-09-06

### Changed

- Weekly generation.

## [0.38.0] - 2022-09-02

### Changed

- Added a context parameter to execution method.
- Removed `WithConfigurationAndResponseHandler` overloads to reduce the size of the SDK.
> `client.Me().Get()` now becomes `client.Me().Get(context.Background(), nil)` (or pass some higher level context).  
> `client.Me().GetWithConfigurationAndResponseHandler(requestConfigurationValue, nil)` now becomes `client.Me().Get(context, requestConfigurationValue)`.  
> `client.Me().GetWithConfigurationAndResponseHandler(nil, responseHandler)` now becomes `client.Me().Get(context, requestConfigurationValue)` where the request configuration has a `RequestHandlerOption` in its options (from abstractions).  
> The body argument (POST, PATCH, PUT) is the second argument. (after the context)

## [0.37.0] - 2022-08-30

### Changed

- Weekly generation.

## [0.36.0] - 2022-08-23

### Changed

- Weekly generation.

## [0.35.0] - 2022-08-17

### Changed

- Weekly generation.

## [0.34.0] - 2022-08-10

### Changed

- Weekly generation.

## [0.33.0] - 2022-08-04

### Changed

- Weekly generation.

## [0.32.0] - 2022-07-20

### Changed

- Weekly generation.

## [0.31.0] - 2022-07-13

### Changed

- Fixed an issue where the odata type property would not serialize properly and make most POST/PUT/PATCH operations fail.

## [0.30.0] - 2022-07-12

### Changed

- Weekly generation.

## [0.29.0] - 2022-06-21

### Changed

- Weekly generation.

## [0.28.0] - 2022-06-20

### Changed

- Weekly generation
- Added missing error mapping for OData actions/functions.

## [0.27.0] - 2022-06-08

### Changed

- Weekly generation.
- Updated kiota references to solve for YAML dependency security bulletin.

## [0.26.0] - 2022-05-31

### Changed

- Weekly generation

## [0.25.0] - 2022-05-27

### Changed

- Weekly generation
- Updated graph core and kiota libraries reference to add support for enum and enum collections responses.

## [0.24.1] - 2022-05-25

### Changed

- Updated reference to microsoft graph core to solve for version misalignment.

## [0.24.0] - 2022-05-10

### Changed

- Weekly generation.

## [0.23.0] - 2022-05-03

### Changed

- Breaking: the parameters have been changed to request body, request configuration(query parameters, options, headers), response handler for a better experience.
- Upgraded JSON reference to fix multiple JSON serialization issues.

## [0.22.0] - 2022-04-28

### Changed

- Weekly generation

## [0.21.0] - 2022-04-19

### Changed

- Upgraded core library to address uri template with quotes issue.
- Upgraded to go 18.

## [0.20.1] - 2022-04-14

### Changed

- Fixed an issue with date json serialization

## [0.20.0] - 2022-04-14

### Changed

- Fixed an issue where query parameters would be missing special characters. #130

## [0.19.0] - 2022-04-12

### Added

- Added support for special characters in query parameters names.

### Changed

- Breaking: updated reference to kiota libraries. (deserialization simplification)
- Weekly generation

## [0.18.0] - 2022-04-05

### Added

- Added support for vendor specific content types.
- Added support for 204 no content responses.

### Changed

- Weekly generation
- Breaking: kiota dependencies have been moved to their own repository.

## [0.17.0] - 2022-03-29

### Changed

- Weekly generation.

## [0.16.0] - 2022-03-22

### Added

- Adds support for text endpoints deserialization. (.Count())

### Changed

- Weekly generation.
- Updated core reference for serialization nil check fix.
- Breaking change: simpler API design for page iterator.

## [0.15.0] - 2022-03-15

### Changed

- Updated reference to core for new serialization types
- Weekly generation

## [0.14.0] - 2022-03-08

### Changed

- Weekly generation, based of the new OpenAPI conversion, lots of new endpoints available.
- Fixed an issue with collections responses for OData functions.
- Breaking: Fixed a bug where using the raw request URL would result in invalid requests.
- Breaking: Splat Parsable interface to AdditionalPropertiesHolder.

## [0.13.0] - 2022-03-01

### Changed

- Weekly generation

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
