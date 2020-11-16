## Running a go module directly
```
go run <module-path>/<module-name>.go
```

## Building a go module
For building a go module, the module must be inside the directory defined by GOPATH env variable, where GOPATH has the following structure:
```
$GOPATH
├── bin
├── lib
└── src
```
and modules are located inside the 'src' directory
Run the following command to build a module
```
go build <module-name>
```
This will generate a binary calling directory.

## Installing a go module
Installing a go module refers to installing it to bin directory in the GOPATH.
To install a module:
```
go install <module-name>
```

## Ownership
This work has been taken from multiple sources, combined into one for a quick study. I do not take ownership for any part.