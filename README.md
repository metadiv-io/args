# args

This package helps to get the arguments passed to the cli tool.

## Installation

```bash
go get -u github.com/WebXense/args
```

## Usage

```go
arg := args.New() // constructor will parse the os.Args

// all non-flag arguments will list in order
allArgs := arg.Args
```

All flag methods are called on this format:

.Method(key string, required bool, defaultValue string, alternativeKeys ...string) returnValue

-   key: the flag name, e.g. "name", you can call the flag with --name or -name
-   required: if the flag is required or not, if it is required and not passed, the program will exit with an error
-   defaultValue: the default value of the flag, if the flag is not passed, the default value will be returned
-   alternativeKeys: alternative keys for the flag, e.g. "n", you can call the flag with --n or -n

```go
// string
s := arg.FlagString("name", false, "default", "n")

// int
i := arg.FlagInt("age", false, 18, "a")

// uint
u := arg.FlagUint("heigh", false, uint(18), "h")

// float
f := arg.FlagFloat("score", false, 18.0, "s")

// bool
b := arg.FlagBool("married", "m") // no default value, pass this flag to set it to true
```