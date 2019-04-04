# go-assetfs-example

Example of [go-assetfs](https://github.com/moisespsena-go/assetfs) project.

See source code for details.

## Get it

```bash
go get github.com/moisespsena-go/assetfs-example
```

Go to project directory:
 
```bash
cd $GOPATH/src/github.com/moisespsena-go/assetfs-example
```

## Start

Init your repository:

```bash
cd 
go run main.go -tags init_repository
```

## Before build and export your project 

Pre-compile your Repository:

```bash
go run main.go -tags pre_compile
```

## Build and export your project

```bash
go build -tags bindata
```

## Clean pre compiled files

```bash
go build -tags bindata_clean
```