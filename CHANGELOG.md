## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2022-08-20
### Added
- Added full tests to `*Command.Run()` and `*Command.Execute()` can now be
tested as seen in `command_test.go` specifically the `TestCommand_Run` function
[see here](./command_test.go).
- Added a default `-help` flag to all commands created by `NewCommand`.

## [0.0.0] - 2022-08-20
### Added
- Type `Command` as the command structure.
- Type `Commands` as a data structure for multiple commands.
- The `NewCommand` function to create a default command.
- The `*Command.Add` and `*Command.AddCommands` methods to be able to add
multiple subcommands.

