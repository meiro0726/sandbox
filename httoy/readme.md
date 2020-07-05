# httoy

## build

build on WSL2 and use on windows10.

``` bash
$ go vet ./... && golint ./... && GOOS=windows GOARCH=amd64 go build -o httoy.exe && cp ./httoy.exe /mnt/c/Users/[username]/Desktop/
```
## License

MIT
