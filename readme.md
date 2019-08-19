# Pushservice golang implementation

## Generate binary for linux

parser:
```
go build -ldflags="-s -w" -o bin/parser parser/*
```

sender:
```
go build -ldflags="-s -w" -o bin/sender sender/*
```

## Generate binary for mac

parser:
```
env GOOS=darwin go build -ldflags="-s -w" -o bin/parser parser/*
```

sender:
```
env GOOS=darwin go build -ldflags="-s -w" -o bin/sender sender/*
```

## Install go & dep on linux

```
sudo snap install go --classic
mkdir -p /home/joynal3483/go/bin
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

configure env variable:
```
PATH=$PATH":/home/joynal3483/go/bin"
```

## clone repo

```
mkdir -p go/src
cd go/src/
git clone git@bitbucket.org:Joynal/pushservice-go.git
cd pushservice-go/
dep ensure
```