# Parser

## Generate binary for linux

parser:
```
go build -ldflags="-s -w" -o bin/parser parser.go
```

sender:
```
go build -ldflags="-s -w" -o bin/sender sender.go
```

## Generate binary for mac

parser:
```
env GOOS=darwin go build -ldflags="-s -w" -o bin/parser parser.go
```

sender:
```
env GOOS=darwin go build -ldflags="-s -w" -o bin/sender sender.go
```

