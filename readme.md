# Pushservice golang implementation

## Generate binary for linux

parser:
```
go build -ldflags="-s -w" -o bin/parser parser/main.go
```

sender:
```
go build -ldflags="-s -w" -o bin/sender sender/main.go
```

## Generate binary for mac

parser:
```
env GOOS=darwin go build -ldflags="-s -w" -o bin/parser parser/main.go
```

sender:
```
env GOOS=darwin go build -ldflags="-s -w" -o bin/sender sender/main.go
```

