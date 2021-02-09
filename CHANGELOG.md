# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2021-02-08
### Added
- `mbctl` CLI tool with `config init`, `notifications create`, and `users generate-hmac` commands 
- goreleaser to build `mbctl` binary and release on new git tags
- `version` package to store versions of this library when built as a binary

### Changed
- `README.md` with CLI information
- `README.md` with link to `pkg.go.dev` documentation
- `README.md` with sending notification example

### Fixed
- `CustomAttributes` to `map[string]interface{}` to support multiple levels of custom attributes

## [0.1.0] - 2021-02-07
### Added
- Basic API usage with error handling
- `GenerateUserEmailHMAC` API method
- `CreateUser` API method
- `CreateUserC` API method
- `UpdateUser` API method
- `UpdateUserC` API method
- `CreateNotification` API method
- `CreateNotificationC` API method
