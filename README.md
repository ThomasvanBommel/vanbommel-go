# vanbommel-go
vanbommel.ca website written in go

## Run this project
Build n' run:
```bash
# Build n' run
make && ./run.build

# Clean up
make clean
```

<br />

Run without building:
```bash
make nobuild
```

<br />

Start project and restart after a modification has been made (for development):
```bash
# Requires inotify-tools to be installed
make watch
```

<br />

Run tests (coming soon)
```bash
make test
```