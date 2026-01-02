# v1.2

## What changed?

Download the binary for your platform from the assets below and run it directly.

## Additions

* Restructure codebase into standard Go CLI project layout for better maintainability
* Add cross-platform build support for Windows, macOS, and Linux (multiple architectures)
* Add Makefile with build, build-all, release, and clean targets
* Add PowerShell and Bash build scripts for cross-platform building
* Add release preparation script to create user-friendly binary names
* Improve documentation with clearer installation instructions and step-by-step guides
* Add .gitignore with proper exclusions for build artifacts and releases

## Improvements

* Better code organization with separated packages (models, themes, styles, config, app)
* Simplified build process with automated scripts
* Enhanced README with professional formatting and clearer structure
* Improved project structure following Go best practices
* Better separation of concerns across internal packages

## Fixes

* Fix import cycles by reorganizing package structure
* Fix build reliability across different platforms
* Clean up unused directories and files from restructuring

