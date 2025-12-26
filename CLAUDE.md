# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a structured Go learning repository following an 8-week backend mastery course. The repository contains exercises, mini-projects, and a capstone task queue API project.

## Repository Structure

- `exercises/` - Weekly coding exercises organized by topic (week1-basics through week5-patterns)
- `mini-projects/` - Small projects for practice (configurable-logger, url-fetcher)
- `taskqueue/` - Capstone project: REST API with background worker queue
  - `cmd/server/` - Application entry point
  - `internal/` - Private application code (handler, repository, worker)
  - `pkg/` - Public packages
- `syllabus.md` - Course curriculum and schedule

## Common Commands

### Running Go Programs
```bash
go run <filename>.go                    # Run a single Go file
go run exercises/week1-basics/day1.go   # Run specific exercise
```

### Building
```bash
go build <filename>.go                  # Build executable
go build -o <binary_name> <filename>.go # Build with custom name
```

### Module Management
```bash
go mod init <module_name>               # Initialize new module
go mod tidy                             # Clean up dependencies
go get <package>                        # Add dependency
```

### For Capstone Project (taskqueue/)
```bash
cd taskqueue/
go run cmd/server/main.go               # Run the task queue server
go build -o taskqueue cmd/server/main.go # Build the server binary
```

## Development Notes

- Go version: 1.25.1
- Module name: `learn-golang`
- The course follows a progressive structure from basic language concepts to production-ready backend development
- Exercises are meant to be completed in sequence following the syllabus
- The capstone project demonstrates production patterns: clean architecture, repository pattern, middleware, graceful shutdown
- Each week builds upon previous concepts, culminating in a full REST API with background job processing