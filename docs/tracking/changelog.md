# SPIKE Changelog

## Recent

* Added configuration options for SPIKE Nexus and SPIKE Keeper.
* Updated quickstart guide.
* Max secret versions is now configurable.
* Introduced standard and configurable logging.
* Add community section and snapshots to the documentation.
* Added sqlite3 as a database backend.
* Enabled cross-compilation and SHA checksums.
* Now admin users can use jwt authentication and short-lived session tokens.

## [0.1.0] - 2024-11-06

### Added

* Implemented `put`, `read`, `delete`, `undelete`, and `list` functionalities.
* Created initial documentation, README, and related files.
* Compiled binaries targeting various platforms (x86, arm64, darwin, linux).
* SPIKE is demoable, however we need to update certain login and initialization
  flows.
* In memory secrets storage only (*using database as a backing store is coming up
  next*)
* Created a `jira.txt` to track things (*to avoid polluting GitHub issues
  unnecessarily*)
* This is an amazing start; more will come. Turtle power 🐢⚡️.