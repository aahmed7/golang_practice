# GOLANG PRACTICE

## Running a go module directly

```bash
go run <module-path>/<module-name>.go
```

## Building a go module

For building a go module, the module must be inside the directory defined by GOPATH env variable, where GOPATH has the following structure:

```bash
$GOPATH
├── bin
├── lib
└── src
```

and modules are located inside the 'src' directory
Run the following command to build a module

```bash
go build <module-name>
```

This will generate a binary calling directory.

## Installing a go module

Installing a go module refers to installing it to bin directory in the GOPATH.
To install a module:

```bash
go install <module-name>
```

## GOPATH settings update

Golang needs to run inside path pointed by GOPATH env variable. To update the GOPATH to use the current directory, run the following command:

On linux, source the file:

```bash
. set_gopath.sh
```

On Windows, run:

```bash
.\set-gopath.ps1
```

GOPATH env variable can have multiple paths, so its better to add the workspace to GOPATH in .profile in linux and env variables in Windows

## Ownership

This work has been taken from multiple sources, combined into one for a quick study. I do not take ownership for any part, except for the bash and pwsh scripts
