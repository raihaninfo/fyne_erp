# How do we compile for different platform?

## Install

```
sudo apt-get install -y *-w64-x86-*
```

```
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-H windowsgui"
```
