## Simple project for learning the Go language.
My quest to become a Gopher.

### Updating go
[Download](https://go.dev/doc/install#install) the version for you OS, and
then remove the current installation:
```
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf /home/danielbevenius/Downloads/go1.19.1.linux-amd64.tar.gz
```
Then update the $PATH environment variable to include `/usr/local/go` if it is
not already included.
```console
$ go version
go version go1.19.1 linux/amd64
```

### Testing
```console
    go test src/hello_test.go
```

With additional logging:
```console
    go test -v src/hello_test.go
```

### Compiling
The `build` command invokes the compiler tool chain and will compile and link
all the file listed on the command line. 
```console
$ go build src/hello.go
```
This will create an executable named hello in the current directory. 
.
```console
$ go install src/hello.go
```
Will compile and copy the executable to `$GOPATH/bin`.

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


### Docs
Install:
```console
$ go get golang.org/x/tools/cmd/godoc
```

Access documentation for a package:
```console
$ go doc fmt
```

### Variables
Go does not allow unused local variables which will result in a compiler
error.

```
var msg string = "Hello"
```
If a variable is not initialized with will get a default value, for the example
above this would be an empty string "".

#### Short variable declaration
```
something := "bajja"
```

### Slices
Are three field data structures:
```
addr
length
capacity
```
You can use `make` to create a slice:
```go
slice := make([]string, 3, 5)
```
```go
slice := []string{"one", "two", "three"}
```
if you specify a value inside the [ ] operator, you’re creating an array. If
you don’t specify a value, you’re creating a slice.

A nil slice is created by declaring a slice without any initialization:
```go
var slice []int
```
They’re useful when you want to represent a slice that doesn’t exist, such as
when an exception occurs in a function that returns a slice 


### Modules
Are a new feature in 1.11 and allows projects to be in directories outside
of the GOPATH. 

You can initialize a project to use modules using:
```console
$ go mod init github.com/danbev/learning-go
```
Now, there is no tool required to use modules, you can simply add imports
to you source code files and then run build and go will get the dependency
and update go.mod (which was created by init).

You can also get specific versions of a dependency using go get:
```console
$ go get something@version

List all the deps:
```console
$ go list -m all
```

Prune deps that are no longer used:
```console
$ go mod tidy
```
