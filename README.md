## Simple project for learning the Go language.
My quest to become a Gopher.


### Setup
Run the ```setenv.sh``` script which will setup the correct GOPATH.

    . ./setenv.sh

### Building

    go install github.com/danbev/algo

### Testing

    go test github.com/danbev/algo

With additional logging:

    go test -v github.com/danbev/algo


### Compiling
The `build` command invokes the compiler tool chain and will compile and link
all the file listed on the command line. 
```console
$ go build src/hello.go
```
This will create an executable named hello.

Let's take a look at the libraries linked to our executable `hello` above.
```console
$ otool -L hello
hello:
	/usr/lib/libSystem.B.dylib (compatibility version 0.0.0, current version 0.0.0)
```
Size:
```console
$ ls -l -h hello
-rwxr-xr-x  1 danielbevenius  staff   2.0M Nov  8 12:08 hello
```


### Debugging
```console
$ export GOFLAGS="-ldflags=-compressdwarf=false"
$ go build -gcflags '-N -l' -o hello hello.go
```
```console
$ lldb hello
(lldb) br s -f hello.go --line 6
  Breakpoint 1: where = hello`main.main + 29, address = 0x000000000109957d
(lldb) r (lldb) bt
* thread #1, stop reason = step in
  * frame #0: 0x0000000001092e4d hello`fmt.Println at print.go:274
    frame #1: 0x00000000010995df hello`main.main at hello.go:6
    frame #2: 0x000000000102aeee hello`runtime.main at proc.go:203
    frame #3: 0x00000000010532f1 hello`runtime.goexit at asm_amd64.s:1357
```




