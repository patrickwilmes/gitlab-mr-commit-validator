# Gitlab MR commit validator

This tool checks if commits of an open merge request are matching
certain guidelines.

## Current state of development
Right now the guidelines are hardcoded and cannot be configured, but that's enough for
right now, considering the use case I have.

### Possible future development
- [ ] Support multiple backends (gitlab, github etc.)
- [ ] Include a commit linter and make guidelines configurable

# Building and running
Make sure to have the latest version of go installed

```bash
go build cmd/validator/
./validator <TOKEN>
```
