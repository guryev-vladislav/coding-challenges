# Algorithmic Task Solutions

Solutions to tasks from programming competitions and practice contests. All solutions are written in Go.

## Structure

Directories are organized by platform, year, and individual task:

```text
<platform>/<year>/<task>/
├── README.md
├── <task>.go
└── <task>_test.go  # when automated tests are available
```

- [True Tech Champ 2025](truetechchamp/2025/README.md)
- [Yandex Contest 2025](yandex_contest/2025/README.md)

## Running Tests

Run this command from the repository root:

```bash
go test ./...
```

Tests are stored next to the corresponding solution and belong to the same Go package. To run tests for one task:

```bash
go test ./truetechchamp/2025/max_cashback
```

Tasks without a `_test.go` file are currently checked only by compilation and manual execution with the examples from their README.

## Files

- Source code is stored in the competition and year directories listed above.
